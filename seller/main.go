package main

import (
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.GET("/invoice/confirm", func(ctx *gofr.Context) (interface{}, error) {

		return nil, nil
	})
	app.Run()
}
