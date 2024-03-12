package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"riddle_with_numbers/riddle"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfAlive(t *testing.T) {
	server := NewServer()
	router := server.Router

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}

func TestSolveRiddle(t *testing.T) {
	server := NewServer()
	router := server.Router

	// Example matrix for testing
	matrix := [][]int{
		{4, 2, 4, 8}, {8, 6, 6, 8}, {4, 2, 6, 6}, {2, 2, 6, 6},
	}
	jsonMatrix, _ := json.Marshal(matrix)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/solve", bytes.NewReader(jsonMatrix))
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	count, err := strconv.Atoi(w.Body.String())
	if err != nil {
		return
	}
	assert.True(t, count == 1)

	// GET THE SOLUTION

	w = httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/solution", bytes.NewReader(jsonMatrix))
	assert.NoError(t, err)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var result [][]riddle.Cell

	println(w.Body.String())

	err = json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err)
	expected := riddle.GetExampleResult()
	assert.True(t, riddle.CompareMatrices(expected, result))

}

func TestErrorSolveRiddle(t *testing.T) {
	server := NewServer()
	router := server.Router

	// Example matrix for testing
	matrix := [][]int{
		{4, 2, 4, 8},
	}
	jsonMatrix, _ := json.Marshal(matrix)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/solve", bytes.NewReader(jsonMatrix))
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	w = httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/solution", bytes.NewReader(jsonMatrix))
	assert.NoError(t, err)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

}
