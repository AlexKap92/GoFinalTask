package configure

import (
	"log"

	"github.com/spf13/viper"
)

type Configs struct {
	DSN           string `yaml:"dsn"`
	ServiceAPIKey string `yaml:"serviceapikey"`
	BaseCurrency  string `yaml:"basecurrency"`
}

func GetConfigs(path string) (*Configs, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	log.Println("\nUsing config file:", viper.ConfigFileUsed())
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var appsettings Configs
	if err := viper.Unmarshal(&appsettings); err != nil {
		log.Fatalf("Error unmarshalling appsettings: %s", err)
		return nil, err
	}
	return &appsettings, nil
}
