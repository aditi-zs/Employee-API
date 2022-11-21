package emp

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetEmployeeData(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expRes      string
		statusCode  int
	}{
		{"All entries are present", "",
			`[{"id":"1","name":"Aditi","age":22,"address":"UP"}]`, 200,
		},
	}
	for _, tc := range tests {
		req, err := http.NewRequest("GET", "/get", nil)
		if err != nil {
			t.Fatal(err)
		}
		//response recorder
		resRec := httptest.NewRecorder()

		GetEmployeeData(resRec, req)

		if status := resRec.Code; status != http.StatusOK {
			t.Errorf("wrong status code returned: got %v want %v", status, http.StatusOK)
		}
		//expected := `[{"id":"1","name":"Aditi","age":22,"address":"UP"}]`
		if strings.TrimSpace(resRec.Body.String()) != tc.expRes {
			t.Errorf("handler returned unexpected body: got %v want [%v]", resRec.Body.String(), tc.expRes)
		}
	}
}

func TestPostEmployeeData(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expRes      string
		statusCode  int
	}{
		{"All entries are present", `{"id":"1","name":"Aditi","age":22,"address":"UP"}`,
			`{"id":"1","name":"Aditi","age":22,"address":"UP"}`, 201,
		},
		{"entries are missing", `{"id":"1","name":"Aditi"}`,
			`{"id":"1","name":"Aditi","age":0,"address":""}`, 201,
		},
	}
	for _, tc := range tests {
		req, err := http.NewRequest("POST", "/post", strings.NewReader(tc.input))
		if err != nil {
			t.Fatal(err)
		}
		//response recorder
		resRec := httptest.NewRecorder()
		PostEmployeeData(resRec, req)
		if tc.statusCode != resRec.Code {
			t.Errorf("handler returned unexpected code: got (%v) want [%v]", resRec.Code, tc.statusCode)
		}
		if strings.TrimSpace(resRec.Body.String()) != tc.expRes {
			t.Errorf("handler returned unexpected body: got (%v) want [%v]", resRec.Body.String(), tc.expRes)
		}
	}
}
