package main

import (
	application "github.com/qz1ch/calc_service/internal"
)

// запуск сервера
func main() {
	app := application.New()
	app.RunServer()
}
