package main

import (
	"fmt"
	"homework/resume_manipulator"
	"net/http"
)

func main() {
	http.HandleFunc("/", print_resume)
	http.ListenAndServe(":8080", nil)
}

func print_resume(w http.ResponseWriter, _ *http.Request) {
	resume_manipulator.GenerateResume()
	fmt.Fprintf(w, string(resume_manipulator.GetResume()))
}
