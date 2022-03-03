package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return

	}
	w.Write([]byte("Hello world"))
}