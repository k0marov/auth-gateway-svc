package service

import (
	"auth-gateway-svc/internal/config"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"log"
	"time"
)

type Tokens struct {
	Access string
}

type TokensService struct {
	cfg config.JWT
}

func NewTokensService(cfg config.JWT) *TokensService {
	return &TokensService{cfg}
}

func (t *TokensService) CreatePair(userId string) (*Tokens, error) {
	access, err := jwt.NewBuilder().
		Subject(userId).
		Expiration(time.Now().Add(time.Duration(t.cfg.AccessExpireInSec) * time.Second)).
		Build()
	if err != nil {
		log.Panicf("while building access token: %v", err)
	}
	accessBytes, err := jwt.Sign(access, jwt.WithKey(jwa.RS256, t.cfg.PrivateKey))
	if err != nil {
		log.Panicf("while signing access token: %v", err)
	}
	return &Tokens{Access: string(accessBytes)}, nil
}

func (t *TokensService) Refresh(refreshToken string) (*Tokens, error) {
	panic("unimplemented")
}

func (t *TokensService) Verify(accessToken string) bool {
	_, err := jwt.Parse([]byte(accessToken), jwt.WithKey(jwa.RS256, t.cfg.PrivateKey), jwt.WithValidate(true))
	return err == nil
}

func (t *TokensService) VerifyAdmin(gotAdminSecret string) bool {
	return gotAdminSecret == t.cfg.AdminSecret
}
