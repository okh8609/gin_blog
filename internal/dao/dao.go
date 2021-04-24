package dao

import "gorm.io/gorm"

type Dao struct{
	gormDB *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{gormDB: engine}
}