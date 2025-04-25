package entity

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Permission uint8

const (
	General       Permission = iota // 0
	Moderator                       // 1
	Administrator                   // 2
	Root                            // 3
)

type RoleEntity struct {
	Permission Permission `validate:"min=0,max=2"`
}

func NewRoleEntity(permission int) (*RoleEntity, error) {
	perm, err := toPermission(permission)
	if err != nil {
		return nil, err
	}
	fmt.Println(perm)
	roleEntity := RoleEntity{
		Permission: perm,
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(roleEntity); err != nil {
		return nil, err
	}
	return &roleEntity, nil
}

func (e *RoleEntity) GetPermission() string {
	switch e.Permission {
	case Moderator:
		return "Moderator"
	case Administrator:
		return "Administrator"
	case Root:
		return "Root"
	default:
		return "General"
	}
}

func (e *RoleEntity) IsStandard() bool {
	return e.Permission == General
}

func (e *RoleEntity) IsAdministrator() bool {
	return e.Permission == Administrator
}

func (e *RoleEntity) IsModerator() bool {
	return e.Permission == Moderator
}

func (e *RoleEntity) IsRoot() bool {
	return e.Permission == Root
}

func toPermission(permission int) (Permission, error) {
	switch permission {
	case int(General):
		return General, nil
	case int(Moderator):
		return Moderator, nil
	case int(Administrator):
		return Administrator, nil
	case int(Root):
		return Root, nil
	}
	return General, errors.New("invalid permission")
}
