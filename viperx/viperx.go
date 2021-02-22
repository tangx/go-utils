package viperx

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	configTypes = []string{"yaml", "yml"}
	configPaths = []string{"."}
)

func Default() {
	SetConfigName("config")
	AddConfigTypes(configTypes...)
	AddConfigPaths(configPaths...)
}

func AddConfigTypes(types ...string) {
	for _, tpe := range types {
		viper.SetConfigType(tpe)
	}
}

func AddConfigPaths(paths ...string) {
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
}

func SetConfigName(name string) {
	viper.SetConfigName(name)
}

func MustReadInConfig(paths ...string) {
	err := ReadInConfig(paths...)
	if err != nil {
		panic(err)
	}
}

func ReadInConfig(paths ...string) error {
	AddConfigPaths(paths...)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Warning(err.Error())
		return err
	}
	return nil
}
