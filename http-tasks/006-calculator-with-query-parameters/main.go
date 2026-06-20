package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calcHandler(w http.ResponseWriter, r *http.Request) {
	aString := r.URL.Query().Get("a")
	bString := r.URL.Query().Get("b")

	if aString == "" {
		fmt.Fprintln(w, "a is empty")
		return
	}
	if bString == "" {
		fmt.Fprintln(w, "b is empty")
		return
	}

	a, err := strconv.Atoi(aString)
	if err != nil {
		fmt.Fprintln(w, "invalid number")
		return
	}

	b, err := strconv.Atoi(bString)
	if err != nil {
		fmt.Fprintln(w, "invalid number")
		return
	}

	sum := a + b

	fmt.Fprintln(w, "Summ of numbers is:", sum)

}

func main() {
	http.HandleFunc("/calc", calcHandler)
	http.ListenAndServe(":8080", nil)
}
