package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GoodsController struct {
	BaseController
}

func (con GoodsController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{})
}

func (con GoodsController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goods/add.html", gin.H{})
}

func (con GoodsController) DoAdd(c *gin.Context) {

}
