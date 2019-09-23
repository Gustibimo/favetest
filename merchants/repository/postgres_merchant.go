package repository

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"

	merchants "github.com/Gustibimo/favetest/merchants"
	model "github.com/Gustibimo/favetest/model"
	"github.com/labstack/gommon/log"
)

type postgresMerchantRepository struct {
	Conn *sql.DB
}

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewPostgresMerchantRepository(Conn *sql.DB) merchants.MerchantRepository {
	return &postgresMerchantRepository{Conn}
}

func (m *postgresMerchantRepository) fetch(query string, args ...interface{}) ([]*model.Merchants, error) {

	rows, err := m.Conn.Query(query, args...)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*model.Merchants, 0)
	for rows.Next() {
		t := new(model.Merchants)
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Address,
			&t.City,
			&t.Category,
			&t.Rating,
			&t.Logo,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (m *postgresMerchantRepository) Fetch(cursor string, limit int64) ([]*model.Merchants, error) {

	query := `SELECT id,name,address, city, category, rating, logo
  						FROM merchants WHERE ID > ? LIMIT ?`

	return m.fetch(query, cursor, limit)

}

func (m *postgresMerchantRepository) GetByID(id int64) (res *model.Merchants, err error) {
	query := `SELECT id,name,address, city, category, rating, logo
	FROM merchants WHERE ID = ?`

	list, err := m.fetch(query, id)
	if err != nil {
		return nil, err
	}

	a := &model.Merchants{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, model.ErrNotFound
	}

	return a, nil
}

func (m *postgresMerchantRepository) Create(a *model.Merchants) (int64, error) {

	query := `INSERT  merchants SET name=? , address=? , rating=?, fave_paycnt=? , city=?, category=?, logo=?`
	stmt, err := m.Conn.Prepare(query)
	if err != nil {

		return 0, err
	}
	res, err := stmt.Exec(a.Name, a.Address, a.Rating, a.FavePayCnt, a.City, a.Category, a.Logo)
	if err != nil {
		return 0, err
	}
	log.Debug("Name: ", a.Name)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (m *postgresMerchantRepository) Delete(id int64) error {
	query := "DELETE FROM merchants WHERE id = ?"

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {

		return err
	}
	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", rowsAfected)
		log.Error(err)
		return err
	}

	return nil
}
func (m *postgresMerchantRepository) Update(mr *model.Merchants) (*model.Merchants, error) {
	query := `UPDATE merchants set name=?, address=?, city=?, category=? WHERE ID = ?`

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return nil, nil
	}

	res, err := stmt.Exec(mr.Name, mr.Address, mr.City, mr.Category, mr.ID)
	if err != nil {
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)
		log.Error(err)
		return nil, err
	}

	return nil, nil
}

// DecodeCursor will decode cursor from user for mysql
func DecodeCursor(encodedTime string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(byt)
	t, err := time.Parse(timeFormat, timeString)

	return t, err
}

// EncodeCursor will encode cursor from mysql to user
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}
