package main

import (
	"github.com/alserok/g8s/internal/app"
	"github.com/alserok/g8s/internal/config"
)

func main() {
	app.MustServe(config.MustLoad())
}
