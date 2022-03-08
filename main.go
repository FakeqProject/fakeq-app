package main

import (
	"log"

	"github.com/FakeqProject/fakeq-app/bind"
	"github.com/FakeqProject/fakeq-app/config"
	"github.com/FakeqProject/fakeq-app/database"
	"github.com/spf13/viper"
	"github.com/webview/webview"
)

const (
	title  = "Fakeq"
	width  = 1024
	height = 768
	port   = ":8080"
)

func main() {
	config.InitConfig("config", "yaml", "./config")
	setupDatabase()

	debug := true
	w := webview.New(debug)
	defer w.Destroy()

	w.SetTitle(title)
	w.SetSize(width, height, webview.HintNone)

	bind.AllBindCollection(w)

	w.Navigate("http://localhost" + port)
	w.Run()
}

func setupDatabase() {
	databaseType := viper.GetString("database.type")
	switch databaseType {
	case "mysql":
		database.InitDatabase(database.MysqlDatabase{
			UserName: viper.GetString("mysql.username"),
			Password: viper.GetString("mysql.password"),
			Host:     viper.GetString("mysql.host"),
			Port:     viper.GetString("mysql.port"),
			Database: viper.GetString("mysql.database"),
			CharSet:  viper.GetString("mysql.charset"),
		})
	case "sqlite":
		database.InitDatabase(database.SqliteDatabase{
			FilePath: viper.GetString("sqlite.path"),
		})
	default:
		log.Fatal("Unsupported database type: ", databaseType)
	}
}
