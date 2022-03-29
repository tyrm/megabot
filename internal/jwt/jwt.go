package jwt

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
	"github.com/tyrm/megabot/internal/db"
	"github.com/tyrm/megabot/internal/id"
	"github.com/tyrm/megabot/internal/kv"
	"github.com/tyrm/megabot/internal/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Module is a module for processing jwt tokens
type Module struct {
	db db.DB
	kv kv.JWT
}

// New creates a new JWT module
func New(db db.DB, kv kv.JWT) (*Module, error) {
	return &Module{
		db: db,
		kv: kv,
	}, nil
}

// Close closes the jwt module
func (m Module) Close() error {
	return nil
}

const (
	claimGroups     = "groups"
	claimAccessID   = "access_id"
	claimAuthorized = "authorized"
	claimExpires    = "exp"
	claimRefreshID  = "refresh_id"
	claimUserID     = "user_id"
)

// AccessDetails contains data stored in the jwt token
type AccessDetails struct {
	AccessID string
	UserID   string
	Groups   []uuid.UUID
}

// TokenDetails contains the metadata for a token
type TokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessID     string
	RefreshID    string
	AtExpires    int64
	RtExpires    int64
}

// CreateAuth inserts the token data into the KV
func (m *Module) CreateAuth(ctx context.Context, userid int64, td *TokenDetails) error {
	l := logrus.WithField("func", "CreateAuth")

	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := m.kv.SetJWTAccessToken(ctx, td.AccessID, userid, at.Sub(now))
	if errAccess != nil {
		l.Debugf("can't save access token: %s", errAccess.Error())
		return errAccess
	}
	errRefresh := m.kv.SetJWTRefreshToken(ctx, td.RefreshID, userid, rt.Sub(now))
	if errRefresh != nil {
		l.Debugf("can't save refresh token: %s", errRefresh.Error())
		return errRefresh
	}
	return nil
}

// CreateToken creates a token based on a user
func (m *Module) CreateToken(ctx context.Context, user *models.User) (*TokenDetails, error) {
	l := logrus.WithField("func", "CreateToken")

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(viper.GetDuration(config.Keys.AccessExpiration)).Unix()
	newAccessToken, err := id.NewRandomULID()
	if err != nil {
		l.Errorf("generating id: %s", err.Error())
		return nil, err
	}
	td.AccessID = newAccessToken

	td.RtExpires = time.Now().Add(viper.GetDuration(config.Keys.RefreshExpiration)).Unix()
	td.RefreshID = td.AccessID + "++" + strconv.FormatInt(user.ID, 10)

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims[claimAuthorized] = true
	atClaims[claimAccessID] = td.AccessID
	atClaims[claimUserID] = user.ID
	atClaims[claimExpires] = td.AtExpires
	atClaims[claimGroups] = user.Groups
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	td.AccessToken, err = at.SignedString([]byte(viper.GetString(config.Keys.AccessSecret)))
	if err != nil {
		l.Errorf("access token signed string: %s", err.Error())
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims[claimRefreshID] = td.RefreshID
	rtClaims[claimUserID] = user.ID
	rtClaims[claimExpires] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS512, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(viper.GetString(config.Keys.RefreshSecret)))
	if err != nil {
		l.Errorf("refresh token signed string: %s", err.Error())
		return nil, err
	}
	return td, nil
}

// DeleteRefreshToken deletes a refresh token from KV
func (m *Module) DeleteRefreshToken(ctx context.Context, refreshToken string) error {
	return m.kv.DeleteJWTRefreshToken(ctx, refreshToken)
}

// DeleteTokens deletes tokens from the KV
func (m *Module) DeleteTokens(ctx context.Context, authD *AccessDetails) error {
	// get the refresh id
	refreshUUID := fmt.Sprintf("%s++%s", authD.AccessID, authD.UserID)
	// delete access token
	err := m.kv.DeleteJWTAccessToken(ctx, authD.AccessID)
	if err != nil {
		return err
	}
	// delete refresh token
	err = m.kv.DeleteJWTRefreshToken(ctx, refreshUUID)
	if err != nil {
		return err
	}
	return nil
}

func (m *Module) extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// ExtractTokenMetadata gets the token from the bearer token
func (m *Module) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := m.verifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		metadata := AccessDetails{
			AccessID: claims[claimAccessID].(string),
			UserID:   claims[claimUserID].(string),
		}
		if claims[claimGroups] != nil {
			groups := claims[claimGroups].([]interface{})
			groupIds := make([]uuid.UUID, len(groups))
			for i, g := range groups {
				gu, err := uuid.Parse(g.(string))
				if err != nil {
					logrus.Tracef("%s is not a uuid: %s", g, err.Error())
					return nil, err
				}
				groupIds[i] = gu
			}
			metadata.Groups = groupIds
		}

		return &metadata, nil
	}
	return nil, err
}

func (m *Module) fetchAuth(ctx context.Context, authD *AccessDetails) (string, error) {
	userid, err := m.kv.GetJWTAccessToken(ctx, authD.AccessID)
	if err != nil {
		return "", err
	}
	if authD.UserID != userid {
		return "", errors.New("unauthorized")
	}
	return userid, nil
}

func (m *Module) tokenValid(r *http.Request) error {
	token, err := m.verifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return err
	}
	return nil
}

// Parse, validate, and return a token.
// keyFunc will receive the parsed token and should return the key for validating.
func (m *Module) verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := m.extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString(config.Keys.AccessSecret)), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// RefreshAccessToken generates a new access token for a given refresh token
func (m *Module) RefreshAccessToken(ctx context.Context, refreshToken string) (*TokenDetails, error) {
	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString(config.Keys.RefreshSecret)), nil
	})

	//if there is an error, the token must have expired
	if err != nil {
		logrus.Tracef("token error: %s", err.Error())
		return nil, err
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, errUnauthorized
	}

	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		// read key data
		refreshString, ok := claims[claimRefreshID].(string) //convert the interface to string
		if !ok {
			logrus.Tracef("claim %s missing", claimRefreshID)
			return nil, errUnprocessableEntity
		}

		// get user
		user, err := m.db.ReadUserByID(ctx, claims[claimUserID].(int64))
		if err != nil {
			logrus.Errorf("getting user: %s", err.Error())
			return nil, err
		}
		if user == nil {
			return nil, errUnauthorized
		}

		// Delete the previous Refresh Token
		err = m.DeleteRefreshToken(ctx, refreshString)
		if err != nil {
			logrus.Errorf("kv error: %s", err.Error())
			return nil, err
		}

		// Create new pairs of refresh and access tokens
		ts, createErr := m.CreateToken(ctx, user)
		if createErr != nil {
			logrus.Tracef("error creating token: %s", createErr)
			return nil, createErr
		}

		// save the tokens metadata to kv
		saveErr := m.CreateAuth(ctx, user.ID, ts)
		if saveErr != nil {
			logrus.Tracef("error saving token: %s", createErr)
			return nil, saveErr
		}

		return ts, nil
	}

	return nil, errRefreshExpired
}
