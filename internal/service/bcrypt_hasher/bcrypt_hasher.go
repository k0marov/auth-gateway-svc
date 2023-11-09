package bcrypt_hasher

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct {
	hashCost int
}

func NewBcryptHasher(hashCost int) *BcryptHasher {
	return &BcryptHasher{hashCost: hashCost}
}

func (b BcryptHasher) Hash(pass string) (string, error) {
	hashedPassBytes, err := bcrypt.GenerateFromPassword([]byte(pass), b.hashCost)
	if err != nil {
		log.Panicf("hashing failed: %v", err)
	}
	return string(hashedPassBytes), nil
}

func (b BcryptHasher) Compare(pass, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	return err == nil
}
