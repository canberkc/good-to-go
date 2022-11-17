package configs

import "github.com/spf13/viper"

var Configuration Config

type Config struct {
	Port         string `mapstructure:"PORT"`
	CertEndpoint string `mapstructure:"CERT_ENDPOINT"`
}

func LoadConfig(paths []string) error {
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	viper.SetConfigName("environment")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Configuration)
	if err != nil {
		return err
	}

	return nil
}
