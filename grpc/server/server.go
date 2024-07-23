package server

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-practice/config"
	paseto "grpc-practice/grpc/pasto"
	auth "grpc-practice/grpc/proto"
	"log"
	"net"
	"time"
)

type GRPCServer struct {
	auth.AuthServiceServer
	pasetoUtil  *paseto.Util
	tokenToAuth map[string]*auth.AuthData
}

const TCP = "tcp"

func NewServer(config *config.Config) error {
	// TCP 연결 해준다.
	listen, err := net.Listen(TCP, config.GRPC.URL)
	if err != nil {
		return err
	}

	// 내가 사용할 서버를 등록하는 과정
	server := grpc.NewServer([]grpc.ServerOption{}...)
	auth.RegisterAuthServiceServer(server, &GRPCServer{
		pasetoUtil:  paseto.CreateInstance(config),
		tokenToAuth: make(map[string]*auth.AuthData),
	})

	// Register registers the server reflection service on the given gRPC server.
	// Both the v1 and v1alpha versions are registered.
	reflection.Register(server)

	go func() {
		log.Println("Hello GRPC!")
		if err = server.Serve(listen); err != nil {
			panic(err)
		}
	}()
	return nil
}

func (server *GRPCServer) CreateAuth(_ context.Context, request *auth.CreateTokenRequest) (*auth.CreateTokenResponse, error) {
	authInfo := request.Auth
	token := authInfo.Token.Value
	server.tokenToAuth[token] = authInfo

	return &auth.CreateTokenResponse{Auth: authInfo}, nil
}

func (server *GRPCServer) VerifyAuth(_ context.Context, request *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error) {
	token := request.Token.Value
	response := &auth.VerifyTokenResponse{Verify: &auth.Verify{Auth: nil}}

	if authInfo, ok := server.tokenToAuth[token]; !ok {
		response.Verify.Status = auth.ResponseType_FAILED
		return response, errors.New("auth map이 없습니다.")
	} else if err := server.pasetoUtil.VerifyToken(token); err != nil {
		return nil, errors.New("토큰 검증에 실패했습니다!")
	} else if authInfo.ExpireDate < time.Now().Unix() {
		delete(server.tokenToAuth, token)
		response.Verify.Status = auth.ResponseType_EXPIRED_DATE
		return response, errors.New("토큰 유효 시간 종료")
	} else {
		response.Verify.Status = auth.ResponseType_SUCCESS
		return response, nil
	}
}
