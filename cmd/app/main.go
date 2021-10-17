package main

import (
	"os"

	"fibo-prj/internal/config"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := config.Config(); err != nil {
		logrus.Errorf("config err: [%v]", err)
		os.Exit(1)
	}

	os.Exit(f0())
}
