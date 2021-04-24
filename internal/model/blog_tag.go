package model

// BlogTag 標籤管理
type BlogTag struct {
	*BlogBasicInfo
	Name  string `json:"name"`  // 標籤名稱
	State uint8  `json:"state"` // 狀態 0為禁用、1為啟用
}

// TableName get sql table name.获取数据库表名
func (m BlogTag) TableName() string {
	return "blog_tag"
}
