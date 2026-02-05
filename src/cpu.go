package main

import (
	"encoding/json"
	"net/http"

	"github.com/shirou/gopsutil/v4/cpu"
)

func CPUHandler(w http.ResponseWriter, req *http.Request) {
	cpuInfos, err := cpu.Info()
	if err != nil {
		http.Error(w, "erreur interne", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cpuInfos)
}
