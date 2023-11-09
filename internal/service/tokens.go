package service

type Tokens struct {
	Access  string
	Refresh string
}

type TokensService struct {
}

func NewTokensService() *TokensService {
	return &TokensService{}
}

func (t *TokensService) CreatePair(userId string) (*Tokens, error) {

}

func (t *TokensService) Refresh(refreshToken string) (*Tokens, error) {
	panic("unimplemented")
}

func (t *TokensService) Verify(accessToken string) bool {

}
