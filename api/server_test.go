package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"riddle_with_numbers/riddle"
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

	expected := riddle.GetExampleResult()

	// Convert matrix to JSON
	jsonMatrix, _ := json.Marshal(matrix)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/solve", bytes.NewReader(jsonMatrix))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var result [][]riddle.Cell

	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err)

	assert.True(t, riddle.CompareMatrices(expected, result))

}
