package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	db "riddle_with_numbers/db/sqlc"
	"riddle_with_numbers/riddle"
	"strconv"

	"github.com/gin-gonic/gin"
)

var results [][][]riddle.Cell

// @Summary get next solution
// @Description get next solution
// @ID get-next-solution
// @Success 200 {object} [][]riddle.Cell  "solved matrix"
// @router /solution [get]
// @Security BearerAuth
func (server *Server) getResults(c *gin.Context) {
	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no results"})
		return
	}
	c.JSON(http.StatusOK, (results)[0])
	results = results[1:]
}

// @Summary solve riddle
// @Description solve riddle
// @ID solve-riddle
// @Accept json
// @Produce json
// @Param matrix body [][]int true "matrix"
// @Success 200 {integer} int "number of solutions"
// @router /solve [post]
// @Security BearerAuth
func (server *Server) solveRiddle(c *gin.Context) {
	var matrix [][]int
	err := c.BindJSON(&matrix)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = riddle.ValidateInputMatrix(&matrix)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	TranslateToCells := riddle.TranslateToCells(matrix)
	results = riddle.SolveMatrix(TranslateToCells)
	c.JSON(http.StatusOK, len(results))
}

// @Success 200 {object} RawMessage "solution"

// @Summary get solution by id
// @Description get solution by id
// @ID get-solution-by-id
// @Param id path int true "solution id"
// @router /condition/{id} [get]
func (server *Server) getSolutionById(c *gin.Context) {
	idStr := c.Param("id")
	println(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Handle error if the ID cannot be converted to an integer
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Now you have the ID, you can use it to fetch the solution from your database or any storage

	// For example:
	solution, err := server.store.GetSolution(context.Background(), int64(id))

	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	if err != nil {
		// Handle error if there's an issue retrieving the solution
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve solution"})
		return
	}

	// Return the solution as JSON
	c.JSON(http.StatusOK, solution.Condition)
}

// @Summary generate conditions
// @Description generate conditions for matrix num x num
// @ID generate-solutions
// @Param num path int true "number of conditions"
// @router /generate/{num} [post]
func (server *Server) generateConditions(c *gin.Context) {

	numStr := c.Param("num")
	num, err := strconv.Atoi(numStr)
	if err != nil || num > 3 || num <= 1 {
		c.JSON(http.StatusBadRequest, errorRes{Error: "Invalid number for the command"})
		return
	}
	err = server.store.DeleteAllSolutions(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorRes{Error: "Failed to delete all solutions"})
		return
	}
	conditions := riddle.GenerateAllMatrices(num)
	for _, con := range conditions {
		jsonData, err := json.Marshal(con)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorRes{Error: "Failed to generate conditions"})
			return
		}
		_, err = server.store.CreateSolution(context.Background(),
			db.CreateSolutionParams{
				Condition: jsonData,
			},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorRes{Error: "Failed to save conditions"})
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d Solutions generated", len(conditions))})

}
