package main

import (
	"work/routes"
)

func main() {
	Init()
	r := routes.NewRouter()
	r.Run(":9090")
}
