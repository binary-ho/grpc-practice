package paseto

import (
	"crypto/rand"
	"github.com/o1egl/paseto"
	"grpc-practice/config"
	auth "grpc-practice/grpc/proto"
)

type Util struct {
	Paseto *paseto.V2
	Key    []byte
}

func CreateInstance(config *config.Config) *Util {
	return &Util{Paseto: paseto.NewV2(), Key: []byte(config.Paseto.Key)}
}

func (paseto *Util) CreateToken(auth *auth.AuthData) (string, error) {
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	return paseto.Paseto.Encrypt(paseto.Key, auth, randomBytes)
}

func (paseto *Util) VerifyToken(token string) error {
	var auth auth.AuthData
	// 타입이 맞는지 검증하는 것이다.
	return paseto.Paseto.Decrypt(token, paseto.Key, &auth, nil)
}
