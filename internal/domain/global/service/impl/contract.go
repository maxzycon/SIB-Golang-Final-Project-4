package impl

import (
	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/config"
	"gorm.io/gorm"
)

type NewGlobalServiceParams struct {
	Conf *config.Config
	Db   *gorm.DB
}
type GlobalService struct {
	conf *config.Config
	db   *gorm.DB
}

func New(params *NewGlobalServiceParams) *GlobalService {
	return &GlobalService{
		conf: params.Conf,
		db:   params.Db,
	}
}
