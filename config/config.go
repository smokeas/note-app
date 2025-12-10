package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config — финальная структурa конфигурации, используемая в main
type Config struct {
	Server struct {
		Port string
	}
	DB struct {
		DSN string
	}
	JWT struct {
		Secret string
		TTL    time.Duration // хранится как Duration в коде
	}
}

// Load загружает config/config.yaml (или переменные окружения) и возвращает *Config
func Load() *Config {
	viper.AddConfigPath("config") // ищем папку config/
	viper.SetConfigName("config") // ищем config.yaml
	viper.SetConfigType("yaml")

	// Поддержка чтения из окружения: например, JWT_TTL или SERVER_PORT
	viper.AutomaticEnv()

	// Попробуем прочитать файл (если нет файла, viper не упадёт — используем значения по умолчанию)
	_ = viper.ReadInConfig() // игнорируем ошибку — допустимо использовать env или defaults

	// Значения по умолчанию (если нужно)
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("jwt.ttl_minutes", 60) // TTL в минутах по умолчанию

	cfg := &Config{}

	// Server
	cfg.Server.Port = viper.GetString("server.port")

	// Database DSN: поддерживаем ключи database.dsn или database.url в конфиге (legacy)
	dsn := viper.GetString("database.dsn")
	if dsn == "" {
		dsn = viper.GetString("database.url")
	}
	cfg.DB.DSN = dsn

	// JWT secret
	cfg.JWT.Secret = viper.GetString("jwt.secret")

	// TTL: ожидаем в конфиге целое число минут (jwt.ttl_minutes) или jwt.ttl
	ttlMinutes := viper.GetInt("jwt.ttl_minutes")
	if ttlMinutes == 0 {
		// fallback: возможно пользователь положил jwt.ttl (int minutes)
		ttlMinutes = viper.GetInt("jwt.ttl")
	}
	if ttlMinutes == 0 {
		// ещё fallback: если строка в jwt.ttl_string (например "1h30m")
		if s := viper.GetString("jwt.ttl_string"); s != "" {
			if d, err := time.ParseDuration(s); err == nil {
				cfg.JWT.TTL = d
			}
		}
	}
	// если cfg.JWT.TTL не установлен строкой, установим из minutes
	if cfg.JWT.TTL == 0 {
		cfg.JWT.TTL = time.Duration(ttlMinutes) * time.Minute
	}

	return cfg
}
