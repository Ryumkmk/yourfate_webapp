package controllers

import (
	"net/http"
	"yourfate_webapp/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "", "layout", "top_navbar", "top")
}

func fate(w http.ResponseWriter, r *http.Request) {
	h := models.GetHoroscope(models.GetTime())

	generateHTML(w, h, "layout", "fate_navbar", "fate")

}
