package usecase

import (
	"strconv"

	"github.com/Gustibimo/favetest/merchants"
	model "github.com/Gustibimo/favetest/model"
	"github.com/labstack/gommon/log"
)

type merchantUcase struct {
	merchantRepos merchants.MerchantRepository
}

type MerchantUcase interface {
	Fetch(cursor string, limit int64) ([]*model.Merchants, string, error)
	GetByID(id int64) (*model.Merchants, error)
	// GetByName(name string) (*model.Merchants, error)
	Update(merchant *model.Merchants) (*model.Merchants, error)
	Store(s *model.Merchants) (*model.Merchants, error)
	Delete(id int64) error
}

func NewMerchantUsecase(m merchants.MerchantRepository) MerchantUcase {
	return &merchantUcase{
		merchantRepos: m,
	}
}

func (m *merchantUcase) Delete(id int64) error {
	existedMerchant, _ := m.merchantRepos.GetByID(id)
	log.Info("masuk sini")
	if existedMerchant == nil {
		log.Info("Masuk Sini2")
		return model.ErrNotFound
	}
	log.Info("Masuk Sini3")

	return m.merchantRepos.Delete(id)
}

func (m *merchantUcase) Fetch(cursor string, limit int64) ([]*model.Merchants, string, error) {
	if limit == 0 {
		limit = 10
	}

	listMerchant, err := m.merchantRepos.Fetch(cursor, limit)
	if err != nil {
		return nil, "", err
	}
	nextCursor := ""
	if size := len(listMerchant); size == int(limit) {
		lastId := listMerchant[limit-1].ID
		nextCursor = strconv.Itoa(int(lastId))
	}

	return listMerchant, nextCursor, nil
}

func (m *merchantUcase) GetByID(id int64) (*model.Merchants, error) {

	res, err := m.merchantRepos.GetByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *merchantUcase) Update(mr *model.Merchants) (*model.Merchants, error) {
	return m.merchantRepos.Update(mr)
}

func (m *merchantUcase) Store(s *model.Merchants) (*model.Merchants, error) {

	existedMerchant, _ := m.GetByID(s.ID)
	if existedMerchant != nil {
		return nil, model.ErrConflict
	}

	id, err := m.merchantRepos.Store(s)
	if err != nil {
		return nil, err
	}
	s.ID = id

	return s, nil
}
