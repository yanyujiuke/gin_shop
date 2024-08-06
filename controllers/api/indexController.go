package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct{}

func (IndexController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "Success",
	})
}
