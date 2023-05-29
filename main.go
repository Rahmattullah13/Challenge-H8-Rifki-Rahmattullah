package main

import (
	"go-trial-class/cli"
	"go-trial-class/config"
)

func main() {
	config.DBConnect()
	cli.MainMenu()
}
