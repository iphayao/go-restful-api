package main

import (
	"github.com/iphayao/go-restful-api/app"
)

func main() {
	app := &app.App{}
	app.Run(":8080")
}
