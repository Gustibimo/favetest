package merchants

import (
	model "github.com/Gustibimo/favetest/model"
)

type MerchantRepository interface {
	Fetch(cursor string, limit int64) (res []*model.Merchants, err error)
	GetByID(id int64) (*model.Merchants, error)
	// GetByName(name string) (*model.Merchants, error)
	Update(merchant *model.Merchants) (*model.Merchants, error)
	Store(s *model.Merchants) (int64, error)
	Delete(id int64) error
}
