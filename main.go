package main

import (
	"github.com/FakeqProject/fakeq-app/config"
	"github.com/FakeqProject/fakeq-app/database"
)

func main() {
	config.InitConfig("config", "yaml", "./config")
	database.InitDatabase(database.SqliteDatabase{
		FilePath: "./fakeq.db",
	})
}
