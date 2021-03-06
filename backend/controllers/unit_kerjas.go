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

func (a *App) CreateUnitKerja(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)
	role := claims.Role

	if role != "superuser" {
		loggers.Log.Errorln("Anda tidak punya kewenangan")
		return echo.NewHTTPError(http.StatusForbidden, "Anda tidak punya kewenangan")
	}

	var r models.UnitKerja

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error bind unit kerja : %v", err))
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	err := a.M.CreateUnitKerja(r)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error create unit kerja : %v", err))
	}

	return c.JSON(http.StatusCreated, constants.H{
		"message": "Created",
	})
}

func (a *App) UpdateUnitKerja(c echo.Context) error {

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
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error unit kerja comvert id to int : %v", err))
	}

	var r models.UnitKerja

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error bind unit kerja : %v", err))
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	err = a.M.UpdateUnitKerja(r, int64(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error create unit kerja : %v", err))
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Updated",
	})
}

func (a *App) RemoveUnitKerja(c echo.Context) error {

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
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error unit kerja convert id to int : %v", err))
	}

	err = a.M.RemoveUnitKerja(int64(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error remove unit kerja : %v", err))
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Removed",
	})
}

func (a *App) GetUnitKerja(c echo.Context) error {

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
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error unit kerja convert limit param to int : %v", err))
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error unit kerja convert last_id param to int : %v", err))
	}

	items, err := a.M.GetUnitKerja(int64(lastID), limit)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error get unit kerja : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) AllUnitKerja(c echo.Context) error {

	items, err := a.M.AllUnitKerja()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error all unit kerja : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) SearchUnitKerja(c echo.Context) error {

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
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error unit kerja convert limit param to int : %v", err))
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error unit kerja convert last_id param to int : %v", err))
	}

	items, err := a.M.SearchUnitKerja(int64(lastID), limit, search, filter)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error search unit kerja : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}
