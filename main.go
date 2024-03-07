package main

import (
	"homework/resume_manipulator"
	"log"
	"net/http"
)

func main() {
	resume_manipulator.GenerateResume()
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Print("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
