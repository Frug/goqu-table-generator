package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config stores configuration values.
type Config struct {
	DBHost     string
	DBUser     string
	DBPass     string
	DBPort     string
	DBName     string
	DBUseSSL   bool
	OutputPath string
	Verbose    bool
}

// NewFromEnv returns a Config populated from mapstructureironment variables.
func NewFromEnv(v *viper.Viper) (Config, error) {

	v.AutomaticEnv()
	// working around viper issue #761 (https://github.com/spf13/viper/issues/761)
	v.BindEnv("DBHost", "DB_HOST")
	v.BindEnv("DBUser", "DB_USER")
	v.BindEnv("DBPass", "DB_PASS")
	v.BindEnv("DBPort", "DB_PORT")
	v.BindEnv("DBName", "DB_NAME")
	v.BindEnv("DBUseSSL", "DB_USE_SSL")
	v.BindEnv("Verbose", "VERBOSE")

	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		return c, err
	}

	if c.DBHost == "" {
		fmt.Printf("DB_HOST must be set\n")
		os.Exit(1)
	}
	if c.DBUser == "" {
		fmt.Printf("DB_USER must be set\n")
		os.Exit(1)
	}

	return c, err
}
