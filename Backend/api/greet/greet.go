package greet

import (
	"fmt"
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	userName := r.PathValue("name")
	if userName != "" {
		fmt.Fprintf(w, "Hello  %s", userName)
	} else {
		fmt.Fprintf(w, "Hello World")
	}
}
