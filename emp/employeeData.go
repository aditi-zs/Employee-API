package emp

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	//"log"
	"net/http"
)

// Employee Struct
type Employee struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

// slice to store all employees
type allEmployees []Employee

var employees = allEmployees{{"1", "Aditi", 22, "UP"}}

// Get all employees data
func GetEmployeeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
}

func GetOneEmployeeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	empID := mux.Vars(r)["id"]
	for _, val := range employees {
		if (val.ID) == empID {
			json.NewEncoder(w).Encode(val)

		}
	}
	w.WriteHeader(http.StatusOK)

}

// Post new employees data
func PostEmployeeData(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	w.Header().Set("Content-Type", "application/json")
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter data")
	}
	json.Unmarshal(req, &emp)
	employees = append(employees, emp)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}
