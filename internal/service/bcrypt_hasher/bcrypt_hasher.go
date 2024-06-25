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

func (b BcryptHasher) Hash(pass string) string {
	hashedPassBytes, err := bcrypt.GenerateFromPassword([]byte(pass), b.hashCost)
	if err != nil {
		log.Panicf("hashing failed: %v", err)
	}
	return string(hashedPassBytes)
}

func (b BcryptHasher) Equals(pass, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	return err == nil
}
