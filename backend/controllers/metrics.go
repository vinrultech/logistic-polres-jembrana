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

func (a *App) CreateMetric(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)
	role := claims.Role

	if role != "superuser" {
		loggers.Log.Errorln("Anda tidak punya kewenangan")
		return echo.NewHTTPError(http.StatusForbidden, "Anda tidak punya kewenangan")
	}

	var r models.Metric

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error bind Metric : %v", err))
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	err := a.M.CreateMetric(r)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error create Metric : %v", err))
	}

	return c.JSON(http.StatusCreated, constants.H{
		"message": "Created",
	})
}

func (a *App) UpdateMetric(c echo.Context) error {

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
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error convert id param to int : %v", err))
	}

	var r models.Metric

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error bind Metric : %v", err))
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	err = a.M.UpdateMetric(r, int64(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error update Metric : %v", err))
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Updated",
	})
}

func (a *App) RemoveMetric(c echo.Context) error {

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
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error convert id param to int : %v", err))
	}

	err = a.M.RemoveMetric(int64(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error remove Metric : %v", err))
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Removed",
	})
}

func (a *App) GetMetric(c echo.Context) error {

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
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error Metric convert limit param to int  : %v", err))
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error Metric convert last_id param : %v", err))
	}

	items, err := a.M.GetMetric(int64(lastID), limit)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error get Metric : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) SearchMetric(c echo.Context) error {

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
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error Metric convert limit param to int : %v", err))
	}

	lastID, err := strconv.Atoi(lastIDParam)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error Metric convert last_id param to int : %v", err))
	}

	items, err := a.M.SearchMetric(int64(lastID), limit, search, filter)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error search Metric : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}

func (a *App) AllMetric(c echo.Context) error {

	items, err := a.M.AllMetric()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error all Metric : %v", err))
	}

	return c.JSON(http.StatusOK, items)
}
