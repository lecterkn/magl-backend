package entity

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Permission int

const (
	Standard      Permission = iota // 0
	Administrator                   // 1
	Root                            // 2
)

type RoleEntity struct {
	Permission Permission `validate:"required,min=0,max=2"`
}

func NewRoleEntity(permission int) (*RoleEntity, error) {
	perm, err := toPermission(permission)
	if err != nil {
		return nil, err
	}
	roleEntity := RoleEntity{
		Permission: perm,
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(roleEntity); err != nil {
		return nil, err
	}
	return &roleEntity, nil
}

func (e *RoleEntity) IsStandard() bool {
	return e.Permission == Standard
}

func (e *RoleEntity) IsAdministrator() bool {
	return e.Permission == Administrator
}

func (e *RoleEntity) IsRoot() bool {
	return e.Permission == Root
}

func toPermission(permission int) (Permission, error) {
	switch permission {
	case int(Standard):
		return Standard, nil
	case int(Administrator):
		return Administrator, nil
	case int(Root):
		return Root, nil
	}
	return Standard, errors.New("invalid permission")
}
