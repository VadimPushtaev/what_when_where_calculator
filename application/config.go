package application

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"reflect"
)

type AppConfiguration struct {
}

func NewConfig(configPath *string) *AppConfiguration {
	var config AppConfiguration

	if configPath != nil && *configPath != "" {
		err := cleanenv.ReadConfig(*configPath, &config)
		if err != nil {
			panic(fmt.Sprintf("Couldn't intialize config from file `%s`: %s", *configPath, err))
		}
	} else {
		err := cleanenv.ReadEnv(&config)
		if err != nil {
			panic(fmt.Sprintf("Couldn't intialize config from env: %s", err))
		}
	}

	return &config
}

func (config *AppConfiguration) Print() {
	v := reflect.ValueOf(*config)

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	fmt.Println(values)
}
