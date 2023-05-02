package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type CaptchaResult struct {
	Id          string `json:"id"`
	Base64Blob  string `json:"base_64_blob"`
	VerifyValue string `json:"verify_value"`
}

func GenerateCaptcha(c *gin.Context) {
	parameters := base64Captcha.ConfigCharacter{
		Height:             30,
		Width:              60,
		Mode:               3,
		ComplexOfNoiseDot:  0,
		ComplexOfNoiseText: 0,
		IsUseSimpleFont:    true,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         4,
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 254,
		},
	}

	id, captchaInterface := base64Captcha.GenerateCaptcha("", parameters)
	base64Encoding := base64Captcha.CaptchaWriteToBase64Encoding(captchaInterface)

	Ok(c, CaptchaResult{
		Id:         id,
		Base64Blob: base64Encoding,
	})
}

func VerifyCaptcha(id string, value string) bool {
	return base64Captcha.VerifyCaptcha(id, value)
}
