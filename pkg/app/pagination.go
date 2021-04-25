package app

import "github.com/gin-gonic/gin"
import "github.com/okh8609/gin_blog/pkg/utils"
import "github.com/okh8609/gin_blog/global"

type Pager struct {
	// 页码
	Page int `json:"page"`
	// 每页数量
	PageSize int `json:"page_size"`
	// 总行数
	TotalRows int64 `json:"total_rows"`
}

func GetPage(c *gin.Context) int { // 第幾頁?
	page := utils.StrMust2Int(c.Query("page"))
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int { // 每頁有幾筆資料?
	pageSize := utils.StrMust2Int(c.Query("page_size"))
	if pageSize <= 0 {
		return global.App.DefaultPageSize
	}
	if pageSize > global.App.MaxPageSize {
		return global.App.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page int, pageSize int) int { // 跳過前面幾筆資料呢?
	
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
