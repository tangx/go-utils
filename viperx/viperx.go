package viperx

import "github.com/spf13/viper"

var (
	configTypes = []string{"yaml", "yml"}
	configPaths = []string{"."}
)

func Default() {
	SetConfigName("config")
	AddTypes(configTypes...)
	AddConfigPath(configPaths...)
}

func AddTypes(types ...string) {
	for _, tpe := range types {
		viper.SetConfigType(tpe)
	}
}

func AddConfigPath(paths ...string) {
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
}

func SetConfigName(name string) {
	viper.SetConfigName(name)
}

func MustReadInConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
