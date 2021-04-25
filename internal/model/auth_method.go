package model

import (
	"gorm.io/gorm"
)

func (obj Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	err := db.Where("uuid = ?", obj.UUID).First(&auth).Error
	return auth, err
}

func (obj Auth) Create(db *gorm.DB) error {
	return db.Create(&obj).Error
}

func (obj Auth) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&obj).Where("uuid = ?", obj.UUID).Updates(values).Error
}

func (obj Auth) Delete(db *gorm.DB) error {
	return db.Where("uuid = ?", obj.UUID).Delete(&obj).Error
}
