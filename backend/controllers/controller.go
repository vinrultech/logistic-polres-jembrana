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

func paginator(count int, firstNo int, lastNo int, maxID int, minID int, limit int, action string) (bool, bool) {
	prevShow := true
	nextShow := true

	if count <= 0 {
		prevShow = false
		nextShow = false
	} else {

		if minID == lastNo {
			nextShow = false
		}

		if maxID == firstNo {
			prevShow = false
		}

		if action == "first" {
			prevShow = false
		}

		if count < limit {
			nextShow = false
		}
	}

	return prevShow, nextShow
}
