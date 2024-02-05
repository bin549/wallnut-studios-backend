package models

type Project struct {
	ID        int64  `binding:"required"`
	Name      string `binding:"required"`
	TypeId    int64  `binding:"required"`
	ProfileId int64  `binding:"required"`
}

var projects = []Project{}
