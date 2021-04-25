package model

// Auth 帳戶認證資訊
type Auth struct {
	UUID     string `gorm:"primaryKey;column:uuid;type:varchar(255);not null" json:"-"` // UUID或帳號
	Password string `gorm:"column:password;type:varchar(255);not null" json:"password"` // 密碼
}

// TableName get sql table name.获取数据库表名
func (obj Auth) TableName() string {
	return "auth"
}
