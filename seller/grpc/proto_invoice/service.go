package proto_invoice

import (
	context "context"
	fmt "fmt"
	"seller/grpc/proto_product"

	"math/rand/v2"

	"gofr.dev/pkg/gofr/container"
)

const INVOICE_STATUS_SEARCH_DRIVER = 1
const INVOICE_STATUS_PROCESS = 2
const INVOICE_STATUS_DELIVERY = 3
const INVOICE_STATUS_FINISH = 4
const INVOICE_STATUS_CANCEL = 5

type Server struct {
	*container.Container

	UnimplementedInvoiceServiceServer
}

func (s *Server) CreateInvoice(ctx context.Context, req *InvoiceFilter) (*InvoiceData, error) {
	driver_id := req.DriverId
	user_id := req.UserId
	product_id := req.ProductId
	from_latitude := req.FromLatitude
	from_longitude := req.FromLongitude
	to_latitude := req.ToLatitude
	to_longitude := req.ToLongitude
	address := req.Address
	product, _ := s.GetProductById(ctx, product_id)
	res, err := s.SQL.ExecContext(ctx, "INSERT INTO invoices (driver_id,user_id,code,shipping_cost,from_latitude,from_longitude,to_latitude,to_longitude,address,status) VALUES (?,?,?,?,?,?,?,?,?,?)",
		driver_id,
		user_id,
		rand.IntN(10000),
		20000,
		from_latitude,
		from_longitude,
		to_latitude,
		to_longitude,
		address,
		INVOICE_STATUS_FINISH)
	if err != nil {
		return nil, err
	}
	lastId, _ := res.LastInsertId()
	_, err = s.SQL.ExecContext(ctx, "INSERT INTO invoice_details (invoice_id,product_id,price,qty) VALUES (?,?,?,?)",
		lastId,
		product_id,
		product.Price,
		1,
	)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (s *Server) GetProductById(ctx context.Context, Id int32) (*proto_product.ProductData, error) {
	var productData proto_product.ProductData
	query := fmt.Sprintf("select p.id,p.user_id,p.name,p.price,p.qty,p.desc from products p WHERE p.id=%d", Id)
	err := s.SQL.QueryRowContext(ctx, query).Scan(&productData.Id, &productData.UserId, &productData.Name, &productData.Price, &productData.Qty, &productData.Desc)
	if err != nil {
		return nil, err
	}
	return &productData, nil
}
