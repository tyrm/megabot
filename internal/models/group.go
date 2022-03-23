package models

import (
	"context"
	"github.com/google/uuid"
	"github.com/tyrm/megabot/internal/id"
	"github.com/uptrace/bun"
	"time"
)

// GroupMembership represents a user's membership in a group
type GroupMembership struct {
	ID        string    `validate:"required,ulid" bun:"type:CHAR(26),pk,nullzero,notnull,unique"`
	CreatedAt time.Time `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UserID    string    `validate:"required,ulid" bun:"type:CHAR(26),unique:groupmembership,notnull,nullzero"`
	User      *User     `validate:"-" bun:"rel:belongs-to,join:user_id=id"`
	GroupID   uuid.UUID `validate:"required" bun:",unique:groupmembership,notnull,nullzero"`
}

var _ bun.BeforeAppendModelHook = (*GroupMembership)(nil)

// BeforeAppendModel runs before a bun append operation
func (gm *GroupMembership) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		if gm.ID == "" {
			newID, err := id.NewULID()
			if err != nil {
				return err
			}
			gm.ID = newID
		}

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
