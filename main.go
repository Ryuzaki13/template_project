package main

import (
	"awesomeProject/server"
	"awesomeProject/setting"
)

func main() {
	cfg := setting.Load("setting.json")
	server.Run(cfg)
}
