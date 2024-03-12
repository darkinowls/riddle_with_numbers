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

// @Summary solve riddle
// @Description solve riddle
// @ID solve-riddle
// @Accept json
// @Produce json
// @Param matrix body [][]int true "matrix"
// @Success 200 {object} [][]riddle.Cell "solved matrix"
// @Router /solve [post]
func solveRiddle(c *gin.Context) {
	var matrix [][]int
	err := c.BindJSON(&matrix)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	TranslateToCells := riddle.TranslateToCells(matrix)
	result := riddle.SolveMatrix(TranslateToCells)
	c.JSON(http.StatusOK, result)
}
