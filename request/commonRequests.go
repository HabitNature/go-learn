package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go-learn/common"
	"strings"
)

func Validate(c *gin.Context, req interface{}) error {
	v, err := getValidator(c)

	if err != nil {
		return err
	}

	t, err := getTranslator(c)

	if err != nil {
		return err
	}

	err = c.ShouldBind(req)

	if err == nil {
		err = v.Struct(req)
	}

	if err != nil {
		errs := err.(validator.ValidationErrors)

		msgs := []string{}

		for _, e := range errs {
			msgs = append(msgs, e.Translate(t))
		}

		return errors.New(strings.Join(msgs, ","))
	}

	return nil
}

func getValidator(c *gin.Context) (*validator.Validate, error) {
	v, exists := c.Get(common.KEY_VALIDATE)

	if !exists {
		return nil, errors.New("未设置验证器")
	}

	vv, ok := v.(*validator.Validate)

	if !ok {
		return nil, errors.New("获取验证器失败")
	}

	return vv, nil
}

func getTranslator(c *gin.Context) (ut.Translator, error) {
	t, exits := c.Get(common.KEY_TRANSLATE)

	if !exits {
		return nil, errors.New("未设置翻译器")
	}

	tt, ok := t.(ut.Translator)

	if !ok {
		return nil, errors.New("获取翻译器失败")
	}

	return tt, nil
}
