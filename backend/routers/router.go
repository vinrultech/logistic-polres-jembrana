package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/controllers"
)

func SetupRouters(e *echo.Echo, app *controllers.App) {

	api := e.Group("/api")
	api.GET("/test", app.TestInsertRow)

	//
	api.POST("/user/login", app.Login)

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &constants.JwtTandaTanganOnlineClaims{},
		SigningKey: []byte(constants.Key),
	}

	admin := api.Group("/admin", middleware.JWTWithConfig(config))
	admin.POST("/user/create", app.CreateUser)
	admin.PUT("/user/change", app.ChangePassword)
	admin.PUT("/user/account", app.ChangeAccount)
	admin.GET("/user/account/:username", app.FetchAccount)
	admin.POST("/user/upload_image/:username", app.UploadImage)
	admin.PUT("/user/remove_image", app.RemoveImage)

	//kategori
	admin.POST("/kategori/create", app.CreateKategori)
	admin.PUT("/kategori/update/:id", app.UpdateKategori)
	admin.GET("/kategori", app.GetKategori)
	admin.GET("/kategori/search", app.SearchKategori)
	admin.DELETE("/kategori/remove/:id", app.RemoveKategori)

	//files
	admin.POST("/files/upload", app.UploadFiles)
	admin.POST("/files/remove", app.RemoveFile)
}
