package main

import (
	"seller/grpc/proto_invoice"
	"seller/grpc/proto_product"
	"seller/grpc/proto_user"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	proto_user.RegisterUserServiceServer(app, &proto_user.Server{})
	proto_product.RegisterProductServiceServer(app, &proto_product.Server{})
	proto_invoice.RegisterInvoiceServiceServer(app, &proto_invoice.Server{})
	app.Run()
}
