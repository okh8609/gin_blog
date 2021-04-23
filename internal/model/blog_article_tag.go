package model

// BlogArticleTag 文章標籤關聯
type BlogArticleTag struct {
	*BasicInfo
	TagID     uint `json:"tab_id"`     // 標籤ID
	ArticleID uint `json:"article_id"` // 文章ID

}

// TableName get sql table name.获取数据库表名
func (m BlogArticleTag) TableName() string {
	return "blog_article_tag"
}
