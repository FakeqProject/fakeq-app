package config

import (
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	CharSet  string `json:"charset"`
}

type SqliteConfig struct {
	FilePath string `json:"filepath"`
}

// InitConfig init viper config
func InitConfig(fileName, fileType, filePath string) error {
	// config file
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	// config file path
	viper.AddConfigPath(filePath)
	// try to load config file
	err := viper.ReadInConfig()
	return err
}

// GetMysqlConfig get mysql config
func GetMysqlConfig() MysqlConfig {
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	charset := viper.GetString("mysql.charset")

	return MysqlConfig{
		UserName: username,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
		CharSet:  charset,
	}
}

// GetSqliteConfig get sqlite config
func GetSqliteConfig() SqliteConfig {
	filePath := viper.GetString("sqlite.filepath")

	return SqliteConfig{
		FilePath: filePath,
	}
}
