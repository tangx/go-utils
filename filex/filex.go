package filex

import (
	"os"
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
