package application

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"reflect"
)

type AppConfiguration struct {
	N           int           `yaml:"n" env:"WHATWW_N" env-default:"50000"`
	Goal        int           `yaml:"goal" env:"WHATWW_GOAL" env-default:"6"`
	Sectors     []SectorSetup `yaml:"sectors" env:"WHATWW_SECTORS" env-default:""`
	PlayedTurns []GameTurn    `yaml:"played_turns" env:"WHATWW_PLAYED_TURNS" env-default:""`
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

	if config.Sectors == nil || len(config.Sectors) == 0 {
		config.Sectors = GetDefaultSectorSetups()
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
