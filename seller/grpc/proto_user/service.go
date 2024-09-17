package proto_user

import (
	context "context"
	fmt "fmt"

	"gofr.dev/pkg/gofr/container"
)

const USER_ROLE_CUSTOMER = 1
const USER_ROLE_SELLER = 2
const USER_ROLE_DRIVER = 3

type Server struct {
	*container.Container

	UnimplementedUserServiceServer
}

func (s *Server) GetUserById(ctx context.Context, req *UserFilter) (*UserData, error) {
	Id := req.Id
	var userData UserData
	query := fmt.Sprintf("SELECT id,name,email,role,phone,latitude,longitude FROM users u WHERE u.role=%d AND u.id=%d", USER_ROLE_SELLER, Id)
	err := s.SQL.QueryRowContext(ctx, query).Scan(&userData.Id, &userData.Name, &userData.Email, &userData.Role, &userData.Phone, &userData.Latitude, &userData.Longitude)
	if err != nil {
		return nil, err
	}
	return &userData, nil
}
