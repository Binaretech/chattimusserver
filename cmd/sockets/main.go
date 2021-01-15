package main

import (
	"github.com/Binaretech/chattimus/config"
	"github.com/Binaretech/chattimus/pkg/server"
)

func main() {
	config.InitConfig()
	server.Server()
}
