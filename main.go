package main

import (
	"github.com/darthxd/tcc-app/config"
	"github.com/darthxd/tcc-app/router"
)

func main() {
	config.Init()
	port := config.GetPort()

	// Run the server
	router.RouterInit(port)
}
