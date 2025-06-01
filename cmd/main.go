package main

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "morse-server: ", log.LstdFlags)
	srv := server.NewServer(logger)
	logger.Println("Server starting on :8080")
	if err := srv.Run(); err != nil {
		logger.Fatal("Server failed: ", err)
	}
}
