package impl

import (
	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/config"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/domain/user/repository"
)

type NewUserServiceParams struct {
	Conf           *config.Config
	UserRepository repository.UserRepository
}
type UserService struct {
	conf           *config.Config
	UserRepository repository.UserRepository
}

func New(params *NewUserServiceParams) *UserService {
	return &UserService{
		conf:           params.Conf,
		UserRepository: params.UserRepository,
	}
}
