package network

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const Authorization = "Authorization"
const TokenIndex = 1

func (network *Network) verifyToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Bearer Token 가져오기.
		token := getToken(context)

		if token == "" {
			context.JSON(http.StatusUnauthorized, nil)
			context.Abort()
		}

		_, err := network.client.VerifyAuth(token)
		if err != nil {
			context.JSON(http.StatusUnauthorized, err.Error())
			context.Abort()
		}
		context.Next()
	}
}

func getToken(c *gin.Context) string {
	var token string

	authToken := c.Request.Header.Get(Authorization)
	authSided := strings.Split(authToken, " ")

	if len(authSided) > 1 {
		token = authSided[TokenIndex]
	}

	return token
}
