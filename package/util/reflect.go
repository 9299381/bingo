package util

import (
	"fmt"
	"reflect"
)

func RefCall(obj interface{}, name string, args ...interface{}) (result []reflect.Value, err error) {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	if !isStructPtr(objT) {
		return nil, fmt.Errorf("%T is not struct pointer", obj)
	}
	method := objV.MethodByName(name)
	if len(args) != method.Type().NumIn() {
		return nil, fmt.Errorf(
			"args num error, is %d not %d",
			method.Type().NumIn(),
			len(args),
		)
	}
	in := make([]reflect.Value, len(args))
	for k, v := range args {
		in[k] = reflect.ValueOf(v)
	}
	return method.Call(in), nil
}

func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}
