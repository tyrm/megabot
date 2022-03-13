package models

import (
	"github.com/google/uuid"
	"github.com/tyrm/megabot/internal/id"
	"time"
)

// GroupMembership represents a user's membership in a group
type GroupMembership struct {
	ID        string    `validate:"required,ulid" bun:"type:CHAR(26),pk,nullzero,notnull,unique"`
	CreatedAt time.Time `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UserID    string    `validate:"required,ulid" bun:"type:CHAR(26),unique:groupmembership,notnull,nullzero"`
	User      *User     `validate:"-" bun:"rel:belongs-to"`
	GroupID   uuid.UUID `validate:"required,uuid" bun:",unique:groupmembership,notnull,nullzero"`
}

// GenID generates a new id for the object
func (g *GroupMembership) GenID() error {
	if g.ID == "" {
		newID, err := id.NewULID()
		if err != nil {
			return err
		}
		g.ID = newID
	}
	return nil
}

// GroupMemberships contains multiple groups
type GroupMemberships []*GroupMembership

// GenID generates a new id for the object
func (g *GroupMemberships) GenID() error {
	for _, group := range *g {
		err := group.GenID()
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
