package main

import (
	"driver/grpc/proto_driver"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	proto_driver.RegisterDriverServiceServer(app, &proto_driver.Server{})
	app.Run()
}
