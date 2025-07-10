package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fealtyx/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/students", handlers.CreateStudent).Methods("POST")
	r.HandleFunc("/students", handlers.GetAllStudents).Methods("GET")
	r.HandleFunc("/students/{id}", handlers.GetStudentByID).Methods("GET")
	r.HandleFunc("/students/{id}", handlers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", handlers.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/students/{id}/summary", handlers.GetStudentSummary).Methods("GET")

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
