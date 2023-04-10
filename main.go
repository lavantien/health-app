package main

import (
	"health-app/server"
	"log"
	"net/http"
)

type config struct {
	ServerPort string
	ClientPort string
	ClientPath string
}

func main() {
	setupLogger()
	config := setupConfig()
	go func() {
		server.Run(config.ServerPort)
	}()
	clientRun(config)
}

func clientRun(config *config) {
	fs := http.FileServer(http.Dir(config.ClientPath))
	log.Println("Serving React client on http://localhost" + config.ClientPort)
	log.Fatal(http.ListenAndServe(config.ClientPort, fs))
}

func setupLogger() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func setupConfig() *config {
	return &config{
		ServerPort: ":8080",
		ClientPort: ":4200",
		ClientPath: "client/build",
	}
}
