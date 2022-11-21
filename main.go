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
	router.HandleFunc("/get", emp.GetEmployeeData).Methods("GET")
	router.HandleFunc("/get/{id}", emp.GetOneEmployeeData).Methods("GET")
	router.HandleFunc("/post", emp.PostEmployeeData).Methods("POST")
	fmt.Println(("server at port 8080"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
