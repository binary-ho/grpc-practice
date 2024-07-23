package network

import (
	"github.com/gin-gonic/gin"
	"grpc-practice/types"
	"net/http"
)

const Login = "/login"

func (network *Network) login() (string, gin.HandlerFunc) {
	return Login, network.loginHandler
}

func (network *Network) loginHandler(context *gin.Context) {
	// Auth Data 생성 필요
	var loginRequest types.LoginRequest

	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	} else if res, err := network.service.CreateAuth(loginRequest.Name); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	} else {
		context.JSON(http.StatusOK, res)
	}
}

const Verify = "/verify"

func (network *Network) verify() (string, gin.HandlerFunc, gin.HandlerFunc) {
	return Verify, network.verifyToken(), network.verifyHandler
}

func (network *Network) verifyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "추카푸카해요~")
}
