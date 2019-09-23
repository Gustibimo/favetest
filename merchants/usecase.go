package merchants

import (
	model "github.com/Gustibimo/favetest/model"
)

type MerchantImpl interface {
	Fetch(cursor string, limit int64) ([]*model.Merchants, string, error)
	GetByID(id int64) (*model.Merchants, error)
	// GetByName(name string) (*model.Merchants, error)
	Update(merchant *model.Merchants) error
	Create(s *model.Merchants) error
	Delete(id int64) error
}
