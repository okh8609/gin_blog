package api

import (
	"github.com/gin-gonic/gin"
	"github.com/okh8609/gin_blog/global"
	"github.com/okh8609/gin_blog/internal/service"
	"github.com/okh8609/gin_blog/pkg/app"
	"github.com/okh8609/gin_blog/pkg/errcode"
	"github.com/okh8609/gin_blog/pkg/utils"
)

// @Summary 新增使用者
// @Produce  json
// @Param uuid body string true "UUID或使用者名稱" maxlength(255)
// @Param password body string true "密碼" maxlength(255)
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /auth [post]
func CreateAuth(c *gin.Context) {
	param := service.CreateAuthParam{}
	response := app.NewGResponse(c)
	valid, errs := utils.BindAndValid(c, &param)
	if !valid {
		global.MyLogger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.SendErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateAuth(&param)
	if err != nil {
		global.MyLogger.Errorf(c, "svc.CreateAuth err: %v", err)
		response.SendErrResponse(errcode.ErrorCreateAuthFail)
		return
	}

	response.SendOkResponse(gin.H{})
}

// @Summary 驗證使用者
// @Produce  json
// @Param uuid body string true "UUID或使用者名稱" maxlength(255)
// @Param password body string true "密碼" maxlength(255)
// @Success 200 {object} string "成功。 回傳JWT token"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /auth/verify [post]
func VerifyAuth(c *gin.Context) {
	param := service.VerifyAuthParam{
		UUID:     c.Param("uuid"),
		Password: c.Param("password"),
	}
	response := app.NewGResponse(c)
	valid, errs := utils.BindAndValid(c, &param)
	if !valid {
		global.MyLogger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.SendErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	token, err := svc.VerifyAuth(&param)
	if err != nil {
		global.MyLogger.Errorf(c, "svc.VerifyAuth err: %v", err)
		response.SendErrResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	if token == "" {
		global.MyLogger.Error(c, "app.GenerateToken err: ... ")
		response.SendErrResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.SendOkResponse(gin.H{
		"token": token,
	})
}

// @Summary 更新使用者
// @Produce  json
// @Param uuid body string true "UUID或使用者名稱" maxlength(255)
// @Param password body string true "舊密碼" maxlength(255)
// @Param new_password body string true "新密碼" maxlength(255)
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /auth [put]
func UpdateAuth(c *gin.Context) {
	param := service.UpdateAuthParam{
		UUID:        c.Param("uuid"),
		Password:    c.Param("password"),
		NewPassword: c.Param("new_password"),
	}
	response := app.NewGResponse(c)
	valid, errs := utils.BindAndValid(c, &param)
	if !valid {
		global.MyLogger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.SendErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateAuth(&param)
	if err != nil {
		global.MyLogger.Errorf(c, "svc.UpdateAuth err: %v", err)
		response.SendErrResponse(errcode.ErrorUpdateAuthFail)
		return
	}

	response.SendOkResponse(gin.H{})
}

// @Summary 刪除使用者
// @Produce  json
// @Param uuid body string true "UUID或使用者名稱" maxlength(255)
// @Param password body string true "密碼" maxlength(255)
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /auth [delete]
func DeleteAuth(c *gin.Context) {
	param := service.DeleteAuthParam{
		UUID:     c.Param("uuid"),
		Password: c.Param("password"),
	}
	response := app.NewGResponse(c)
	valid, errs := utils.BindAndValid(c, &param)
	if !valid {
		global.MyLogger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.SendErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteAuth(&param)
	if err != nil {
		global.MyLogger.Errorf(c, "svc.DeleteAuth err: %v", err)
		response.SendErrResponse(errcode.ErrorDeleteAuthFail)
		return
	}

	response.SendOkResponse(gin.H{})
}
