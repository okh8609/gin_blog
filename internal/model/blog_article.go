package model

// BlogArticle 文章管理
type BlogArticle struct {
	*BasicInfo
	Title         string `json:"title"`           // 文章標題
	Desc          string `json:"desc"`            // 文章簡述
	Content       string `json:"content"`         // 文章內容
	CoverImageURL string `json:"cover_image_url"` // 封面圖片地址
	State         uint8  `json:"state"`           // 狀態 0為禁用、1為啟用
}

// TableName get sql table name.获取数据库表名
func (m BlogArticle) TableName() string {
	return "blog_article"
}
