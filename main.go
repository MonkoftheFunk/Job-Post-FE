package main

import (
	server "Job-Post-FE/srv"
	"embed"
)

//go:embed dist/*
var dist embed.FS

func main() {
	// todo set from docker env
	err := server.Run(9080, dist)
	if err != nil {
		panic("Failed to setup server: " + err.Error())
	}
}
