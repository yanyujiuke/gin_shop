package admin

import (
	"ginShop/common"
	"ginShop/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type FocusController struct {
	BaseController
}

func (con FocusController) Index(c *gin.Context) {
	focusList := []models.Focus{}
	common.DB.Find(&focusList)

	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})
}

func (con FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}

func (con FocusController) DoAdd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	focusType, focusTypeErr := common.StringToInt(c.PostForm("focus_type"))
	link := strings.Trim(c.PostForm("link"), " ")
	sort, sortErr := common.StringToInt(c.PostForm("sort"))
	status, statusErr := common.StringToInt(c.PostForm("status"))

	if focusTypeErr != nil && sortErr != nil || statusErr != nil {
		con.Error(c, "参数错误", "/admin/focus/add")
		return
	}

	// 文件上传
	url, uploadErr := common.UploadImg(c, "focus_img")
	if uploadErr != nil {
		con.Error(c, "文件上传失败", "/admin/focus/add")
		return
	}

	focus := models.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  url,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(common.GetUnix()),
	}

	createErr := common.DB.Create(&focus).Error
	if createErr != nil {
		con.Error(c, "添加轮播图失败", "/admin/focus/add")
		return
	}

	con.Success(c, "添加轮播图成功", "/admin/focus")
}

func (con FocusController) Edit(c *gin.Context) {
	id, _ := common.StringToInt(c.Query("id"))

	focus := models.Focus{Id: id}
	common.DB.Find(&focus)

	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus": focus,
	})
}

func (con FocusController) DoEdit(c *gin.Context) {
	id, _ := common.StringToInt(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	focusType, focusTypeErr := common.StringToInt(c.PostForm("focus_type"))
	link := strings.Trim(c.PostForm("link"), " ")
	sort, sortErr := common.StringToInt(c.PostForm("sort"))
	status, statusErr := common.StringToInt(c.PostForm("status"))

	if focusTypeErr != nil && sortErr != nil || statusErr != nil {
		con.Error(c, "参数错误", "/admin/focus/edit?id="+common.IntToString(id))
		return
	}

	// 文件上传
	url, _ := common.UploadImg(c, "focus_img")

	focus := models.Focus{Id: id}
	common.DB.Find(&focus)

	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	if url != "" {
		focus.FocusImg = url
	}

	saveErr := common.DB.Save(&focus).Error
	if saveErr != nil {
		con.Error(c, "轮播图编辑失败", "/admin/focus/edit?id="+common.IntToString(id))
		return
	}

	con.Success(c, "轮播图编辑成功", "/admin/focus")
}

func (con FocusController) Delete(c *gin.Context) {
	id, _ := common.StringToInt(c.Query("id"))

	focus := models.Focus{Id: id}
	common.DB.Delete(&focus)

	con.Success(c, "删除成功", "/admin/focus")
}
