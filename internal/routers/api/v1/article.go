package v1

import "github.com/gin-gonic/gin"

type Article struct{}

func NewArticle() *Article {
	return &Article{}
}

// @Summary 取得單篇文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.BlogArticle "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [get]
func (t *Article) Get(c *gin.Context) {}

// @Summary 取得多篇文章
// @Produce json
// @Param name query string false "文章名稱"
// @Param tag_id query int false "標籤ID"
// @Param state query int false "狀態"
// @Param page query int false "頁碼"
// @Param page_size query int false "每頁數量"
// @Success 200 {object} model.BlogArticle "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles [get]
func (t *Article) List(c *gin.Context) {}

// @Summary 創建文章
// @Produce json
// @Param tag_id body int true "標籤ID"
// @Param title body string true "文章標題"
// @Param desc body string false "文章簡介"
// @Param cover_image_url body string true "封面圖片地址"
// @Param content body string true "文章內容"
// @Param created_by body string true "創建者"
// @Param state body int false "狀態"
// @Success 200 {object} model.BlogArticle "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles [post]
func (t *Article) Create(c *gin.Context) {}

// @Summary 更新文章
// @Produce json
// @Param tag_id body int false "標籤ID"
// @Param title body string false "文章標題"
// @Param desc body string false "文章簡介"
// @Param cover_image_url body string false "封面圖片地址"
// @Param content body string false "文章內容"
// @Param modified_by body string true "修改者"
// @Success 200 {object} model.BlogArticle "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [put]
func (t *Article) Update(c *gin.Context) {}

// @Summary 刪除文章
// @Produce  json
// @Param id path int true "文章ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [delete]
func (t *Article) Delete(c *gin.Context) {}
