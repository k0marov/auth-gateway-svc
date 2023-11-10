package service

import (
	"auth-gateway-svc/internal/config"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"log"
	"os"
	"time"
)

type Tokens struct {
	Access string
}

type TokensService struct {
	cfg        config.JWT
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewTokensService(cfg config.JWT) *TokensService {
	priv, err := os.ReadFile(cfg.PrivateKeyFilePath)
	if err != nil {
		log.Panicf("while reading private key file: %v", err)
	}
	block, _ := pem.Decode(priv)
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Panicf("unable to parse pkcs1 rsa private key from %s: %v", cfg.PrivateKeyFilePath, err)
	}
	rsaPublicKey := &privKey.PublicKey
	return &TokensService{cfg, privKey, rsaPublicKey}
}

func (t *TokensService) Create(userLogin string) *Tokens {
	access, err := jwt.NewBuilder().
		Subject(userLogin).
		Expiration(time.Now().Add(time.Duration(t.cfg.AccessExpireInSec) * time.Second)).
		Build()
	if err != nil {
		log.Panicf("while building access token: %v", err)
	}
	accessBytes, err := jwt.Sign(access, jwt.WithKey(jwa.RS256, t.privateKey))
	if err != nil {
		log.Panicf("while signing access token: %v", err)
	}
	return &Tokens{Access: string(accessBytes)}
}

func (t *TokensService) Verify(accessToken string) bool {
	_, err := jwt.Parse([]byte(accessToken), jwt.WithKey(jwa.RS256, t.publicKey), jwt.WithValidate(true))
	return err == nil
}

func (t *TokensService) VerifyAdmin(gotAdminSecret string) bool {
	return gotAdminSecret == t.cfg.AdminSecret
}
