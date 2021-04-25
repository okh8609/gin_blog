package api

import (
	"github.com/gin-gonic/gin"
	"github.com/okh8609/gin_blog/global"
	"github.com/okh8609/gin_blog/internal/service"
	"github.com/okh8609/gin_blog/pkg/app"
	"github.com/okh8609/gin_blog/pkg/errcode"
	"github.com/okh8609/gin_blog/pkg/upload"
	"github.com/okh8609/gin_blog/pkg/utils"
)

// @Summary 上傳檔案
// @Produce  json
// @Param file body string true "檔案路徑"
// @Param type body int true "檔案類型[1:ImageFile(.jpg .gif .png), 2:DocFile(.pdf)]" Enums(1, 2)
// @Success 200 {object} string "成功  {"file_access_url": XXX}  "
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /upload/file [post]
func UploadFile(c *gin.Context) {
	response := app.NewGResponse(c)

	fileType := utils.StrMust2Int(c.PostForm("type"))
	if fileType <= 0 {
		response.SendErrResponse(errcode.InvalidParams)
		return
	}

	file, fileHeader, err := c.Request.FormFile("file") // 上傳的檔案的POST的參數名稱
	if err != nil {
		response.SendErrResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	if fileHeader == nil {
		response.SendErrResponse(errcode.InvalidParams)
		return
	}

	ss := service.New(c.Request.Context())
	fileInfo, err := ss.UploadFile(file, fileHeader, upload.FileType(fileType))
	if err != nil {
		global.MyLogger.Errorf(c, "service.UploadFile err: %v", err)
		response.SendErrResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.SendOkResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
