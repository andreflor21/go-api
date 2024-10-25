package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBDriver string
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
	WebServerPort string
	JwtSecret string
	JwtExpiresIn int
	TokenAuth *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JwtSecret), nil)
	
	return cfg, nil
}