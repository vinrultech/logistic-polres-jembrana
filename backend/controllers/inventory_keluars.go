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

func (a *App) CreateInventoryKeluar(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)

	var ik struct {
		BarangID string `json:"barang_id" validate:"required"`
		Tanggal  string `json:"tanggal" validate:"required"`
		Jumlah   int64  `json:"jumlah" validate:"required"`
	}

	if err := c.Bind(&ik); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error bind InventoryKeluar : %v", err))
	}

	if err := c.Validate(&ik); err != nil {
		return err
	}

	barang, err := a.M.FetchBarang(ik.BarangID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error create InventoryKeluar, fetch barang : %v", err))
	}

	r := models.InventoryKeluar{}

	var unitKerja models.UnitKerja

	if claims.UnitKerjaID == 0 {
		unitKerja = models.GetUnitKerja()
		r.UnitKerja = unitKerja
	} else {
		unitKerja, err := a.M.FetchUnitKerja(int64(claims.UnitKerjaID))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error fetch unit kerja InventoryKeluar : %v", err))
		}
		r.UnitKerja = unitKerja
	}

	u1 := uuid.NewV1()

	r.RowID = u1.String()
	r.Barang = barang
	r.Tanggal = ik.Tanggal
	r.Jumlah = ik.Jumlah

	if err := a.M.CreateInventoryKeluar(r, barang); err != nil {
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error create InventoryKeluar : %v", err))
		}
	}

	return c.JSON(http.StatusCreated, constants.H{
		"message": "Created",
	})
}

func (a *App) GetInventoryKeluar(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)

	limitParam := c.QueryParam("limit")
	lastIDParam := c.QueryParam("last_id")

	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error InventoryKeluar convert limit param to int : %v", err))
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error InventoryKeluar convert last_id param to int : %v", err))
	}

	var filterValues []interface{}
	var filters []string

	//var items []models.SuratKeluar

	if claims.UnitKerjaID != 0 {
		filters = append(filters, " ik.unit_kerja_id=? ")
		filterValues = append(filterValues, int64(claims.UnitKerjaID))
	}

	filterByParam := c.QueryParam("filter_by")
	filterByValue := c.QueryParam("filter_value")

	if filterByParam != "" && filterByValue != "" {
		filterValue, err := strconv.Atoi(filterByValue)

		if err != nil {
			loggers.Log.Errorln(err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error InventoryKeluar convert filter_value param to int : %v", err))
		}
		filters = append(filters, fmt.Sprintf("b.%s=?", filterByParam))
		filterValues = append(filterValues, filterValue)
	}

	items, err := a.M.GetInventoryKeluar(int64(lastID), limit, filters, filterValues)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error get InventoryKeluar : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) SearchInventoryKeluar(c echo.Context) error {

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
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error InventoryKeluar convert limit param to int : %v", err))
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error InventoryKeluar convert last_id param to int : %v", err))
	}

	var filterValues []interface{}
	var filters []string

	if claims.UnitKerjaID != 0 {
		filters = append(filters, " ik.unit_kerja_id=? ")
		filterValues = append(filterValues, int64(claims.UnitKerjaID))
	}

	filterByParam := c.QueryParam("filter_by")
	filterByValue := c.QueryParam("filter_value")

	if filterByParam != "" && filterByValue != "" {
		filterValue, err := strconv.Atoi(filterByValue)

		if err != nil {
			loggers.Log.Errorln(err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error InventoryKeluar convert filter_value param to int : %v", err))
		}
		filters = append(filters, fmt.Sprintf("b.%s=?", filterByParam))
		filterValues = append(filterValues, filterValue)
	}

	items, err := a.M.SearchInventoryKeluar(int64(lastID), limit, []string{filter, search}, filters, filterValues)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error search InventoryKeluar : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) RemoveInventoryKeluar(c echo.Context) error {

	rowID := c.Param("row_id")

	im, err := a.M.FetchInventoryKeluar(rowID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error remove InventoryKeluar, fetch : %v", err))
	}

	if err := a.M.RemoveInventoryKeluar(im, im.Barang); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error remove InventoryKeluar : %v", err))
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Removed",
	})
}
