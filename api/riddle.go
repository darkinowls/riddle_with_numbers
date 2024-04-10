package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"riddle_with_numbers/riddle"
	"strconv"
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
// @Security BearerAuth
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
	if err != nil {
		// Handle error if there's an issue retrieving the solution
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve solution"})
		return
	}

	// Return the solution as JSON
	c.JSON(http.StatusOK, solution.Condition)
}
