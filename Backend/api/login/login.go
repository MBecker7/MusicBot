package login

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

var userDB = map[string]string{
	"Marc": "Passwort123",
}

func Login(w http.ResponseWriter, r *http.Request) {
	fileName, err := filepath.Abs("../../Frontend/login.html")
	if err != nil {
		fmt.Println("Error with filepath", err)
		return
	}

	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}
	err = t.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		fmt.Println("Error when executing template", err)
	}

}

func LoginSumbit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if userDB[username] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Logged in successfully, Hello %s", username)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User or Password do not match!")
	}
}
