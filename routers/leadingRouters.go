package routers

import (
	"mygo_demo/controllers/leading"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", leading.DefaultController{}.Index)
		defaultRouters.GET("/thumbnail1", leading.DefaultController{}.Thumbnail1)
		defaultRouters.GET("/thumbnail2", leading.DefaultController{}.Thumbnail2)
		defaultRouters.GET("/qrcode1", leading.DefaultController{}.Qrcode1)
		defaultRouters.GET("/qrcode2", leading.DefaultController{}.Qrcode2)

	}
}
