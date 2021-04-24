package model

import "gorm.io/gorm"

func (t BlogTag) Count(db *gorm.DB) (int64, error) {
	db = db.Model(&BlogTag{})

	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	db = db.Where("is_del = ?", 0)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t BlogTag) Get(db *gorm.DB) (BlogTag, error) {
	var tag BlogTag
	err := db.Where("id = ? AND is_del = ? AND state = ?", t.ID, 0, t.State).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}

	return tag, nil
}

func (t BlogTag) List(db *gorm.DB, pageOffset, pageSize int) ([]*BlogTag, error) {
	var tags []*BlogTag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	db = db.Where("is_del = ?", 0)
	if err = db.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t BlogTag) ListByIDs(db *gorm.DB, ids []uint32) ([]*BlogTag, error) {
	var tags []*BlogTag
	db = db.Where("state = ? AND is_del = ?", t.State, 0)
	db = db.Where("id IN (?)", ids)
	err := db.Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

func (t BlogTag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t BlogTag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error
}

func (t BlogTag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.ID, 0).Delete(&t).Error
}
