package output

import (
	"time"

	"github.com/google/uuid"
)

type CategoryQueryOutput struct {
	Id          uuid.UUID
	Name        string
	Description string
	ImageUrl    *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
