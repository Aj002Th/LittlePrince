package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// use in router [0-100]
func GetPageParams(ctx *gin.Context) (page int, pageSize int) {
	pageStr, ok := ctx.GetQuery("page")
	if !ok {
		page = 1
	} else {
		page, _ = strconv.Atoi(pageStr)
		if page == 0 {
			page = 1
		}
	}

	pageSizeStr, ok := ctx.GetQuery("page_size")
	if !ok {
		pageSize = 10
	} else {
		pageSize, _ = strconv.Atoi(pageSizeStr)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
	}
	return
}
