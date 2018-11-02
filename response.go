package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, sc int, msg string) {
	w.WriteHeader(sc)
	io.WriteString(w, msg)
	//http.Error(w, msg, sc)
	return
}
