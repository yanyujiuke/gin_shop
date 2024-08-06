package admin

import (
	"ginShop/common"
	"ginShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type GoodsTypeController struct {
	BaseController
}

func (con GoodsTypeController) Index(c *gin.Context) {
	goodsTypeList := []models.GoodsType{}
	common.DB.Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goodsType/index.html", gin.H{
		"goodsTypeList": goodsTypeList,
	})
}

func (con GoodsTypeController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goodsType/add.html", gin.H{})
}

func (con GoodsTypeController) DoAdd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	status, _ := common.StringToInt(c.PostForm("status"))

	goodsType := models.GoodsType{}
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status
	goodsType.AddTime = int(common.GetUnix())

	err := common.DB.Create(&goodsType).Error
	if err != nil {
		con.Error(c, "添加失败", "/admin/goodsType/add")
	} else {
		con.Success(c, "添加成功", "/admin/goodsType")
	}
}

func (con GoodsTypeController) Edit(c *gin.Context) {
	id, _ := common.StringToInt(c.Query("id"))

	goodsType := models.GoodsType{Id: id}
	err := common.DB.Find(&goodsType).Error
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goodsType")
		return
	}

	c.HTML(http.StatusOK, "admin/goodsType/edit.html", gin.H{
		"goodsType": goodsType,
	})
}

func (con GoodsTypeController) DoEdit(c *gin.Context) {
	id, _ := common.StringToInt(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	status, _ := common.StringToInt(c.PostForm("status"))

	goodsType := models.GoodsType{Id: id}
	common.DB.Find(&goodsType)
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status

	err := common.DB.Save(&goodsType).Error
	if err != nil {
		con.Error(c, "编辑失败", "/admin/goodsType")
	} else {
		con.Success(c, "编辑成功", "/admin/goodsType")
	}
}

func (con GoodsTypeController) Delete(c *gin.Context) {
	id, err := common.StringToInt(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goodsType")
	} else {
		goodsType := models.GoodsType{Id: id}
		common.DB.Delete(&goodsType)
		con.Success(c, "删除成功", "/admin/goodsType")
	}
}
