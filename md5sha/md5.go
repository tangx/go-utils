package md5sha

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func File(filename string) (s string, err error) {
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

func String(str string) (sha string, err error) {

	h := md5.New()
	_, err = io.WriteString(h, str)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil

}
