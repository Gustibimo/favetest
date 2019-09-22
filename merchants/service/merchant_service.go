package service

import (
	"net/http"
	"strconv"

	merchantUcase "github.com/Gustibimo/favetest/merchants/usecase"
	"github.com/Gustibimo/favetest/model"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

type HttpMerchantHandler struct {
	MUsecase merchantUcase.MerchantUcase
}

func (a *HttpMerchantHandler) FetchMerchant(c echo.Context) error {

	limitS := c.QueryParam("limit")
	limit, _ := strconv.Atoi(limitS)

	cursor := c.QueryParam("cursor")

	listAr, nextCursor, err := a.MUsecase.Fetch(cursor, int64(limit))

	if err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}

func (a *HttpMerchantHandler) GetByID(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	merch, err := a.MUsecase.GetByID(id)

	if err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}
	return c.JSON(http.StatusOK, merch)
}

func isRequestValid(m *model.Merchants) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *HttpMerchantHandler) Store(c echo.Context) error {
	var merchant model.Merchants
	err := c.Bind(&merchant)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&merchant); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ar, err := a.MUsecase.Store(&merchant)

	if err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}
	return c.JSON(http.StatusCreated, ar)
}
func (a *HttpMerchantHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	err = a.MUsecase.Delete(id)

	if err != nil {

		return c.JSON(getStatusCode(err), err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case model.ErrInternalServerError:

		return http.StatusInternalServerError
	case model.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func NewArticleHttpHandler(e *echo.Echo, us merchantUcase.MerchantUcase) {
	handler := &HttpMerchantHandler{
		MUsecase: us,
	}

	e.GET("/merchant", handler.FetchMerchant)
	e.POST("merchant", handler.Store)
	e.GET("/merchant/:id", handler.GetByID)
	e.DELETE("/merchant/:id", handler.Delete)

}
