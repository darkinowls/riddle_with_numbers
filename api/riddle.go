package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"riddle_with_numbers/riddle"
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
