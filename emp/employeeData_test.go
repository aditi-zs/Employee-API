package emp

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEmployeeData(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expRes      []Employee
		statusCode  int
	}{
		{"All entries are present",
			"",
			[]Employee{
				{"1", "Aditi", 22, "UP"},
			},
			200,
		},
	}
	//comment
	for _, tc := range tests {
		req, err := http.NewRequest("GET", "/emp", nil)
		if err != nil {
			t.Errorf(err.Error()) //err.Error() will return a string
		}
		//response recorder
		resRec := httptest.NewRecorder()

		GetEmployeeData(resRec, req)
		var val []Employee
		_ = json.Unmarshal(resRec.Body.Bytes(), &val) //json to go

		assert.Equal(t, tc.statusCode, resRec.Code)
		assert.Equal(t, tc.expRes, val)
	}
}

func TestPostEmployeeData(t *testing.T) {
	tests := []struct {
		description string
		input       Employee
		expRes      Employee
		statusCode  int
	}{
		{"All entries are present",
			Employee{
				"1", "Aditi", 22, "UP",
			},
			Employee{
				"1", "Aditi", 22, "UP",
			},
			201,
		},
	}

	for _, tc := range tests {
		val, _ := json.Marshal(tc.input) //go to json
		req, err := http.NewRequest("POST", "/emp", bytes.NewReader(val))
		if err != nil {
			t.Errorf(err.Error())
		}
		//response recorder
		resRec := httptest.NewRecorder()
		PostEmployeeData(resRec, req)
		var actRes Employee
		_ = json.Unmarshal(resRec.Body.Bytes(), &actRes) //json to go
		assert.Equal(t, tc.statusCode, resRec.Code)
		assert.Equal(t, tc.expRes, actRes)
	}
}
