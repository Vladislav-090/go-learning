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
		fmt.Fprintln(w, "empty a")
		return
	}
	if bString == "" {
		fmt.Fprintln(w, "empty b")
		return
	}
	a, err := strconv.Atoi(aString)
	if err != nil {
		fmt.Fprintln(w, "strToInt error!", a)
		return
	}
	b, err := strconv.Atoi(bString)
	if err != nil {
		fmt.Fprintln(w, "strToInt error!", b)
		return
	}
	op := r.URL.Query().Get("op")
	var result int
	switch op {
	case "sum":
		result = a + b
	case "sub":
		result = a - b
	case "mul":
		result = a * b
	case "div":
		if b == 0 {
			fmt.Fprintln(w, "division by zero")
			return
		}
		result = a / b

	default:
		fmt.Fprintln(w, "Uknown operation")
		return
	}

	fmt.Fprintln(w, "Result:", result)
}

func main() {
	http.HandleFunc("/calc", calcHandler)
	http.ListenAndServe(":8080", nil)
}
