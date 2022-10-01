package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/dog", dog)
	http.Handle("/everything/", http.StripPrefix("/everything", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/everything/toby.jpg">`)
}
