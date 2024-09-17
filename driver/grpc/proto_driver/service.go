package proto_driver

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

	UnimplementedDriverServiceServer
}

func (s *Server) GetDriverByDistance(ctx context.Context, req *DriverFilter) (*DriverData, error) {
	latitude := req.Latitude
	longitude := req.Longitude

	var driverData DriverData
	s.Logger.Debug("container injected!")
	query := fmt.Sprintf(`SELECT u.id,u.name,u.email,u.role,u.phone,u.latitude,u.longitude
	FROM users u
	WHERE (
		6371 * acos(
			cos(radians(%f)) 
			* cos(radians(latitude)) 
			* cos(radians(longitude) - radians(%f)) 
			+ sin(radians(%f)) 
			* sin(radians(latitude))
		)
	) < 7
	AND u.role = %d
	ORDER BY RAND()
	LIMIT 1;`, latitude, longitude, latitude, USER_ROLE_DRIVER)
	err := s.SQL.QueryRowContext(ctx, query).Scan(&driverData.Id, &driverData.Name, &driverData.Email, &driverData.Role, &driverData.Phone, &driverData.Latitude, &driverData.Longitude)
	if err != nil {
		return nil, err
	}
	return &driverData, nil
}
