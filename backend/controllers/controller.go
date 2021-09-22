package controllers

import (
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/models"
)

type App struct {
	M *models.Model
}

func New(m *models.Model) *App {
	a := new(App)
	a.M = m

	return a
}
