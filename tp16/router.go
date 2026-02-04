package main

import (
	"net/http"
)

func router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /cpu", CPUHandler)
	mux.HandleFunc("GET /ps", PSHandler)
	mux.HandleFunc("GET /ps/{user}", PSUserHandler)
	mux.HandleFunc("GET /load", LoadHandler)
	mux.HandleFunc("GET /disk", DiskHandler)
	return mux
}
