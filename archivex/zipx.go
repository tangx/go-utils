package zipx

import (
	"archive/zip"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tangx/goutils/filex"
)

func ZipFiles(name string, files ...string) error {
	fobj, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fobj.Close()

	zw := zip.NewWriter(fobj)
	defer zw.Close()

	for _, file := range files {
		logrus.Debugf("Adding %s", file)
		fobj, err := os.Open(file)
		if err != nil {
			logrus.Error(err)
			continue
		}

		fi, err := fobj.Stat()
		if err != nil {

			logrus.Error(err)
			return err
		}

		// todo: 需要实现目录结构
		fh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}

		w, err := zw.CreateHeader(fh)
		if err != nil {
			logrus.Error(err)
			return err
		}

		_, err = io.Copy(w, fobj)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	return nil
}

func ZipDir(name string, dirpath string) error {
	files, err := filex.Walkdir(dirpath, false)
	if err != nil {
		return err
	}

	return ZipFiles(name, files...)
}

func ZipDirRecursive(name string, dirpath string) error {

	files, err := filex.Walkdir(dirpath, true)
	if err != nil {
		return err
	}

	return ZipFiles(name, files...)

}
