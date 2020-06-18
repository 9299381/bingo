package util

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	elem := reflect.ValueOf(obj).Elem()
	relType := elem.Type()

	var data = make(map[string]interface{})
	for i := 0; i < relType.NumField(); i++ {
		data[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	return data
}
func Struct2Struct(src, dst interface{}) error {
	m := Struct2Map(src)
	return Map2Struct(m, dst)
}

func Map2Struct(m, st interface{}) error {
	// map转struct时, key与struct的字段应该相同,忽略大小写,
	// 注意 不可随意增加_,json中可以有_
	mm, ok := m.(map[string]interface{})
	if ok == false {
		return errors.New("Map2Struct Convert Error")
	}
	for k, v := range mm {
		if strings.Contains(k, "_") {
			kk := strings.ReplaceAll(k, "_", "")
			mm[kk] = v
		}
	}
	err := mapstructure.WeakDecode(mm, st)
	if err != nil {
		return errors.New("Map2Struct Convert Error")
	}

	return nil
}

func FormEncode(params map[string]interface{}) url.Values {
	data := url.Values{}
	for k, param := range params {
		paramsType := reflect.TypeOf(param)
		switch paramsType.String() {
		case "string":
			data.Set(k, param.(string))
		case "int":
			data.Set(k, strconv.Itoa(param.(int)))
		default:
			str, _ := json.Marshal(param)
			data.Set(k, string(str))

		}
	}
	return data
}
