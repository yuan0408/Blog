package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/yuan0408/go-gin-example/models"
	"github.com/yuan0408/go-gin-example/pkg/e"
	"github.com/yuan0408/go-gin-example/pkg/util"
	//"log"
	log "github.com/yuan0408/go-gin-example/pkg/logging"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	valid := validation.Validation{}
	username := c.Query("username")
	password := c.Query("password")
	a := auth{username, password}
	ok, _ := valid.Valid(&a)

	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	if ok {
		if models.CheckAuth(username, password) {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			}
			data["token"] = token
			code = e.SUCCESS
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Info(err.Key, err.Value)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    data,
	})
}
