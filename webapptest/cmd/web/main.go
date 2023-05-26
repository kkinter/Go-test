package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	Session *scs.SessionManager
}

func main() {
	app := application{}

	// get a session manager
	app.Session = getSession()

	log.Println("Starting server on port 8080")

	err := http.ListenAndServe(":8080", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
