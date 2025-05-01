package main

import (
	"github.com/alserok/g8s/internal/app"
	"github.com/alserok/g8s/internal/config"
)

// @title           g8s API
// @version         0.0.1
// @description     g8s http api
// @host            localhost:{{PORT}}
// @schemes         http
func main() {
	app.MustServe(config.MustLoad())
}
