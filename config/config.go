package config

import "github.com/spf13/viper" //  viper — помощник для чтения конфигов

type Config struct {
	Server struct {
		Port string //Здесь будет храниться порт из YAML
	}
	Database struct {
		URL string // Адрес базы данных
	}
	// ... остальные поля
}

func Load() *Config {
	viper.AddConfigPath("config") //  Где искать config.yaml
	viper.SetConfigName("config") //  Имя файла без расширения
	viper.ReadInConfig()          //  Читаем файл

	var cfg Config
	viper.Unmarshal(&cfg) //  Конвертируем YAML в структуру Go
	return &cfg
}
