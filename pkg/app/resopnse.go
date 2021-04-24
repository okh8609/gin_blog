package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okh8609/gin_blog/pkg/errcode"
)

type GResponse struct {
	gc *gin.Context
}

func NewGResponse(c *gin.Context) *GResponse {
	return &GResponse{
		gc: c,
	}
}

func (r *GResponse) SendOkResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.gc.JSON(http.StatusOK, data)
}

func (r *GResponse) SendErrResponse(e *errcode.Error) {
	response := gin.H{"code": e.GetCode(), "msg": e.GetMsg()}
	details := e.GetDetails()
	if len(details) > 0 {
		response["details"] = details
	}
	r.gc.JSON(e.GetHttpStatusCode(), response)
}

func (r *GResponse) SendOkResponseList(list interface{}, totalRows int64) {
	r.gc.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.gc),
			PageSize:  GetPageSize(r.gc),
			TotalRows: totalRows,
		},
	})
}
