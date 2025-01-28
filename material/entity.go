package material

import (
	"serinitystore/user"
	"time"
)

type Material struct {
	ID           int
	UserID       int
	MaterialName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         user.User
}
