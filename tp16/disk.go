package main

import (
	"encoding/json"
	"net/http"

	"github.com/shirou/gopsutil/v4/disk"
)

func DiskHandler(w http.ResponseWriter, req *http.Request) {
	diskAvg, err := disk.Usage("/")
	if err != nil {
		http.Error(w, "erreur interne", http.StatusInternalServerError)
		return
	}
	diskAvg.UsedPercent = diskAvg.UsedPercent * 100
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(diskAvg)
}
