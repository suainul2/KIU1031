package main

import (
	"customer/grpc_driver/proto_driver"
	"customer/grpc_seller/proto_invoice"
	"customer/grpc_seller/proto_product"
	"customer/grpc_seller/proto_user"
	"fmt"

	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.GET("/order/{user_id}/{product_id}", func(ctx *gofr.Context) (interface{}, error) {
		user_id, _ := strconv.Atoi(ctx.PathParam("user_id"))
		user := GetUserById(ctx, int32(user_id))
		product_id, _ := strconv.Atoi(ctx.PathParam("product_id"))
		connSellerRpc := GetConRpcSeller()
		connDriverRpc := GetConRpcDriver()
		defer connSellerRpc.Close()
		defer connDriverRpc.Close()
		pbProduct := proto_product.NewProductServiceClient(connSellerRpc)
		pbSeller := proto_user.NewUserServiceClient(connSellerRpc)
		dataProduct, _ := pbProduct.GetProductById(ctx, &proto_product.ProductFilter{Id: int32(product_id)})
		dataSeller, _ := pbSeller.GetUserById(ctx, &proto_user.UserFilter{Id: dataProduct.UserId})
		if dataSeller != nil && user != nil {
			pbDriver := proto_driver.NewDriverServiceClient(connDriverRpc)
			dataDriver, _ := pbDriver.GetDriverByDistance(ctx, &proto_driver.DriverFilter{Latitude: dataSeller.Latitude, Longitude: dataSeller.Longitude})
			fmt.Println(dataDriver)
			pbInvoice := proto_invoice.NewInvoiceServiceClient(connSellerRpc)
			pbInvoice.CreateInvoice(ctx, &proto_invoice.InvoiceFilter{
				UserId:        int32(user_id),
				DriverId:      dataDriver.Id,
				ProductId:     dataProduct.Id,
				FromLatitude:  dataSeller.Latitude,
				FromLongitude: dataSeller.Longitude,
				ToLatitude:    user.Latitude,
				ToLongitude:   user.Longitude,
			})
			return "ok", nil
		}
		return "error", nil
	})
	app.Run()
}
func GetConRpcSeller() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:9000", opts...)
	if err != nil {
		log.Fatalf("Tidak dapat terhubung: %v", err)
	}
	return conn
}
func GetConRpcDriver() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:9001", opts...)
	if err != nil {
		log.Fatalf("Tidak dapat terhubung: %v", err)
	}
	return conn
}
func GetUserById(ctx *gofr.Context, id int32) *proto_user.UserData {
	var userData proto_user.UserData
	query := fmt.Sprintf("SELECT id,name,email,role,phone,latitude,longitude FROM users u WHERE u.role=1 AND u.id=%d", id)
	err := ctx.SQL.QueryRowContext(ctx, query).Scan(&userData.Id, &userData.Name, &userData.Email, &userData.Role, &userData.Phone, &userData.Latitude, &userData.Longitude)
	if err != nil {
		return nil
	}
	return &userData

}
