package routers

import (
	"ginShop/controllers/api"
	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/api")
	{
		adminRouters.GET("/", api.IndexController{}.Index)

	}
}
