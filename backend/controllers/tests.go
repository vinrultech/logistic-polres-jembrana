package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/models"
)

func (a *App) TestInsertRow(c echo.Context) error {

	for i := 0; i < 100000; i++ {
		nama := "Alvin "
		alamat := "Negara "
		if i%2 == 0 {
			nama = "Arul "
			alamat = "Pati "
		}

		nama += strconv.Itoa(i)
		alamat += strconv.Itoa(i)

		t := models.Test{
			Nama:   nama,
			Alamat: alamat,
		}

		err := a.M.CreateTest(t)

		if err != nil {
			loggers.Log.Errorln(err.Error())

		}

	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Successfully",
	})

}
