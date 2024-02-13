package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/amalia-fadilastuti/sbx3-golang-level4/sum"
)

func main() {
	// fmt.Println("This is not a web server")
	// http.HandleFunc("/sum", sumHandler)
	// http.ListenAndServe(":8080", nil)

	http.ListenAndServe(":8080", handler())
}

func handler() http.Handler {
	s := http.NewServeMux()
	s.HandleFunc("/sum", sumHandler)
	return s
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	a := r.FormValue("a")
	b := r.FormValue("b")
	va, _ := strconv.Atoi(a)
	vb, _ := strconv.Atoi(b)
	s := sum.Ints(va, vb)

	w.Write([]byte(fmt.Sprintf("%d", s)))
}
