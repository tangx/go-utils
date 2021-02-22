package filex

import (
	"io/ioutil"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func Walkdir(dirpath string, recursive bool) ([]string, error) {
	fsList, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return nil, err
	}

	zipfiles := []string{}
	for _, fs := range fsList {
		subpath := filepath.Join(dirpath, fs.Name())

		if fs.IsDir() {
			logrus.Debugf("walking dir: %s", subpath)

			if recursive {
				files, err := Walkdir(subpath, true)
				if err != nil {
					logrus.Fatal(err)
				}

				zipfiles = append(zipfiles, files...)
			}

			continue
		}

		logrus.Debugf("walking file: %s", subpath)
		zipfiles = append(zipfiles, filepath.Join(dirpath, fs.Name()))
	}

	return zipfiles, nil
}
