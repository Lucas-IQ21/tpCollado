package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/shirou/gopsutil/v4/process"
)

type Proc struct {
	Pid    int32   `json:"pid"`
	Name   string  `json:"name"`
	User   string  `json:"user"`
	CPU    float64 `json:"cpu"`
	Memory float32 `json:"memory"`
	Status string  `json:"status"`
}

func PSHandler(w http.ResponseWriter, req *http.Request) {
	var myProcs []Proc
	procs, err := process.Processes()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, proc := range procs {
		var myProc Proc
		myProc.Pid = proc.Pid
		myProc.Name, _ = proc.Name()
		myProc.User, _ = proc.Username()
		myProc.CPU, _ = proc.CPUPercent()
		myProc.Memory, _ = proc.MemoryPercent()
		status, _ := proc.Status()
		myProc.Status = status[0]
		myProcs = append(myProcs, myProc)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myProcs)
}

func ReadProc(user string) (*[]Proc, error) {
	var myProcs []Proc
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}
	for _, proc := range procs {
		var myProc Proc
		myProc.User, _ = proc.Username()
		if user == "" || user == myProc.User {
			myProc.Pid = proc.Pid
			myProc.Name, _ = proc.Name()
			myProc.CPU, _ = proc.CPUPercent()
			myProc.Memory, _ = proc.MemoryPercent()
			status, _ := proc.Status()
			myProc.Status = status[0]
			myProcs = append(myProcs, myProc)
		}
	}
	return &myProcs, nil
}

func PSUserHandler(w http.ResponseWriter, req *http.Request) {
	user := req.URL.Query().Get("user")
	procs, err := ReadProc(user)
	if err != nil {
		http.Error(w, "erreur interne", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(procs)
}
