package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Listen, DatabaseURL string
	ReminderInterval    time.Duration
}

func Load() Config {
	v := viper.New()
	v.SetEnvPrefix("CLASSROOM")
	v.AutomaticEnv()
	v.SetDefault("LISTEN", ":8080")
	v.SetDefault("REMINDER_INTERVAL", "1m")
	return Config{v.GetString("LISTEN"), v.GetString("DATABASE_URL"), v.GetDuration("REMINDER_INTERVAL")}
}
