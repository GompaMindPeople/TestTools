package main

import (
	_ "TestTools/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	server := g.Server()

	server.SetFileServerEnabled(true)
	server.Run()
}
