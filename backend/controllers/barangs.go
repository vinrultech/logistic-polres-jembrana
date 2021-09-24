package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/models"
)

func (a *App) CreateBarang(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)

	var r models.Barang

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error bind barang : %v", err))
	}

	var unitKerja models.UnitKerja

	if claims.UnitKerjaID == 0 {
		unitKerja = models.GetUnitKerja()
		r.UnitKerja = unitKerja
	} else {
		unitKerja, err := a.M.FetchUnitKerja(int64(claims.UnitKerjaID))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error fetch unit kerja barang : %v", err))
		}
		r.UnitKerja = unitKerja
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	u1 := uuid.NewV1()

	r.RowID = u1.String()

	err := a.M.CreateBarang(r)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error create barang : %v", err))
	}

	return c.JSON(http.StatusCreated, constants.H{
		"message": "Created",
	})
}

func (a *App) GetBarang(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)

	limitParam := c.QueryParam("limit")
	lastIDParam := c.QueryParam("last_id")

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error barang convert limit param to int : %v", err))
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error barang convert last_id param to int : %v", err))
	}

	var filterValues []interface{}
	var filters []string

	//var items []models.SuratMasuk

	if claims.UnitKerjaID != 0 {
		filters = append(filters, " b.unit_kerja_id=? ")
		filterValues = append(filterValues, int64(claims.UnitKerjaID))
	}

	items, err := a.M.GetBarang(int64(lastID), limit, filters, filterValues)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error get barang : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) SearchBarang(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)

	limitParam := c.QueryParam("limit")
	lastIDParam := c.QueryParam("last_id")
	filter := c.QueryParam("filter")
	search := c.QueryParam("search")

	filter = fmt.Sprintf("b.%s", filter)

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error barang convert limit param to int : %v", err))
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error barang convert last_id param to int : %v", err))
	}

	var filterValues []interface{}
	var filters []string

	if claims.UnitKerjaID != 0 {
		filters = append(filters, " b.unit_kerja_id=? ")
		filterValues = append(filterValues, int64(claims.UnitKerjaID))
	}

	items, err := a.M.SearchBarang(int64(lastID), limit, []string{filter, search}, filters, filterValues)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error search barang : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) UpdateBarang(c echo.Context) error {

	rowID := c.Param("row_id")

	var r models.Barang

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error bind barang : %v", err))
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	err := a.M.UpdateBarang(r, rowID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error update barang : %v", err))
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Updated",
	})
}

func (a *App) RemoveBarang(c echo.Context) error {

	rowID := c.Param("row_id")

	err := a.M.RemoveBarang(rowID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error remove barang : %v", err))
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Removed",
	})
}

func (a *App) FetchBarang(c echo.Context) error {

	rowID := c.Param("row_id")

	Barang, err := a.M.FetchBarang(rowID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error get barang : %v", err))
	}

	return c.JSON(http.StatusOK, Barang)
}
