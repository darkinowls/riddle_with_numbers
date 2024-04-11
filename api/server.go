package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	db "riddle_with_numbers/db/sqlc"
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

func NewTestServer() (*Server, error) {
	conf, err := util.LoadConfig("..")
	if err != nil {
		log.Fatal("cannot load config:", err)
		return nil, err
	}
	dbCon, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
		return nil, err
	}
	server, err := NewServer(&conf, db.NewStore(dbCon))
	if err != nil {
		fmt.Println("Error creating server: ", err.Error())
		return nil, err
	}
	server.router = server.buildRoutes()
	return server, nil
}

func (server *Server) Start(address string) error {
	server.router = server.buildRoutes()
	return server.router.Run(address)
}

func (server *Server) buildRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("ping", checkIfAlive)
	r.GET("/condition/:id", server.getSolutionById)
	r.POST("/generate", server.generateConditions)
	r.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ga := r.Group("/auth")
	ga.POST("create", server.createUser)
	ga.POST("login", server.loginUser)

	gr := r.Group("/").Use(authMiddleware(server.tokenMaker))
	gr.GET("solution", server.getResults)
	gr.POST("solve", server.solveRiddle)
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
func checkIfAlive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
