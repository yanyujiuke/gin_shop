package admin

import (
	"ginShop/common"
	"ginShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GoodsCateController struct {
	BaseController
}

func (con GoodsCateController) Index(c *gin.Context) {
	goodsCateList := []models.GoodsCate{}
	common.DB.Where("pid = ?", 0).Preload("GoodsCateItems").Find(&goodsCateList)

	c.HTML(http.StatusOK, "admin/goodsCate/index.html", gin.H{
		"goodsCateList": goodsCateList,
	})
}

func (con GoodsCateController) Add(c *gin.Context) {
	// 获取顶级分类
	goodsCateList := []models.GoodsCate{}

	common.DB.Where("pid = ?", 0).Find(&goodsCateList)

	c.HTML(http.StatusOK, "admin/goodsCate/add.html", gin.H{
		"goodsCateList": goodsCateList,
	})
}

func (con GoodsCateController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	pid, err1 := common.StringToInt(c.PostForm("pid"))
	link := c.PostForm("link")
	template := c.PostForm("template")
	subTitle := c.PostForm("sub_title")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sort, err2 := common.StringToInt(c.PostForm("sort"))
	status, err3 := common.StringToInt(c.PostForm("status"))

	if err1 != nil || err3 != nil {
		con.Error(c, "传入参数类型不正确", "/goodsCate/add")
		return
	}
	if err2 != nil {
		con.Error(c, "排序值必须是整数", "/goodsCate/add")
		return
	}
	cateImgDir, _ := common.UploadImg(c, "cate_img")
	goodsCate := models.GoodsCate{
		Title:       title,
		Pid:         pid,
		SubTitle:    subTitle,
		Link:        link,
		Template:    template,
		Keywords:    keywords,
		Description: description,
		CateImg:     cateImgDir,
		Sort:        sort,
		AddTime:     int(common.GetUnix()),
		Status:      status,
	}
	err := common.DB.Create(&goodsCate).Error
	if err != nil {
		con.Error(c, "增加数据失败", "/admin/goodsCate/add")
		return
	}
	con.Success(c, "增加数据成功", "/admin/goodsCate")
}

func (con GoodsCateController) Edit(c *gin.Context) {
	id, _ := common.StringToInt(c.Query("id"))

	goodsCate := models.GoodsCate{Id: id}
	err := common.DB.Find(&goodsCate).Error
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goodsCate")
		return
	}

	goodsCateList := []models.GoodsCate{}
	common.DB.Where("pid = ?", 0).Find(&goodsCateList)

	c.HTML(http.StatusOK, "admin/goodsCate/edit.html", gin.H{
		"goodsCate":     goodsCate,
		"goodsCateList": goodsCateList,
	})
}

func (con GoodsCateController) DoEdit(c *gin.Context) {
	id, err1 := common.StringToInt(c.PostForm("id"))
	title := c.PostForm("title")
	pid, err2 := common.StringToInt(c.PostForm("pid"))
	link := c.PostForm("link")
	template := c.PostForm("template")
	subTitle := c.PostForm("sub_title")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sort, err3 := common.StringToInt(c.PostForm("sort"))
	status, err4 := common.StringToInt(c.PostForm("status"))

	if err1 != nil || err2 != nil || err4 != nil {
		con.Error(c, "传入参数类型不正确", "/goodsCate/add")
		return
	}
	if err3 != nil {
		con.Error(c, "排序值必须是整数", "/goodsCate/add")
		return
	}
	cateImgDir, _ := common.UploadImg(c, "cate_img")

	goodsCate := models.GoodsCate{Id: id}
	common.DB.Find(&goodsCate)
	goodsCate.Title = title
	goodsCate.Pid = pid
	goodsCate.Link = link
	goodsCate.Template = template
	goodsCate.SubTitle = subTitle
	goodsCate.Keywords = keywords
	goodsCate.Description = description
	goodsCate.Sort = sort
	goodsCate.Status = status
	if cateImgDir != "" {
		goodsCate.CateImg = cateImgDir
	}
	err := common.DB.Save(&goodsCate).Error
	if err != nil {
		con.Error(c, "修改失败", "/admin/goodsCate/edit?id="+common.IntToString(id))
		return
	}
	con.Success(c, "修改成功", "/admin/goodsCate")
}

func (con GoodsCateController) Delete(c *gin.Context) {
	id, err := common.StringToInt(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goodsCate")
		return
	}

	goodsCateList := []models.GoodsCate{}
	common.DB.Where("pid = ?", id).Find(&goodsCateList)
	if len(goodsCateList) > 0 {
		con.Error(c, "当前分类下还有子分类，不可删除", "/admin/goodsCate")
		return
	}

	goodsCate := models.GoodsCate{Id: id}
	common.DB.Delete(&goodsCate)
	con.Success(c, "删除成功", "/admin/goodsCate")
}
