package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOddNumbers(t *testing.T) {
	req, err := http.NewRequest("GET", "/odd?limit=5", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOddNumbers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response NumbersResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int{1, 3, 5, 7, 9}
	if len(response.Numbers) != len(expected) {
		t.Errorf("handler returned unexpected count: got %v want %v", len(response.Numbers), len(expected))
	}

	for i, num := range response.Numbers {
		if num != expected[i] {
			t.Errorf("handler returned unexpected number at index %d: got %v want %v", i, num, expected[i])
		}
	}
}

func TestGetEvenNumbers(t *testing.T) {
	req, err := http.NewRequest("GET", "/even?limit=5", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEvenNumbers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response NumbersResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int{2, 4, 6, 8, 10}
	if len(response.Numbers) != len(expected) {
		t.Errorf("handler returned unexpected count: got %v want %v", len(response.Numbers), len(expected))
	}

	for i, num := range response.Numbers {
		if num != expected[i] {
			t.Errorf("handler returned unexpected number at index %d: got %v want %v", i, num, expected[i])
		}
	}
}

func TestGetOddNumbersDefaultLimit(t *testing.T) {
	req, err := http.NewRequest("GET", "/odd", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOddNumbers)

	handler.ServeHTTP(rr, req)

	var response NumbersResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Count != 10 {
		t.Errorf("handler returned unexpected default count: got %v want %v", response.Count, 10)
	}
}

func TestGetEvenNumbersMaxLimit(t *testing.T) {
	req, err := http.NewRequest("GET", "/even?limit=2000", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEvenNumbers)

	handler.ServeHTTP(rr, req)

	var response NumbersResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Count != 1000 {
		t.Errorf("handler returned unexpected max limit count: got %v want %v", response.Count, 1000)
	}
}

func TestGetOddNumbersInvalidLimit(t *testing.T) {
	req, err := http.NewRequest("GET", "/odd?limit=invalid", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOddNumbers)

	handler.ServeHTTP(rr, req)

	var response NumbersResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Count != 10 {
		t.Errorf("handler returned unexpected count for invalid limit: got %v want %v", response.Count, 10)
	}
}
