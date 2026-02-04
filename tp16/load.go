package main

import (
	"encoding/json"
	"net/http"

	"github.com/shirou/gopsutil/v4/load"
)

func LoadHandler(w http.ResponseWriter, req *http.Request) {
	loadAvg, err := load.Avg()
	if err != nil {
		http.Error(w, "erreur interne", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loadAvg)
}
