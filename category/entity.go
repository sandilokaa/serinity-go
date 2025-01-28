package category

import (
	"serinitystore/user"
	"time"
)

type Category struct {
	ID        int
	UserID    int
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.User
}
