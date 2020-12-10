package md5sha

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"time"
)

func File(filename string) (sha string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		// log.Fatal(err)
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		// log.Fatal(err)
		return "", nil
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func MustFile(filename string) (sha string) {
	sha, err := File(filename)
	if err != nil {
		panic(err)
	}
	return
}

func String(str string) (sha string, err error) {

	h := md5.New()
	_, err = io.WriteString(h, str)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil

}

func MustString(str string) (sha string) {
	sha, err := String(str)
	if err != nil {
		panic(err)
	}
	return
}

// Time return current time md5sha
func Time() (sha string) {
	ts := time.Now().UnixNano()
	return MustString(fmt.Sprint(ts))
}
