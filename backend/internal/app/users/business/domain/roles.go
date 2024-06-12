package domain

import (
	"fmt"
)

// RoleUser represents a role of user type
var RoleUser = Role{1, "USER"}

// roles is a map representing a set of roles
var roles = map[RoleType]Role{
	RoleUser.roleType: RoleUser,
}

// RoleID represents the role unique identifier
type RoleID uint8

func (r RoleID) Validate() (*RoleID, error) {
	for _, v := range roles {
		if v.roleID.ID() == uint8(r) {
			return &r, nil
		}
	}

	return &r, fmt.Errorf("%w: invalid role type id '%v'", InvalidRoleTypeID, uint8(r))
}

func (r RoleID) ID() uint8 {
	return uint8(r)
}

// RoleType represents a role type
type RoleType string

func (r RoleType) Validate() (*RoleType, error) {
	role, exists := roles[r]

	if !exists {
		return &r, fmt.Errorf("%w: invalid role type '%v'", InvalidRoleType, r)
	}

	return &role.roleType, nil
}

func (r RoleType) RoleType() string {
	return string(r)
}

// Role is a structure containing the values needed to define a role
type Role struct {
	roleID   RoleID
	roleType RoleType
}

func NewRole(roleID uint8, roleType string) (*Role, error) {
	roleIDVO, err := RoleID(roleID).Validate()
	if err != nil {
		return &Role{}, nil
	}

	roleTypeVO, err := RoleType(roleType).Validate()
	if err != nil {
		return &Role{}, nil
	}

	return &Role{
		roleID:   *roleIDVO,
		roleType: *roleTypeVO,
	}, nil
}

func (r *Role) ID() uint8 {
	return r.roleID.ID()
}

func (r *Role) Role() string {
	return r.roleType.RoleType()
}
