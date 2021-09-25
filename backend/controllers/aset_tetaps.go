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

func (a *App) CreateAsetTetap(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)

	var r models.AsetTetap

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error bind AsetTetap : %v", err))
	}

	var unitKerja models.UnitKerja

	if claims.UnitKerjaID == 0 {
		unitKerja = models.GetUnitKerja()
		r.UnitKerja = unitKerja
	} else {
		unitKerja, err := a.M.FetchUnitKerja(int64(claims.UnitKerjaID))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error fetch unit kerja AsetTetap : %v", err))
		}
		r.UnitKerja = unitKerja
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	u1 := uuid.NewV1()

	r.RowID = u1.String()

	err := a.M.CreateAsetTetap(r)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error create AsetTetap : %v", err))
	}

	return c.JSON(http.StatusCreated, constants.H{
		"message": "Created",
	})
}

func (a *App) GetAsetTetap(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)

	limitParam := c.QueryParam("limit")
	lastIDParam := c.QueryParam("last_id")

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error AsetTetap convert limit param to int : %v", err))
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error AsetTetap convert last_id param to int : %v", err))
	}

	var filterValues []interface{}
	var filters []string

	//var items []models.SuratMasuk

	if claims.UnitKerjaID != 0 {
		filters = append(filters, " a.unit_kerja_id=? ")
		filterValues = append(filterValues, int64(claims.UnitKerjaID))
	}

	filterByParam := c.QueryParam("filter_by")
	filterByValue := c.QueryParam("filter_value")

	if filterByParam != "" && filterByValue != "" {
		filterValue, err := strconv.Atoi(filterByValue)

		if err != nil {
			loggers.Log.Errorln(err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error AsetTetap convert filter_value param to int : %v", err))
		}
		filters = append(filters, fmt.Sprintf("a.%s=?", filterByParam))
		filterValues = append(filterValues, filterValue)
	}

	items, err := a.M.GetAsetTetap(int64(lastID), limit, filters, filterValues)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error get AsetTetap : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) SearchAsetTetap(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)

	limitParam := c.QueryParam("limit")
	lastIDParam := c.QueryParam("last_id")
	filter := c.QueryParam("filter")
	search := c.QueryParam("search")

	filter = fmt.Sprintf("a.%s", filter)

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error AsetTetap convert limit param to int : %v", err))
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error AsetTetap convert last_id param to int : %v", err))
	}

	var filterValues []interface{}
	var filters []string

	if claims.UnitKerjaID != 0 {
		filters = append(filters, " a.unit_kerja_id=? ")
		filterValues = append(filterValues, int64(claims.UnitKerjaID))
	}

	filterByParam := c.QueryParam("filter_by")
	filterByValue := c.QueryParam("filter_value")

	if filterByParam != "" && filterByValue != "" {
		filterValue, err := strconv.Atoi(filterByValue)

		if err != nil {
			loggers.Log.Errorln(err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error AsetTetap convert filter_value param to int : %v", err))
		}
		filters = append(filters, fmt.Sprintf("a.%s=?", filterByParam))
		filterValues = append(filterValues, filterValue)
	}

	items, err := a.M.SearchAsetTetap(int64(lastID), limit, []string{filter, search}, filters, filterValues)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error search AsetTetap : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) UpdateAsetTetap(c echo.Context) error {

	rowID := c.Param("row_id")

	var r models.AsetTetap

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error bind AsetTetap : %v", err))
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	err := a.M.UpdateAsetTetap(r, rowID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error update AsetTetap : %v", err))
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Updated",
	})
}

func (a *App) RemoveAsetTetap(c echo.Context) error {

	rowID := c.Param("row_id")

	err := a.M.RemoveAsetTetap(rowID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error remove AsetTetap : %v", err))
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Removed",
	})
}

func (a *App) FetchAsetTetap(c echo.Context) error {

	rowID := c.Param("row_id")

	AsetTetap, err := a.M.FetchAsetTetap(rowID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error get AsetTetap : %v", err))
	}

	return c.JSON(http.StatusOK, AsetTetap)
}
