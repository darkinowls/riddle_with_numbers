package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	token "riddle_with_numbers/token"
)

const (
	authorizationHeaderKey = "Authorization"
	bearerPrefix           = "Bearer "
	authPayloadKey         = "authorization_payload_key"
)

func authMiddleware(maker token.ITokenMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.Request.Header.Get(authorizationHeaderKey)
		if len(value) < len(bearerPrefix) {
			c.AbortWithStatusJSON(401, errorResponse(errors.New("authorization header is not provided")))
			return
		}
		tokenString := value[len(bearerPrefix):]
		verifyToken, err := maker.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, errorResponse(errors.New("unauthorized")))
			return
		}
		// TODO: check if the user exists

		c.Set(authPayloadKey, verifyToken)
		c.Next()
		// Reqeust is passed to the next middleware
	}
}
