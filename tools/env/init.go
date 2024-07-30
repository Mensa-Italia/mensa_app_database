package env

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type config struct {
	PasswordUUID string `env:"PASSWORD_UUID" envDefault:""`
	PasswordSalt string `env:"PASSWORD_SALT" envDefault:""`
}

var cfg = config{}

func init() {
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func GetPasswordUUID() string {
	return cfg.PasswordUUID
}

func GetPasswordSalt() string {
	return cfg.PasswordSalt
}
