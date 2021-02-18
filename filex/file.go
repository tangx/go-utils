package filex

import (
	"os"
	"time"
)

func Exist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsExist(err) {
		return true
	}

	return false
}

func IsDir(path string) (bool, error) {
	s, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if s.IsDir() {
		return true, nil
	}

	return false, nil
}

func IsFile(path string) (bool, error) {
	ok, err := IsDir(path)
	if err != nil {
		return false, err
	}

	if !ok {
		return true, nil
	}

	return false, nil
}

// 文件时间对比
func TimestampNewer(filepath string, ts time.Time) bool {
	fs, _ := os.Stat(filepath)
	return ts.Local().After(fs.ModTime().Local())
}

func TimestampEqual(filepath string, ts time.Time) bool {
	fs, _ := os.Stat(filepath)
	return ts.Local().Equal(fs.ModTime().Local())
}
