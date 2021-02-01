package api

import (
	"encoding/json"
	"mutant/app"
	"mutant/db"
	"mutant/models"
	"net/http"
)

//IsMutant get request
func IsMutant(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")

	var mutant models.MutantModel

	err := json.NewDecoder(r.Body).Decode((&mutant))

	if err != nil {

		http.Error(w, "invalid json object format for dna , please check ", 400)

	}
	result, err := app.IsMutantLogic(&mutant)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	if result {
		w.WriteHeader((http.StatusOK))
	} else {
		w.WriteHeader((http.StatusForbidden))
	}

	json.NewEncoder(w).Encode(result)
}

//GetStats Route
func GetStats(w http.ResponseWriter, r *http.Request) {

	stats, err := db.GetStatsBD()

	if err != nil {

		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")

	w.WriteHeader((http.StatusCreated))

	json.NewEncoder(w).Encode(stats)

}
