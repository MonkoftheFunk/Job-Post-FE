package main

import (
	server "Job-Post-FE/srv"
	"Job-Post-FE/srv/mongo"
	"embed"
	"log"
	"os"
	"strconv"
	"time"
)

//go:embed dist/*
var dist embed.FS

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic("Failed to get PORT: " + err.Error())
	}
	config := mongo.Config{
		DSN:         os.Getenv("MONGO_DB_DSN"),
		ConnTimeout: time.Second * 10,
	}
	if config.DSN == "" {
		log.Fatal("You must set your 'MONGO_DB_DSN' environment variable.")
	}
	err = server.Run(port, dist, mongo.NewClient(&config))
	if err != nil {
		panic("Failed to setup server: " + err.Error())
	}
}
