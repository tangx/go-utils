package regex

import (
	"fmt"
	"regexp"
)

// Regex 主要用于正则表达式 **命名字段** 的匹配和获取
// 参考 https://stackoverflow.com/a/20751656
type Regex struct {
	re   *regexp.Regexp
	smap map[string]string
	bmap map[string][]byte
}

func MustNew(expr string) *Regex {
	return &Regex{
		re:   regexp.MustCompile(expr),
		smap: make(map[string]string),
		bmap: make(map[string][]byte),
	}
}
func New(expr string) (*Regex, error) {
	re, err := regexp.Compile(expr)
	if err != nil {
		return nil, fmt.Errorf("regex compile failed %v", err)
	}
	return &Regex{re: re}, nil
}

// FindStringNamedSubmatch 匹配并获取命名字符串
func (r *Regex) FindStringNamedSubmatch(s string) *Regex {
	result := make(map[string]string)

	// FindStringSubmatch 返回匹配信息数组
	match := r.re.FindStringSubmatch(s)

	// SubexpNames 获取返回 named 字段
	for i, name := range r.re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	r.smap = result

	return r
}

// FindNamedSubmatch 匹配并获取命名 byte
func (r *Regex) FindNamedSubmatch(b []byte) *Regex {

	// FindStringSubmatch 返回匹配信息数组
	match := r.re.FindSubmatch(b)

	// SubexpNames 获取返回 named 字段
	for i, name := range r.re.SubexpNames() {
		if i != 0 && name != "" {
			r.bmap[name] = match[i]
		}
	}

	return r
}

// Get return byte by name
// return an empty []byte slice if name does not exist
func (r *Regex) Get(name string) []byte {
	return r.bmap[name]
}

// GetString return string by name
// return an empty string if name does not exist
func (r *Regex) GetString(name string) string {
	return r.smap[name]
}

// Lookup return a couple of []byte and bool
func (r *Regex) Lookup(name string) ([]byte, bool) {
	data, ok := r.bmap[name]
	return data, ok
}

// LookupString return a couple of string and bool
func (r *Regex) LookupString(name string) (string, bool) {
	data, ok := r.smap[name]
	return data, ok
}
