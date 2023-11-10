package service

import (
	"auth-gateway-svc/internal/core"
	"fmt"
)

type JWTPairCreator interface {
	Create(userLogin string) *Tokens
}

type UserRepo interface {
	Create(login, hashedPassword string) error
	GetStoredPass(login string) (string, error)
}

type Hasher interface {
	Hash(string) string
	Equals(raw string, hashed string) bool
}

type Auth struct {
	jwt    JWTPairCreator
	repo   UserRepo
	hasher Hasher
}

func NewAuth(jwt JWTPairCreator, repo UserRepo, hasher Hasher) *Auth {
	return &Auth{jwt, repo, hasher}
}

func (a *Auth) Register(login, password string) (*Tokens, error) {
	passHashed := a.hasher.Hash(password)
	err := a.repo.Create(login, passHashed)
	if err != nil {
		return nil, fmt.Errorf("failed to create user in repo: %w", err)
	}
	return a.jwt.Create(login), nil
}

func (a *Auth) Login(login, password string) (*Tokens, error) {
	hashedPass, err := a.repo.GetStoredPass(login)
	if err != nil {
		return nil, fmt.Errorf("while getting stored pass from repo: %w", err)
	}
	if !a.hasher.Equals(password, hashedPass) {
		return nil, core.CEInvalidCredentials
	}
	return a.jwt.Create(login), nil
}
