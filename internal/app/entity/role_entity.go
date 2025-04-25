package entity

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Permission uint8

const (
	General       Permission = iota // 0
	Administrator                   // 1
	Root                            // 2
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
	case Root:
		return "Root"
	case Administrator:
		return "Administrator"
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

func (e *RoleEntity) IsRoot() bool {
	return e.Permission == Root
}

func toPermission(permission int) (Permission, error) {
	switch permission {
	case int(General):
		return General, nil
	case int(Administrator):
		return Administrator, nil
	case int(Root):
		return Root, nil
	}
	return General, errors.New("invalid permission")
}
