package output

import (
	"time"

	"github.com/google/uuid"
)

type UserOutput struct {
	Id        uuid.UUID
	Name      string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
