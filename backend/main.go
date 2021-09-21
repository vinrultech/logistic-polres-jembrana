package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/controllers"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/db"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/models"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/routers"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

var config *viper.Viper

func init() {
	config = viper.New()
	config.SetConfigType("json")
	config.AddConfigPath("./")
	config.SetConfigName("app.config")

	err := config.ReadInConfig()
	if err != nil {
		loggers.NewLogger().Errorln(err.Error())
		os.Exit(0)
	}
}

func main() {

	loggers.NewLogger()

	db := db.InitDb(config)

	m := models.New(db)

	a := controllers.New(m)

	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.HTTPErrorHandler = func(err error, c echo.Context) {

		report, ok := err.(*echo.HTTPError)

		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		} else {
			if report.Code == 400 {
				report.Message = "Anda tidak punya autorisasi"
			} else if report.Code == 401 {
				report.Message = "Token anda invalid atau sesi anda berakhir"
			}
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s dibutuhkan", err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s harus lebih besar dari %s", err.Field(), err.Param())

				}

				break
			}
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}

	routers.SetupRouters(e, a)

	e.Static("/upload", "upload")
	e.Static("/files", "files")

	/*

		e.GET("/hello", func(c echo.Context) error {

			//return c.String(http.StatusOK, "Aku padamu")
			password, _ := constants.HashPassword("25d55ad283aa400af464c76d713c07ad")
			return c.String(http.StatusOK, password)

		})

	*/

	go func() {
		port := config.GetString("port")
		if err := e.Start(port); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
