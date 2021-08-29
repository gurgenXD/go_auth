package main

import (
	"auth/api/config"
	"auth/api/routers"
	"fmt"
	"log"
)

func main() {
	config.LoadEnv()
	config.Setup()

	addr := fmt.Sprintf(":%d", config.Server.Port)
	if err := routers.New().Run(addr); err != nil {
		log.Fatal(err)
	}
}
