package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/models"
)

func (a *App) CreateSuratMasuk(c echo.Context) error {

	var r models.SuratMasuk

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	u1 := uuid.NewV1()

	r.RowID = u1.String()

	err := a.M.CreateSuratMasuk(r)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	for _, file := range r.Files {
		file.RowID = r.RowID
		file.Tipe = "surat_masuk"
		err = a.M.CreateFile(file)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, constants.H{
		"message": "Created",
	})
}

func (a *App) GetSuratMasuk(c echo.Context) error {

	limitParam := c.QueryParam("limit")
	lastIDParam := c.QueryParam("last_id")

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	startDate := c.QueryParam("start")
	endDate := c.QueryParam("end")

	var items []models.SuratMasuk

	if startDate != "" && endDate != "" {
		items, err = a.M.GetSuratMasuk(int64(lastID), limit, startDate, endDate)
	} else {
		items, err = a.M.GetSuratMasuk(int64(lastID), limit)
	}

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) SearchSuratMasuk(c echo.Context) error {

	limitParam := c.QueryParam("limit")
	lastIDParam := c.QueryParam("last_id")
	filter := c.QueryParam("filter")
	search := c.QueryParam("search")

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	startDate := c.QueryParam("start")
	endDate := c.QueryParam("end")

	var items []models.SuratMasuk

	if startDate != "" && endDate != "" {
		items, err = a.M.SearchSuratMasuk(int64(lastID), limit, search, filter, startDate, endDate)
	} else {
		items, err = a.M.SearchSuratMasuk(int64(lastID), limit, search, filter)
	}

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) UpdateSuratMasuk(c echo.Context) error {

	rowID := c.Param("row_id")

	var r models.SuratMasuk

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	err := a.M.UpdateSuratMasuk(r, rowID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	for _, file := range r.Files {
		if file.Status == "new" {
			file.RowID = r.RowID
			file.Tipe = "surat_masuk"
			err = a.M.CreateFile(file)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Updated",
	})
}

func (a *App) RemoveSuratMasuk(c echo.Context) error {

	rowID := c.Param("row_id")

	row, err := a.M.FetchSuratMasuk(rowID)

	err = a.M.RemoveSuratMasuk(rowID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	for _, file := range row.Files {
		err := a.M.RemoveFile(file.FileID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		filename := fmt.Sprintf("files/%s_%s", file.FileID, file.Filename)

		err = os.Remove(filename)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Removed",
	})
}
