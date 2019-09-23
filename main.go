package main

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"

	cfg "github.com/Gustibimo/favetest/config/env"
	merchantRepo "github.com/Gustibimo/favetest/merchants/repository"
	httpDeliver "github.com/Gustibimo/favetest/merchants/service"
	merchantUcase "github.com/Gustibimo/favetest/merchants/usecase"
	"github.com/Gustibimo/go-clean-arch/middleware"
	"github.com/labstack/echo"
)

var config cfg.Config

func init() {
	config = cfg.NewViperConfig()

	if config.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

}

func main() {
	dbHost := config.GetString(`database.host`)
	dbPort := config.GetString(`database.port`)
	dbUser := config.GetString(`database.user`)
	dbPass := config.GetString(`database.pass`)
	dbName := config.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Singapore")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && config.GetBool("debug") {
		fmt.Println(err)
	}
	defer dbConn.Close()
	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)
	mr := merchantRepo.NewPostgresMerchantRepository(dbConn)
	mu := merchantUcase.NewMerchantUsecase(mr)

	httpDeliver.NewMerchantHttpHandler(e, mu)

	e.Start(config.GetString("server.address"))
}
