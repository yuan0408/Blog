package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/yuan0408/go-gin-example/pkg/setting"
)

func GetPage(c *gin.Context) int {
	res := 0
	page, _ := com.StrTo(c.Query("Page")).Int()
	if page > 0 {
		res = (page - 1) * setting.PageSize
	}
	return res
}
