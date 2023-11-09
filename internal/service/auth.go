package service

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) Register(login, password string) (*Tokens, error) {

}

func (a *Auth) Login(login, password string) (*Tokens, error) {

}
