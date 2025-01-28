package sizechart

import (
	"serinitystore/user"
	"time"
)

type SizeChart struct {
	ID        int
	UserID    int
	Name      string
	FileName  string
	User      user.User
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SizeChart) TableName() string {
	return "SizeCharts"
}
