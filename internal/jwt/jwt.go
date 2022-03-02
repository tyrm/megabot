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
	"github.com/tyrm/megabot/internal/id"
	"github.com/tyrm/megabot/internal/kv"
	"github.com/tyrm/megabot/internal/models"
	"net/http"
	"strings"
	"time"
)

type Module struct {
	kv kv.JWT
}

func New(kv kv.JWT) (*Module, error) {
	return &Module{
		kv: kv,
	}, nil
}

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

type accessDetails struct {
	AccessID string
	UserID   string
	Groups   []uuid.UUID
}

type tokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessID     string
	RefreshID    string
	AtExpires    int64
	RtExpires    int64
}

func (m *Module) CreateAuth(ctx context.Context, userid string, td *tokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := m.kv.SetJWTAccessToken(ctx, td.AccessID, userid, at.Sub(now))
	if errAccess != nil {
		logrus.Debugf("can't save access token: %s", errAccess.Error())
		return errAccess
	}
	errRefresh := m.kv.SetJWTRefreshToken(ctx, td.RefreshID, userid, rt.Sub(now))
	if errRefresh != nil {
		logrus.Debugf("can't save refresh token: %s", errRefresh.Error())
		return errRefresh
	}
	return nil
}

func (m *Module) CreateToken(ctx context.Context, user *models.User) (*tokenDetails, error) {
	td := &tokenDetails{}
	td.AtExpires = time.Now().Add(viper.GetDuration(config.Keys.AccessExpiration)).Unix()
	newAccessToken, err := id.NewRandomULID()
	if err != nil {
		return nil, err
	}
	td.AccessID = newAccessToken

	td.RtExpires = time.Now().Add(viper.GetDuration(config.Keys.RefreshExpiration)).Unix()
	td.RefreshID = td.AccessID + "++" + user.ID

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims[claimAuthorized] = true
	atClaims[claimAccessID] = td.AccessID
	atClaims[claimUserID] = user.ID
	atClaims[claimExpires] = td.AtExpires
	atClaims[claimGroups] = user.Groups
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	td.AccessToken, err = at.SignedString(viper.GetString(config.Keys.AccessSecret))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims[claimRefreshID] = td.RefreshID
	rtClaims[claimUserID] = user.ID
	rtClaims[claimExpires] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS512, rtClaims)
	td.RefreshToken, err = rt.SignedString(viper.GetString(config.Keys.RefreshSecret))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (m *Module) deleteTokens(ctx context.Context, authD *accessDetails) error {
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

func (m *Module) extractTokenMetadata(r *http.Request) (*accessDetails, error) {
	token, err := m.verifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		metadata := accessDetails{
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

func (m *Module) fetchAuth(ctx context.Context, authD *accessDetails) (string, error) {
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
		return viper.GetString(config.Keys.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
