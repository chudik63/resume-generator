package main

import (
	"homework/resume_manipulator"
	"log"
	"net/http"
)

func main() {
	resume_manipulator.GenerateResume()

	http.HandleFunc("/", requestHandler)

	log.Print("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "public/index.html")
	case "POST":
		resume_manipulator.GenerateResume()
		http.ServeFile(w, r, "public/index.html")
	}

}
