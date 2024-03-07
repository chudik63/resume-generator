package main

import (
	"fmt"
	"homework/resume_manipulator"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", print_resume)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func print_resume(w http.ResponseWriter, _ *http.Request) {
	resume_manipulator.GenerateResume()
	_, err := fmt.Fprintf(w, string(resume_manipulator.GetResume()))
	if err != nil {
		log.Fatal(err)
	}
}
