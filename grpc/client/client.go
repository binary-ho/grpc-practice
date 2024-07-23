package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-practice/config"
	paseto "grpc-practice/grpc/pasto"
	"grpc-practice/grpc/proto"
	"time"
)

type GRPCClient struct {
	client     *grpc.ClientConn
	authClient auth.AuthServiceClient
	pasetoUtil *paseto.Util
}

func NewGRPCClient(config *config.Config) (*GRPCClient, error) {
	client := new(GRPCClient)
	dial, err := grpc.Dial(config.GRPC.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client.client = dial
	client.authClient = auth.NewAuthServiceClient(client.client)
	client.pasetoUtil = paseto.CreateInstance(config)
	return client, nil
}

func (grpc *GRPCClient) CreateAuth(name string) (*auth.AuthData, error) {
	now := time.Now()
	expiredTime := now.Add(30 * time.Minute)

	authData := &auth.AuthData{
		Name:       name,
		CreateDate: now.Unix(),
		ExpireDate: expiredTime.Unix(),
	}

	token, err := grpc.pasetoUtil.CreateToken(authData)
	if err != nil {
		return nil, err
	}
	authData.Token.Value = token

	// MEMO : Server에 있는 CreateAuth를 호출한다.(RPC)
	// 이 프로젝트에서는 Server Package에 선언된다.
	response, err := grpc.authClient.CreateAuth(context.Background(), &auth.CreateTokenRequest{Auth: authData})

	if err != nil {
		return nil, err
	}
	return response.GetAuth(), nil
}

func (grpc *GRPCClient) VerifyAuth(token string) (*auth.Verify, error) {
	response, err := grpc.authClient.VerifyAuth(context.Background(), &auth.VerifyTokenRequest{Token: &auth.Token{Value: token}})
	if err != nil {
		return nil, err
	}
	return response.Verify, nil
}
