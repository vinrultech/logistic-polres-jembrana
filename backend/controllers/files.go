package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
)

func (a *App) UploadFiles(c echo.Context) error {

	form, err := c.MultipartForm()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	file := form.File["file"][0]
	fileId := c.FormValue("file_id")
	fileName := c.FormValue("filename")
	status := c.FormValue("status")

	src, err := file.Open()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	defer src.Close()

	fileurl := fmt.Sprintf("files/%s_%s", fileId, fileName)

	dst, err := os.Create(fileurl)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		loggers.Log.Errorln(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, constants.H{
		"filename": fileName,
		"url":      fileurl,
		"file_id":  fileId,
		"status":   status,
	})
}

func (a *App) RemoveFile(c echo.Context) error {

	var f struct {
		Filename string `json:"filename" validate:"required"`
		FileID   string `json:"file_id" validate:"required"`
		Status   string `json:"status" validate:"required"`
	}

	if err := c.Bind(&f); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(&f); err != nil {
		return err
	}

	if f.Status == "edit" {
		err := a.M.RemoveFile(f.FileID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	filename := fmt.Sprintf("files/%s_%s", f.FileID, f.Filename)

	err := os.Remove(filename)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, constants.H{
		"filename": f.Filename,
		"url":      filename,
		"file_id":  f.FileID,
		"status":   f.Status,
	})
}
