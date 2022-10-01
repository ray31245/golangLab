package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

var count uint64

func set(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		// fmt.Println(err)
		// http.Error(w, http.StatusText(400), http.StatusBadRequest)
	} else {
		count, err = strconv.ParseUint(c.Value, 10, 8)
		if err != nil {
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}
	}
	countStr := strconv.FormatUint(count+1, 10)
	fmt.Println(countStr)
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: countStr,
		Path:  "/",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}

// Using cookies, track how many times a user has been to your website domain.
