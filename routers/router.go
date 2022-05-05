package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuan0408/go-gin-example/pkg/setting"
	v1 "github.com/yuan0408/go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//retrieve
		apiv1.GET("/tags", v1.GetTags)
		//create
		apiv1.POST("/tags", v1.AddTag)
		//update
		apiv1.PUT("/tags/:id", v1.EditTag)
		//delete
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}
