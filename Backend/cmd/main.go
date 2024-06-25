package main

import (
	"Backend/api/greet"
	"Backend/api/login"
	"Backend/api/songs"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /songs/next", songs.NextSongHandler)
	router.HandleFunc("POST /songs/add", songs.AddSongHandler)

	//Default Route to Hello World function
	router.HandleFunc("/{name}", greet.Greet)
	router.HandleFunc("/login", login.Login)
	router.HandleFunc("/login-sumbit", login.LoginSumbit)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server listening on port :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}
