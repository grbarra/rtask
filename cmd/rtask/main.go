package main

import (
	"github.com/sirupsen/logrus"

	"gitlab.com/grbarra/rtask/config"
	"gitlab.com/grbarra/rtask/internal/pkg/app"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	server := app.New(conf)

	server.Run()
}
