package main

import (
	server "Job-Post-FE/srv"
	"Job-Post-FE/srv/mongo"
	"Job-Post-FE/srv/session"
	"embed"
	"github.com/redis/go-redis/v9"
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

	// Mongo
	mconfig := mongo.Config{
		DSN:         os.Getenv("MONGO_DB_DSN"),
		ConnTimeout: time.Second * 10,
		Database:    "platform",
	}
	if mconfig.DSN == "" {
		log.Fatal("You must set your 'MONGO_DB_DSN' environment variable.")
	}

	// Redis - read only
	opts, err := redis.ParseURL(os.Getenv("REDIS_DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	sconfig := session.Config{
		Config:  *opts,
		Key:     os.Getenv("APP_KEY"),
		Cookie:  os.Getenv("SESSION_COOKIE_NAME"),
		Prefix:  os.Getenv("SESSION_PREFIX"),
		UserKey: os.Getenv("SESSION_USER_KEY"),
	}

	err = server.Run(port, dist, &mconfig, &sconfig)
	if err != nil {
		panic("Failed to setup server: " + err.Error())
	}
}
