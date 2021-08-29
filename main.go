package main

import (
	"auth/api/routers"
)

func main() {
	routers.New().Run()
}
