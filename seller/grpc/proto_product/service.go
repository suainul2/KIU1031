package proto_product

import (
	context "context"
	fmt "fmt"

	"gofr.dev/pkg/gofr/container"
)

type Server struct {
	*container.Container
	UnimplementedProductServiceServer
}

func (s *Server) GetProductById(ctx context.Context, req *ProductFilter) (*ProductData, error) {
	Id := req.Id
	var productData ProductData
	query := fmt.Sprintf("select p.id,p.user_id,p.name,p.price,p.qty,p.desc from products p WHERE p.id=%d", Id)
	err := s.SQL.QueryRowContext(ctx, query).Scan(&productData.Id, &productData.UserId, &productData.Name, &productData.Price, &productData.Qty, &productData.Desc)
	if err != nil {
		return nil, err
	}
	return &productData, nil
}
