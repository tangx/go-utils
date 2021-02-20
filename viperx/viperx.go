package viperx

import "github.com/spf13/viper"

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
	AddConfigPaths(paths...)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
