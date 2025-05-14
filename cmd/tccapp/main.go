package main

import (
	"github.com/darthxd/tcc-app/internal/config"
	"github.com/darthxd/tcc-app/internal/http/router"
)

func main() {
	config.Init()
	port := config.GetPort()

	// Run the server
	router.RouterInit(port)
}
