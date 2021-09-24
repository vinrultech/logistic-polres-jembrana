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

	//api.GET("/fetch/:row_id", app.FetchSuratMasuk)

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
	admin.PUT("/user/reset_password", app.ResetPassword)
	admin.PUT("/user/account", app.ChangeAccount)
	admin.GET("/user/account/:username", app.FetchAccount)
	admin.POST("/user/upload_image/:username", app.UploadImage)
	admin.PUT("/user/remove_image", app.RemoveImage)
	admin.PUT("/user/update/:id", app.UpdateUser)
	admin.DELETE("/user/remove/:id", app.RemoveUser)
	admin.GET("/user", app.GetsUser)
	admin.GET("/user/search", app.SearchUser)

	//kategori
	admin.POST("/kategori/create", app.CreateKategori)
	admin.PUT("/kategori/update/:id", app.UpdateKategori)
	admin.GET("/kategori", app.GetKategori)
	admin.GET("/kategori/search", app.SearchKategori)
	admin.GET("/kategori/all", app.AllKategori)
	admin.DELETE("/kategori/remove/:id", app.RemoveKategori)

	//files
	admin.POST("/files/upload", app.UploadFiles)
	admin.POST("/files/remove", app.RemoveFile)

	//surat masuk
	admin.POST("/surat_masuk/create", app.CreateSuratMasuk)
	admin.PUT("/surat_masuk/update/:row_id", app.UpdateSuratMasuk)
	admin.GET("/surat_masuk", app.GetSuratMasuk)
	admin.GET("/surat_masuk/:row_id", app.FetchSuratMasuk)
	admin.GET("/surat_masuk/search", app.SearchSuratMasuk)
	admin.DELETE("/surat_masuk/remove/:row_id", app.RemoveSuratMasuk)

	//surat keluar
	admin.POST("/surat_keluar/create", app.CreateSuratKeluar)
	admin.PUT("/surat_keluar/update/:row_id", app.UpdateSuratKeluar)
	admin.GET("/surat_keluar", app.GetSuratKeluar)
	admin.GET("/surat_keluar/:row_id", app.FetchSuratKeluar)
	admin.GET("/surat_keluar/search", app.SearchSuratKeluar)
	admin.DELETE("/surat_keluar/remove/:row_id", app.RemoveSuratKeluar)

	//juknis
	admin.POST("/juknis/create", app.CreateJuknis)
	admin.PUT("/juknis/update/:row_id", app.UpdateJuknis)
	admin.GET("/juknis", app.GetJuknis)
	admin.GET("/juknis/:row_id", app.FetchJuknis)
	admin.GET("/juknis/search", app.SearchJuknis)
	admin.DELETE("/juknis/remove/:row_id", app.RemoveJuknis)

	//barang
	admin.POST("/barang/create", app.CreateBarang)
	admin.PUT("/barang/update/:row_id", app.UpdateBarang)
	admin.GET("/barang", app.GetBarang)
	admin.GET("/barang/:row_id", app.FetchBarang)
	admin.GET("/barang/search", app.SearchBarang)
	admin.DELETE("/barang/remove/:row_id", app.RemoveBarang)

	//unit_kerja
	admin.POST("/unit_kerja/create", app.CreateUnitKerja)
	admin.PUT("/unit_kerja/update/:id", app.UpdateUnitKerja)
	admin.GET("/unit_kerja", app.GetUnitKerja)
	admin.GET("/unit_kerja/all", app.AllUnitKerja)
	admin.GET("/unit_kerja/search", app.SearchUnitKerja)
	admin.DELETE("/unit_kerja/remove/:id", app.RemoveUnitKerja)

	//metric
	admin.POST("/metric/create", app.CreateMetric)
	admin.PUT("/metric/update/:id", app.UpdateMetric)
	admin.GET("/metric", app.GetMetric)
	admin.GET("/metric/search", app.SearchMetric)
	admin.GET("/metric/all", app.AllMetric)
	admin.DELETE("/metric/remove/:id", app.RemoveMetric)
}
