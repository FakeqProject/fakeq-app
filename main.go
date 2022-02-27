package main

import (
	"github.com/FakeqProject/fakeq-app/bind"
	"github.com/FakeqProject/fakeq-app/config"
	"github.com/FakeqProject/fakeq-app/database"
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
	database.InitDatabase(database.SqliteDatabase{
		FilePath: "./fakeq.db",
	})

	debug := true
	w := webview.New(debug)
	defer w.Destroy()

	w.SetTitle(title)
	w.SetSize(width, height, webview.HintNone)

	bind.AllBindCollection(w)

	w.Navigate("http://localhost" + port)
	w.Run()
}
