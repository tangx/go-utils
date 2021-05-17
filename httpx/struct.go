package httpx

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// StructToHashmap 将 struct 包含 `map` tag 的字段及其值提出，并保存到 map[string]string 中
// 使用场景:
//   在做接口验证签名的时候， 常常需要对传入数据进行一定的排列后在进行签名计算
//     example: https://ai.qq.com/doc/auth.shtml
//   使用 map[string]string 传参后在使用 for 循环就很方便
//   为了保留 **注释即文档** 的特性， 因此常常选用 struct 作为参数管理。
func StructToHashmap(v interface{}, params map[string]string) error {

	// 获取对象的反射值
	// valueof 获取反射值
	// indirect 获取真是反射值。 即， 如果对象是指针， 则返回指针指向的内容
	rv := reflect.Indirect(reflect.ValueOf(v))

	// 判断是否为 struct
	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("want a struct, but got a %#v", rv.Kind())
	}

	// 获取对象的 反射类型
	rt := rv.Type()

	/*
		遍历struct 中的所有字段
	*/
	for i := 0; i < rv.NumField(); i++ {
		// 字段信息
		// fv => filed value。 字段值
		fv := rv.Field(i)
		// ft => filed type。 字段类型
		ft := rt.Field(i)

		// 获取字段tag
		tag, ok := ft.Tag.Lookup("map")
		if !ok {
			// 如果没有定义 tag， 则略过
			continue
		}

		// tag 第一个值为 tagname
		name := tagname(tag)
		// 如果 tag 没有定义name，或字段为空，则默认为字段名称
		if name == "" {
			name = strings.ToLower(ft.Name)
		}

		switch fv.Interface().(type) {
		case string:
			params[name] = fv.String()
		case int, int8, int16, int32, int64:
			params[name] = strconv.FormatInt(fv.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			params[name] = strconv.FormatUint(fv.Uint(), 10)
		case float32, float64:
			params[name] = strconv.FormatFloat(fv.Float(), 'g', 10, 64)
		case bool:
			params[name] = strconv.FormatBool(fv.Bool())

			// default:
			// 	return fmt.Errorf("unsupported filed type: %s -> %v", rt.Name(), rv.Type().String())
		}
	}

	return nil
}

func tagname(tag string) string {
	parts := strings.Split(tag, ",")

	if len(parts) > 0 {
		return parts[0]
	}

	return ""
}
