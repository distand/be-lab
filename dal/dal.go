package dal

import (
	"be-lab/common/infra"
	"gorm.io/gorm"
)

type Dal struct {
	DB *gorm.DB
}

func NewDal() *Dal {
	return &Dal{
		DB: infra.DB,
	}
}
