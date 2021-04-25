package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/okh8609/gin_blog/global"
	"github.com/okh8609/gin_blog/internal/service"
	"github.com/okh8609/gin_blog/pkg/app"
	"github.com/okh8609/gin_blog/pkg/errcode"
	"github.com/okh8609/gin_blog/pkg/utils"
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

	param := service.GetTagsParam{}
	ok, verrs := utils.BindAndValid(c, &param)
	response := app.NewGResponse(c)
	if !ok {
		global.MyLogger.Errorf(c, "app.BindAndValid errs: %v", verrs.Error())
		response.SendErrResponse(errcode.InvalidParams.WithDetails(verrs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.GetTagsParam{Name: param.Name, State: param.State})
	if err != nil {
		global.MyLogger.Errorf(c, "svc.CountTag err: %v", err)
		response.SendErrResponse(errcode.ErrorCountTagFail)
		return
	}
	tags, err := svc.GetTagList(&param, &pager) // router(路徑) -> service(資料驗證與轉送) -> dao(準備完整的資料) -> model(sql查詢)
	if err != nil {
		global.MyLogger.Errorf(c, "svc.GetTagList err: %v", err)
		response.SendErrResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.SendOkResponseList(tags, totalRows)
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
func (t *Tag) Create(c *gin.Context) {
	param := service.CreateTagParam{}
	response := app.NewGResponse(c)
	valid, errs := utils.BindAndValid(c, &param)
	if !valid {
		global.MyLogger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.SendErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.MyLogger.Errorf(c, "svc.CreateTag err: %v", err)
		response.SendErrResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.SendOkResponse(gin.H{})
}

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
func (t *Tag) Update(c *gin.Context) {
	param := service.UpdateTagParam{ID: utils.StrMust2UInt(c.Param("id"))}
	response := app.NewGResponse(c)
	valid, errs := utils.BindAndValid(c, &param)
	if !valid {
		global.MyLogger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.SendErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.MyLogger.Errorf(c, "svc.UpdateTag err: %v", err)
		response.SendErrResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.SendOkResponse(gin.H{})
}

// @Summary 刪除標籤
// @Produce  json
// @Param id path int true "標籤ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/tags/{id} [delete]
func (t *Tag) Delete(c *gin.Context) {
	param := service.DeleteTagParam{ID: utils.StrMust2UInt(c.Param("id"))}
	response := app.NewGResponse(c)
	valid, errs := utils.BindAndValid(c, &param)
	if !valid {
		global.MyLogger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.SendErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.MyLogger.Errorf(c, "svc.DeleteTag err: %v", err)
		response.SendErrResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.SendOkResponse(gin.H{})
}
