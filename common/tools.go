package common

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path"
	"strconv"
)

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 表示把string转换成int
func StringToInt(str string) (int, error) {
	n, err := strconv.Atoi(str)
	return n, err
}

// 表示把int转换成string
func IntToString(n int) string {
	str := strconv.Itoa(n)
	return str
}

func UploadImg(c *gin.Context, imgFieldName string) (string, error) {
	file, fielErr := c.FormFile(imgFieldName)
	if fielErr != nil {
		return "", fielErr
	}

	// 获取文件后缀名
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}

	dir := "./static/upload/" + GetDay()
	mkdirErr := os.MkdirAll(dir, 0666)
	if mkdirErr != nil {
		return "", mkdirErr
	}

	fileName := strconv.FormatInt(GetUnix(), 10) + "_" + file.Filename

	dstUrl := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dstUrl)

	return dstUrl, nil
}
