package sizechart

import "cheggstore/user"

type CreateSizeChartInput struct {
	User     user.User
	Name     string `form:"name" binding:"required"`
	FileName string
}

type SizeChartInputDetail struct {
	ID int `uri:"id" binding:"required"`
}

type UpdateSizeChartInput struct {
	User     user.User
	Name     string `form:"name" binding:"required"`
	FileName string
}
