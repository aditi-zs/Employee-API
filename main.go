package main

import (
	"employee/emp"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/emp", emp.GetEmployeeData).Methods("GET")
	router.HandleFunc("/emp/{id}", emp.GetOneEmployeeData).Methods("GET")
	router.HandleFunc("/emp", emp.PostEmployeeData).Methods("POST")
	fmt.Println(("server at port 8080"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
