package main

import (
	"github.com/hSATAC/go-zxing-qrencoder/qrencode"
	"image/png"
	"net/http"
	"strconv"
)

func getQR(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Header().Set("Status", "500 Internal Server Error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	size := r.Form.Get("size")
	if size == "" {
		size = "10"
	}
	txt := r.Form.Get("text")
	grid, err := qrencode.Encode(txt, qrencode.ECLevelQ)
	if err != nil {
		w.Header().Set("Status", "500 Internal Server Error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	isize, err := strconv.Atoi(size)
	if err != nil {
		w.Header().Set("Status", "500 Internal Server Error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	png.Encode(w, grid.Image(isize))
}
