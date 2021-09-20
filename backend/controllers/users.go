package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/models"
)

func (a *App) Login(c echo.Context) error {
	var u struct {
		Username string `json:"username" validate:"required,gte=5"`
		Password string `json:"password" validate:"required"`
	}

	if err := c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(&u); err != nil {
		return err
	}

	user, err := a.M.GetUser(u.Username)

	if err != nil {

		return echo.NewHTTPError(http.StatusInternalServerError, "Username tidak ada")
	}

	err = constants.CheckPasswordHash(u.Password, user.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Password salah")
	}

	claims := &constants.JwtTandaTanganOnlineClaims{
		user.Username,
		user.Role,
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(constants.Key))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, constants.H{
		"message":  "Ok",
		"token":    t,
		"username": user.Username,
		"nama":     user.Nama,
		"foto":     user.Foto,
		"role":     user.Role,
	})
}

func (a *App) ChangePassword(c echo.Context) error {
	var u struct {
		Username        string `json:"username" validate:"required"`
		Password        string `json:"password" validate:"required"`
		ConfirmPassword string `json:"confirm_password" validate:"required"`
		OldPassword     string `json:"old_password" validate:"required"`
	}

	if err := c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(&u); err != nil {
		return err
	}

	if u.Password != u.ConfirmPassword {
		return echo.NewHTTPError(http.StatusInternalServerError, "Password tidak sama dengan konfirmasi password")
	}

	user, err := a.M.GetUser(u.Username)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Username tidak ada")
	}

	err = constants.CheckPasswordHash(u.OldPassword, user.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Password lama salah")
	}

	err = a.M.ChangePassword(u.Password, user.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error saat merubah password")
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Ok",
	})
}

func (a *App) CreateUser(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*constants.JwtTandaTanganOnlineClaims)
	role := claims.Role

	if role != "superuser" {
		loggers.Log.Errorln("Anda tidak punya kewenangan")
		return echo.NewHTTPError(http.StatusForbidden, "Anda tidak punya kewenangan")
	}

	var r models.User

	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(&r); err != nil {
		return err
	}

	if !constants.ContainsString([]string{"superuser", "staff"}, r.Role) {
		return echo.NewHTTPError(http.StatusInternalServerError, "Role tidak ada")
	}

	if a.M.CheckUser(r.Username) {
		return echo.NewHTTPError(http.StatusInternalServerError, "Username sudah ada, cari yang lain")
	}

	err := a.M.CreateUser(r)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error create")
	}

	return c.JSON(http.StatusCreated, constants.H{
		"message": "Created",
	})
}

func (a *App) FetchAccount(c echo.Context) error {
	username := c.Param("username")

	user, err := a.M.GetUser(username)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Username tidak ada")
	}

	account, err := a.M.FetchAccount(user.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error mendapatkan account")
	}

	return c.JSON(http.StatusOK, constants.H{
		"nama": account.Nama,
		"foto": account.Foto,
		"hp":   account.Hp,
	})
}

func (a *App) ChangeAccount(c echo.Context) error {

	var u struct {
		Nama     string `json:"nama" validate:"required,gte=3"`
		Hp       string `json:"hp"`
		Username string `json:"username" validate:"required"`
	}

	if err := c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(&u); err != nil {
		return err
	}

	user, err := a.M.GetUser(u.Username)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Username tidak ada")
	}

	err = a.M.ChangeAccount(u.Nama, u.Hp, user.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error merubah account")
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Ok",
		"nama":    u.Nama,
	})
}

func (a *App) UploadImage(c echo.Context) error {

	username := c.Param("username")

	user, err := a.M.GetUser(username)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Username tidak ada")
	}

	form, err := c.MultipartForm()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	image := form.File["image"][0]

	src, err := image.Open()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	defer src.Close()

	imagename := fmt.Sprintf("%s_%s", constants.GetFilename(), image.Filename)

	filename := fmt.Sprintf("upload/%s", imagename)

	dst, err := os.Create(filename)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	a.M.ChangeImage(imagename, user.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error simpan image ke database")
	}

	return c.JSON(http.StatusOK, constants.H{
		"message":  "Ok",
		"filename": imagename,
	})
}

func (a *App) RemoveImage(c echo.Context) error {

	var f struct {
		Filename string `json:"filename" validate:"required"`
		Username string `json:"username" validate:"required"`
	}

	if err := c.Bind(&f); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(&f); err != nil {
		return err
	}

	user, err := a.M.GetUser(f.Username)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Username tidak ada")
	}

	filename := fmt.Sprintf("upload/%s", f.Filename)

	err = os.Remove(filename)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	a.M.ChangeImage("", user.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error remove image dari database")
	}

	return c.JSON(http.StatusOK, constants.H{
		"message": "Ok",
	})
}
