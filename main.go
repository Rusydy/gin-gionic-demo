package main

import "gin-gionic-demo/routers"

func main() {
	routers.StartServer().Run(":8080")
}
