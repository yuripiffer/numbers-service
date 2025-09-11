package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type NumbersResponse struct {
	Numbers []int `json:"numbers"`
	Count   int   `json:"count"`
}

func GetOddNumbers(w http.ResponseWriter, r *http.Request) {
	limit := getLimitParam(r)
	numbers := generateOddNumbers(limit)

	response := NumbersResponse{
		Numbers: numbers,
		Count:   len(numbers),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetEvenNumbers(w http.ResponseWriter, r *http.Request) {
	limit := getLimitParam(r)
	numbers := generateEvenNumbers(limit)

	response := NumbersResponse{
		Numbers: numbers,
		Count:   len(numbers),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getLimitParam(r *http.Request) int {
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		return 10 // default
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		return 10 // default on error
	}

	if limit > 1000 {
		return 1000 // max limit
	}

	return limit
}

func generateOddNumbers(limit int) []int {
	numbers := make([]int, 0, limit)
	for i, count := 1, 0; count < limit; i += 2 {
		numbers = append(numbers, i)
		count++
	}
	return numbers
}

func generateEvenNumbers(limit int) []int {
	numbers := make([]int, 0, limit)
	for i, count := 2, 0; count < limit; i += 2 {
		numbers = append(numbers, i)
		count++
	}
	return numbers
}
