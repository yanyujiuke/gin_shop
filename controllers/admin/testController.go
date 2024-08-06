package admin

import (
	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
	"io/ioutil"
)

type TestController struct {
	BaseController
}

func (con TestController) Index(c *gin.Context) {
	savepath := "static/upload/qrcode.png"
	err := qrcode.WriteFile("https://www.16type.com", qrcode.Medium, 556, savepath)
	if err != nil {
		c.String(200, "生成二维码失败")
		return
	}
	file, _ := ioutil.ReadFile(savepath)
	c.String(200, string(file))
}
