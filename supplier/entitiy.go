package supplier

import (
	"cheggstore/user"
	"time"
)

type Supplier struct {
	ID        int
	UserID    int
	Name      string
	Address   string
	Postal    string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.User
}
