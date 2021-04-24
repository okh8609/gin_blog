package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okh8609/gin_blog/global"
	"github.com/okh8609/gin_blog/internal/service"
	"github.com/okh8609/gin_blog/pkg/errcode"
	"github.com/okh8609/gin_blog/pkg/validation"
)

type Tag struct{}

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) Get(c *gin.Context) {}

// @Summary 取得多個標籤
// @Produce  json
// @Param name query string false "標籤名稱" maxlength(100)
// @Param state query int false "狀態" Enums(0, 1) default(1)
// @Param page query int false "頁碼"
// @Param page_size query int false "每頁數量"
// @Success 200 {object} model.BlogTag "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags [get]
func (t *Tag) List(c *gin.Context) {

	//region --- validation.BindAndValid test ---
	param := service.GetTagsParam{}
	ok, es := validation.BindAndValid(c, &param)
	if !ok {
		global.MyLogger.Errorf(c, "validation.BindAndValid err: %v", es)

		ers := errcode.InvalidParams.WithDetails(es.Errors()...)
		response := gin.H{"code": ers.ErrCode, "msg": ers.GetMsg()}
		details := ers.GetDetails()
		if len(details) > 0 {
			response["details"] = details
		}
		c.JSON(ers.GetHttpStatusCode(), response)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
	//endregion --- validation.BindAndValid test ---

}

// @Summary 新增標籤
// @Produce  json
// @Param name body string true "標籤名稱" minlength(3) maxlength(100)
// @Param state body int false "狀態" Enums(0, 1) default(1)
// @Param created_by body string false "創建者" minlength(3) maxlength(100)
// @Success 200 {object} model.BlogTag "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags [post]
func (t *Tag) Create(c *gin.Context) {}

// @Summary 更新標籤
// @Produce  json
// @Param id path int true "標籤ID"
// @Param name body string false "標籤名稱" minlength(3) maxlength(100)
// @Param state body int false "狀態" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.BlogTag "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags/{id} [put]
func (t *Tag) Update(c *gin.Context) {}

// @Summary 刪除標籤
// @Produce  json
// @Param id path int true "標籤ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags/{id} [delete]
func (t *Tag) Delete(c *gin.Context) {}
