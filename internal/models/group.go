package models

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

// GroupMembership represents a user's membership in a group
type GroupMembership struct {
	ID        int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UserID    int64     `validate:"min=1" bun:",unique:groupmembership,notnull,nullzero"`
	User      *User     `validate:"-" bun:"rel:belongs-to,join:user_id=id"`
	GroupID   []byte    `validate:"required" bun:",unique:groupmembership,notnull,nullzero"`
}

var _ bun.BeforeAppendModelHook = (*GroupMembership)(nil)

// BeforeAppendModel runs before a bun append operation
func (gm *GroupMembership) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		gm.CreatedAt = now
		gm.UpdatedAt = now

		err := validate.Struct(gm)
		if err != nil {
			return err
		}
	case *bun.UpdateQuery:
		gm.UpdatedAt = time.Now()

		err := validate.Struct(gm)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetGroupID sets bytes for GroupID
func (gm *GroupMembership) GetGroupID() (uuid.UUID, error) {
	if len(gm.GroupID) == 0 {
		return uuid.Nil, nil
	}
	return uuid.FromBytes(gm.GroupID)
}

// SetGroupID sets bytes for GroupID
func (gm *GroupMembership) SetGroupID(u uuid.UUID) error {
	d, err := u.MarshalBinary()
	if err != nil {
		return err
	}
	gm.GroupID = d
	return nil
}

// groupSuperAdmin is the uuid of the Super Administrators group
var groupSuperAdmin = uuid.Must(uuid.Parse("11a08aec-b7e0-46b4-ba53-e95a858d4cad"))

// groupName contains the names of the groups.
var groupName = map[string]uuid.UUID{
	"admin": groupSuperAdmin,
}

// groupTitle contains the titles of the groups.
var groupTitle = map[uuid.UUID]string{
	groupSuperAdmin: "Super Admin",
}

// GroupSuperAdmin returns the uuid for the Super Admin
func GroupSuperAdmin() uuid.UUID {
	return groupSuperAdmin
}

// GroupName return a uuid for the group name
func GroupName(g string) uuid.UUID {
	if s, ok := groupName[g]; ok {
		return s
	}
	return uuid.Nil
}

// GroupTitle return a pretty text name for the group
func GroupTitle(g uuid.UUID) string {
	if s, ok := groupTitle[g]; ok {
		return s
	}
	return ""
}
