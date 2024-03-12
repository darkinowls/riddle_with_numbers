package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"riddle_with_numbers/riddle"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()
	buildRoutes(r)
	return &Server{Router: r}
}

func buildRoutes(r *gin.Engine) {
	r.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("ping", checkIfAlive)
	r.GET("solution", getResults)
	r.POST("solve", solveRiddle)
}

// ///////////////////////////////////////////////////////////////////////////////////////////////

// @Summary ping example
// @Description do ping
// @ID ping-example
// @Accept json
// @Produce json
// @Success 200 {object} string "pong"
// @Router /ping [get]
func checkIfAlive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

var results [][][]riddle.Cell

// @Summary get next solution
// @Description get next solution
// @ID get-next-solution
// @Success 200 {object} [][]riddle.Cell  "solved matrix"
// @Router /solution [get]
func getResults(c *gin.Context) {
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
// @Router /solve [post]
func solveRiddle(c *gin.Context) {
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
