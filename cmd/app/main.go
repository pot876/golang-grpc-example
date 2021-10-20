package main

import (
	"os"

	"github.com/pot876/golang-grpc-example/internal/config"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := config.Config(); err != nil {
		logrus.Errorf("config err: [%v]", err)
		os.Exit(1)
	}

	os.Exit(f0())
}
