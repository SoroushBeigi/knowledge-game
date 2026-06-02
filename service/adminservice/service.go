package adminservice

import "github.com/SoroushBeigi/knowledge-game/entity"

type Service struct {
}

func New() *Service {
	return &Service{}
}

// TODO: implement admin api later
func (s Service) ListAllUsers() ([]entity.User, error) {
	list := make([]entity.User, 0)

	list = append(list, entity.User{
		ID:          0,
		PhoneNumber: "tst",
		Avatar:      "tst",
		Name:        "tst",
		Password:    "tst",
		Role:        entity.AdminRole,
	})

	return list, nil

}
