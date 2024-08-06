package common

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store base64Captcha.Store = RedisStore{}

func MakeCaptcha() (string, string, error) {
	var driver base64Captcha.Driver
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          1,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driver = driverString.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := c.Generate()

	return id, b64s, err
}

func VerifyCaptcha(captchaId string, verifyValue string) bool {
	if store.Verify(captchaId, verifyValue, true) {
		return true
	}
	return false
}
