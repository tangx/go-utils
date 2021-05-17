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

	return structToHashmap(rv, params, "")

}

func structToHashmap(rv reflect.Value, params map[string]string, parent string) error {

	// 获取对象的 反射类型
	rt := rv.Type()

	/*
		遍历struct 中的所有字段
	*/
	for i := 0; i < rv.NumField(); i++ {
		// 字段信息
		// fv => filed value。 字段值
		fv := reflect.Indirect(rv.Field(i))
		// ft => filed type。 字段类型
		ft := rt.Field(i)

		// 获取字段tag
		tag, ok := ft.Tag.Lookup("map")
		if !ok {
			// 如果没有定义 tag， 则略过
			continue
		}

		// tag 第一个值为 tagname
		parts := strings.Split(tag, ",")
		name := tagname(parts)
		// 如果 tag 没有定义name，或字段为空
		// 规则为 parent.filedname
		if name == "" {
			name = ft.Name
		}

		// inline 内联模式, 结构体中的同名 tag 可能覆盖与被覆盖, 取决于其相对位置
		// 否则字段名默认为字段名称与父字段的组合 Parent__Filedname
		inlineSep := `__`
		if !isInline(parts) {
			name = strings.Trim(strings.Join([]string{parent, name}, inlineSep), inlineSep)
		}

		// 如果字段是 struct 结构，则递归循环。
		if fv.Kind() == reflect.Struct {
			structToHashmap(fv, params, name)
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
			// 	return fmt.Errorf("unsupported filed type: %s -> %#v", rt.Name(), rv.Type())
		}
	}

	return nil
}

func tagname(parts []string) string {

	if len(parts) > 0 {
		return parts[0]
	}

	return ""
}

func isInline(parts []string) bool {

	for _, value := range parts {
		if value == "inline" {
			return true
		}
	}

	return false
}
