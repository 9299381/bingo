package valid

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"strings"
)

func Handle(m map[string]interface{}, obj interface{}) error {
	//map -> obj,然后在验证y
	//由于mapstructure->转struct时,不支持key下划线,替换下
	mm := make(map[string]interface{})
	for k, v := range m {
		if strings.Contains(k, "_") {
			kk := strings.ReplaceAll(k, "_", "")
			mm[kk] = v
		}
	}
	err := mapstructure.WeakDecode(mm, obj)
	if err != nil {
		return errors.New("request params valid convert error")
	}
	//如何验证嵌套的问题
	valid := Validation{}
	ok, err := valid.Valid(obj)
	if err != nil {
		//验证中出错
		return err
	}
	if !ok {
		//验证结果失败
		var msg string
		for _, err := range valid.Errors {
			m := strings.Join([]string{err.Field, err.Message}, ":")
			msg = strings.Join([]string{m, msg}, ";")
		}
		return errors.New(msg)
	}
	return nil
}
