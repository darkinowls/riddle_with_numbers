package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	db "riddle_with_numbers/db/sqlc"
	"riddle_with_numbers/riddle"
	"riddle_with_numbers/token"
	"riddle_with_numbers/util"
)

type Server struct {
	router     *gin.Engine
	store      db.IStore
	tokenMaker token.ITokenMaker
	config     *util.Config
}

func NewServer(conf *util.Config, store db.IStore) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(conf.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}
	return &Server{store: store, tokenMaker: tokenMaker, config: conf}, nil
}

func (server *Server) Start(address string) error {
	server.router = server.buildRoutes()
	return server.router.Run(address)
}

func (server *Server) buildRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("ping", checkIfAlive)
	r.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ga := r.Group("/auth")
	ga.POST("create", server.createUser)
	ga.POST("login", server.loginUser)

	gr := r.Group("/").Use(authMiddleware(server.tokenMaker))
	gr.GET("solution", getResults)
	gr.POST("solve", solveRiddle)
	return r
}

type errorRes struct {
	Error string `json:"error"`
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// ///////////////////////////////////////////////////////////////////////////////////////////////

// @Summary ping example
// @Description do ping
// @ID ping-example
// @Accept json
// @Produce json
// @Success 200 {object} string "pong"
// @router /ping [get]
// @Security BearerAuth
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
// @router /solution [get]
// @Security BearerAuth
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
// @router /solve [post]
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
