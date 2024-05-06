package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/teacher", TeacherPage)
	mux.HandleFunc("/student", StudentPage)
	mux.HandleFunc("/parent", ParentPage)

	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Println(err)
		return
	}
}
