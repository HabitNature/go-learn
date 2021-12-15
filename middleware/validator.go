package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_t "github.com/go-playground/validator/v10/translations/en"
	zh_t "github.com/go-playground/validator/v10/translations/zh"
	"go-learn/common"
	"reflect"
)

func ValidateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		en := en.New()
		zh := zh.New()

		t := ut.New(zh, zh, en)
		v := binding.Validator.Engine().(*validator.Validate)

		locale := c.GetHeader("Accept-Language")

		if locale == "" {
			locale = "zh"
		}

		tt, _ := t.GetTranslator(locale)

		switch c.GetHeader("Accept-Language") {
		case "en":
			en_t.RegisterDefaultTranslations(v, tt)
			v.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fmt.Sprintf("<%s>", fld.Tag.Get("json"))
			})
		default:
			zh_t.RegisterDefaultTranslations(v, tt)
			v.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fmt.Sprintf("<%s>", fld.Tag.Get("json"))
			})
		}

		c.Set(common.KEY_TRANSLATE, tt)
		c.Set(common.KEY_VALIDATE, v)

		c.Next()
	}
}
