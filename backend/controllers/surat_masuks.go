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

func (a *App) CreateSuratMasuk(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)

	var r models.SuratMasuk

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error bind surat masuk : %v", err))
	}

	var unitKerja models.UnitKerja

	if claims.UnitKerjaID == 0 {
		unitKerja = models.GetUnitKerja()
		r.UnitKerja = unitKerja
	} else {
		unitKerja, err := a.M.FetchUnitKerja(int64(claims.UnitKerjaID))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error fetch unit kerja surat masuk: %v", err))
		}
		r.UnitKerja = unitKerja
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	u1 := uuid.NewV1()

	r.RowID = u1.String()

	err := a.M.CreateSuratMasukWithTransaction(r)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error create surat masuk with transaction : %v", err))
	}

	return c.JSON(http.StatusCreated, constants.H{
		"message": "Created",
	})
}

func (a *App) GetSuratMasuk(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)

	limitParam := c.QueryParam("limit")
	lastIDParam := c.QueryParam("last_id")

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error surat masuk convert limit param to int : %v", err))
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error surat masuk convert last_id param to int : %v", err))
	}

	startDate := c.QueryParam("start")
	endDate := c.QueryParam("end")

	var items []models.SuratMasuk

	if startDate != "" && endDate != "" {
		if claims.UnitKerjaID == 0 {
			items, err = a.M.GetSuratMasuk(int64(lastID), limit, startDate, endDate)
		} else {
			items, err = a.M.GetSuratMasukWithFilter(int64(lastID), limit, int64(claims.UnitKerjaID), startDate, endDate)
		}

	} else {
		if claims.UnitKerjaID == 0 {
			items, err = a.M.GetSuratMasuk(int64(lastID), limit)
		} else {
			items, err = a.M.GetSuratMasukWithFilter(int64(lastID), limit, int64(claims.UnitKerjaID))
		}

	}

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error get surat masuk : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) SearchSuratMasuk(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)

	limitParam := c.QueryParam("limit")
	lastIDParam := c.QueryParam("last_id")
	filter := c.QueryParam("filter")
	search := c.QueryParam("search")

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error surat masuk convert limit param to int : %v", err))
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error surat masuk convert last_id param to int : %v", err))
	}

	startDate := c.QueryParam("start")
	endDate := c.QueryParam("end")

	var items []models.SuratMasuk

	if startDate != "" && endDate != "" {
		if claims.UnitKerjaID == 0 {
			items, err = a.M.SearchSuratMasuk(int64(lastID), limit, search, filter, startDate, endDate)
		} else {
			items, err = a.M.SearchSuratMasukWithFilter(int64(lastID), limit, search, filter,
				int64(claims.UnitKerjaID), startDate, endDate)
		}

	} else {
		if claims.UnitKerjaID == 0 {
			items, err = a.M.SearchSuratMasuk(int64(lastID), limit, search, filter)
		} else {
			items, err = a.M.SearchSuratMasukWithFilter(int64(lastID), limit, search, filter,
				int64(claims.UnitKerjaID))
		}

	}

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error search surat masuk : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) UpdateSuratMasuk(c echo.Context) error {

	rowID := c.Param("row_id")

	var r models.SuratMasuk

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error bind surat masuk : %v", err))
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	err := a.M.UpdateSuratMasukWithTransaction(r, rowID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error update surat masuk with transaction : %v", err))
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Updated",
	})
}

func (a *App) RemoveSuratMasuk(c echo.Context) error {

	rowID := c.Param("row_id")

	files, err := a.M.RemoveSuratMasukTransaction(rowID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error remove surat masuk transaction : %v", err))
	}

	for _, file := range files {

		filename := fmt.Sprintf("files/%s_%s", file.FileID, file.Filename)

		err = os.Remove(filename)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error surat masuk,remove file : %v", err))
		}
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Removed",
	})
}

func (a *App) FetchSuratMasuk(c echo.Context) error {

	rowID := c.Param("row_id")

	suratMasuk, err := a.M.FetchSuratMasuk(rowID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error get surat masuk : %v", err))
	}

	return c.JSON(http.StatusOK, suratMasuk)
}
