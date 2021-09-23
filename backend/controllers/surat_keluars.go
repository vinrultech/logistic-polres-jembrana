package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/models"
)

func (a *App) CreateSuratKeluar(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)

	var r models.SuratKeluar

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error bind surat keluar : %v", err))
	}

	var unitKerja models.UnitKerja

	if claims.UnitKerjaID == 0 {
		unitKerja = models.GetUnitKerja()
		r.UnitKerja = unitKerja
	} else {
		unitKerja, err := a.M.FetchUnitKerja(int64(claims.UnitKerjaID))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error fetch unit kerja surat keluar : %v", err))
		}
		r.UnitKerja = unitKerja
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	u1 := uuid.NewV1()

	r.RowID = u1.String()

	err := a.M.CreateSuratKeluarWithTransaction(r)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error create surat keluar with transaction : %v", err))
	}

	return c.JSON(http.StatusCreated, constants.H{
		"message": "Created",
	})
}

func (a *App) GetSuratKeluar(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)

	limitParam := c.QueryParam("limit")
	lastIDParam := c.QueryParam("last_id")

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error surat keluar convert limit param to int : %v", err))
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error surat keluar convert last_id param to int : %v", err))
	}

	startDate := c.QueryParam("start")
	endDate := c.QueryParam("end")

	var items []models.SuratKeluar

	if startDate != "" && endDate != "" {
		if claims.UnitKerjaID == 0 {
			items, err = a.M.GetSuratKeluar(int64(lastID), limit, startDate, endDate)
		} else {
			items, err = a.M.GetSuratKeluarWithFilter(int64(lastID), limit, int64(claims.UnitKerjaID), startDate, endDate)
		}

	} else {
		if claims.UnitKerjaID == 0 {
			items, err = a.M.GetSuratKeluar(int64(lastID), limit)
		} else {
			items, err = a.M.GetSuratKeluarWithFilter(int64(lastID), limit, int64(claims.UnitKerjaID))
		}

	}

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error get surat keluar : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) SearchSuratKeluar(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)

	limitParam := c.QueryParam("limit")
	lastIDParam := c.QueryParam("last_id")
	filter := c.QueryParam("filter")
	search := c.QueryParam("search")

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error surat keluar convert limit param to int : %v", err))
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error surat keluar convert last_id param to int : %v", err))
	}

	startDate := c.QueryParam("start")
	endDate := c.QueryParam("end")

	var items []models.SuratKeluar

	if startDate != "" && endDate != "" {
		if claims.UnitKerjaID == 0 {
			items, err = a.M.SearchSuratKeluar(int64(lastID), limit, search, filter, startDate, endDate)
		} else {
			items, err = a.M.SearchSuratKeluarWithFilter(int64(lastID), limit, search, filter,
				int64(claims.UnitKerjaID), startDate, endDate)
		}

	} else {
		if claims.UnitKerjaID == 0 {
			items, err = a.M.SearchSuratKeluar(int64(lastID), limit, search, filter)
		} else {
			items, err = a.M.SearchSuratKeluarWithFilter(int64(lastID), limit, search, filter,
				int64(claims.UnitKerjaID))
		}

	}

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error search surat keluar : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) UpdateSuratKeluar(c echo.Context) error {

	rowID := c.Param("row_id")

	var r models.SuratKeluar

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error bind surat keluar : %v", err))
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	err := a.M.UpdateSuratKeluarWithTransaction(r, rowID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error update surat keluar with transaction : %v", err))
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Updated",
	})
}

func (a *App) RemoveSuratKeluar(c echo.Context) error {

	rowID := c.Param("row_id")

	files, err := a.M.RemoveSuratKeluarTransaction(rowID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error remove surat keluar transaction : %v", err))
	}

	for _, file := range files {

		filename := fmt.Sprintf("files/%s_%s", file.FileID, file.Filename)

		err = os.Remove(filename)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error surat keluar,remove file : %v", err))
		}
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Removed",
	})
}

func (a *App) FetchSuratKeluar(c echo.Context) error {

	rowID := c.Param("row_id")

	SuratKeluar, err := a.M.FetchSuratKeluar(rowID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error get surat keluar : %v", err))
	}

	return c.JSON(http.StatusOK, SuratKeluar)
}
