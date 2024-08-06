package admin

import (
	"fmt"
	"ginShop/common"
	"ginShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type GoodsTypeAttributeController struct {
	BaseController
}

func (con GoodsTypeAttributeController) Index(c *gin.Context) {
	typeId, _ := common.StringToInt(c.Query("type_id"))

	goodsTypeAttributeList := []models.GoodsTypeAttribute{}
	common.DB.Where("type_id = ?", typeId).Find(&goodsTypeAttributeList)

	goodsType := models.GoodsType{}
	common.DB.Find(&goodsType)

	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/index.html", gin.H{
		"typeId":                 typeId,
		"goodsTypeAttributeList": goodsTypeAttributeList,
		"goodsType":              goodsType,
	})
}

func (con GoodsTypeAttributeController) Add(c *gin.Context) {
	typeId, _ := common.StringToInt(c.Query("type_id"))

	goodsTypeList := []models.GoodsTypeAttribute{}
	common.DB.Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/add.html", gin.H{
		"typeId":        typeId,
		"goodsTypeList": goodsTypeList,
	})
}

func (con GoodsTypeAttributeController) DoAdd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	typeId, err1 := common.StringToInt(c.PostForm("type_id"))
	attrType, err2 := common.StringToInt(c.PostForm("attr_type"))
	attrValue := c.PostForm("attr_value")
	sort, err3 := common.StringToInt(c.PostForm("sort"))

	fmt.Println(err1, err2)

	if err1 != nil || err2 != nil {
		con.Error(c, "非法请求", "/admin/goodsType")
		return
	}
	if title == "" {
		con.Error(c, "商品类型属性名称不能为空", "/admin/goodsTypeAttribute/add?type_id="+common.IntToString(typeId))
		return
	}

	if err3 != nil {
		con.Error(c, "排序值不对", "/admin/goodsTypeAttribute/add?type_id="+common.IntToString(typeId))
		return
	}

	goodsTypeAttr := models.GoodsTypeAttribute{
		Title:     title,
		TypeId:    typeId,
		AttrType:  attrType,
		AttrValue: attrValue,
		Status:    1,
		Sort:      sort,
		AddTime:   int(common.GetUnix()),
	}
	err := common.DB.Create(&goodsTypeAttr).Error
	if err != nil {
		con.Error(c, "增加商品类型属性失败 请重试", "/admin/goodsTypeAttribute/add?type_id="+common.IntToString(typeId))
	} else {
		con.Success(c, "增加商品类型属性成功", "/admin/goodsTypeAttribute?type_id="+common.IntToString(typeId))
	}
}

func (con GoodsTypeAttributeController) Edit(c *gin.Context) {
	id, _ := common.StringToInt(c.Query("id"))

	goodsTypeAttribute := models.GoodsTypeAttribute{Id: id}
	err := common.DB.Find(&goodsTypeAttribute).Error
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goodsTypeAttribute")
		return
	}

	goodsTypeList := []models.GoodsType{}
	common.DB.Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/edit.html", gin.H{
		"goodsTypeAttribute": goodsTypeAttribute,
		"goodsTypeList":      goodsTypeList,
	})
}

func (con GoodsTypeAttributeController) DoEdit(c *gin.Context) {
	id, err1 := common.StringToInt(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	typeId, err2 := common.StringToInt(c.PostForm("type_id"))
	attrType, err3 := common.StringToInt(c.PostForm("attr_type"))
	attrValue := c.PostForm("attr_value")
	sort, err4 := common.StringToInt(c.PostForm("sort"))

	if err1 != nil || err2 != nil || err3 != nil {
		con.Error(c, "非法请求", "/admin/goodsType")
		return
	}
	if title == "" {
		con.Error(c, "商品类型属性名称不能为空", "/admin/goodsTypeAttribute/edit?id="+common.IntToString(id))
		return
	}
	if err4 != nil {
		con.Error(c, "排序值不对", "/admin/goodsTypeAttribute/edit?id="+common.IntToString(id))
		return
	}

	goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
	common.DB.Find(&goodsTypeAttr)
	goodsTypeAttr.Title = title
	goodsTypeAttr.TypeId = typeId
	goodsTypeAttr.AttrType = attrType
	goodsTypeAttr.AttrValue = attrValue
	goodsTypeAttr.Sort = sort
	err := common.DB.Save(&goodsTypeAttr).Error
	if err != nil {
		con.Error(c, "修改数据失败", "/admin/goodsTypeAttribute/edit?id="+common.IntToString(id))
		return
	}
	con.Success(c, "需改数据成功", "/admin/goodsTypeAttribute?type_id="+common.IntToString(typeId))
}

func (con GoodsTypeAttributeController) Delete(c *gin.Context) {
	id, err := common.StringToInt(c.Query("id"))
	typeId, _ := common.StringToInt(c.Query("type_id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goodsTypeAttribute?type_id="+common.IntToString(typeId))
	} else {
		goodsTypeAttribute := models.GoodsTypeAttribute{Id: id}
		common.DB.Delete(&goodsTypeAttribute)
		con.Success(c, "删除成功", "/admin/goodsTypeAttribute?type_id="+common.IntToString(typeId))
	}
}
