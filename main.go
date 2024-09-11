package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	defaultPort = "7541"
	webDir      = "./web"
)

func main() {
	fmt.Println("Запуск сервера")
	// Проверка на существование переменной окружения TODO_PORT
	port := defaultPort
	envPort := os.Getenv("TODO_PORT")
	if envPort != "" {
		port = envPort
	}
	port = ":" + port

	serverFiles := http.FileServer(http.Dir(webDir))
	http.Handle("/", serverFiles)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Остановка сервера")
}
