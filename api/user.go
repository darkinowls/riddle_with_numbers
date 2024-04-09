package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"riddle_with_numbers/db/sqlc"
	"riddle_with_numbers/util"
	"time"
)

type createUserRequest struct {
	Password string `json:"password" binding:"required,min=6,max=32"`
	Email    string `json:"email" binding:"required,email"`
}

// DTO
type userResponse struct {
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type loginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginUserResponse struct {
	AccessToken string        `json:"access_token"`
	User        *userResponse `json:"user"`
}

func newUserResponse(user db.User) *userResponse {
	return &userResponse{
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

// @Summary create user
// @Description create user
// @ID create-user
// @Accept json
// @Produce json
// @Param user body createUserRequest true "user"
// @Success 200 {object} userResponse "user"
// @Failure 400 {object} errorRes "error"
// @Failure 500 {object} errorRes "error"
// @Router /auth/create [post]
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		HashedPassword: hashedPassword,
		Email:          req.Email,
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pErr, ok := err.(*pq.Error); ok {
			switch pErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newUserResponse(user))
}

// @Summary login user
// @Description login user
// @ID login-user
// @Accept json
// @Produce json
// @Param user body loginUserRequest true "user"
// @Success 200 {object} loginUserResponse "user"
// @Failure 400 {object} errorRes "error"
// @Failure 404 {object} errorRes "error"
// @Failure 401 {object} errorRes "error"
// @Failure 500 {object} errorRes "error"
// @Router /auth/login [post]
func (s *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := s.store.GetUser(ctx, req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := s.tokenMaker.CreateToken(user.Email, s.config.TokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	})
}
