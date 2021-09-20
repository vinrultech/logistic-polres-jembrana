package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/models"
)

func (a *App) CreateKategori(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)
	role := claims.Role

	if role != "superuser" {
		loggers.Log.Errorln("Anda tidak punya kewenangan")
		return echo.NewHTTPError(http.StatusForbidden, "Anda tidak punya kewenangan")
	}

	var r models.Kategori

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	if a.M.CheckKategoriKode(r.Kode) {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Kode kategori %s sudah ada, cari yang lain", r.Kode))
	}

	err := a.M.CreateKategori(r)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, constants.H{
		"message": "Created",
	})
}

func (a *App) UpdateKategori(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)
	role := claims.Role

	if role != "superuser" {
		loggers.Log.Errorln("Anda tidak punya kewenangan")
		return echo.NewHTTPError(http.StatusForbidden, "Anda tidak punya kewenangan")
	}

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var r models.Kategori

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	item, err := a.M.FetchKategori(int64(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if r.Kode != item.Kode {
		if a.M.CheckKategoriKode(r.Kode) {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Kode kategori %s sudah ada, cari yang lain", r.Kode))
		}
	}

	err = a.M.UpdateKategori(r, int64(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Updated",
	})
}

func (a *App) RemoveKategori(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)
	role := claims.Role

	if role != "superuser" {
		loggers.Log.Errorln("Anda tidak punya kewenangan")
		return echo.NewHTTPError(http.StatusForbidden, "Anda tidak punya kewenangan")
	}

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = a.M.RemoveKategori(int64(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Removed",
	})
}

func (a *App) GetKategori(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)
	role := claims.Role

	if role != "superuser" {
		loggers.Log.Errorln("Anda tidak punya kewenangan")
		return echo.NewHTTPError(http.StatusForbidden, "Anda tidak punya kewenangan")
	}

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

	items, err := a.M.GetKategori(int64(lastID), limit)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) SearchKategori(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)
	role := claims.Role

	if role != "superuser" {
		loggers.Log.Errorln("Anda tidak punya kewenangan")
		return echo.NewHTTPError(http.StatusForbidden, "Anda tidak punya kewenangan")
	}

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

	items, err := a.M.SearchKategori(int64(lastID), limit, search, filter)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, items)
}
