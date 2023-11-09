package bcrypt_hasher_test

import (
	"auth-gateway-svc/internal/service/bcrypt_hasher"
	"testing"
	"testing/quick"
)

func TestBcryptHasher(t *testing.T) {
	t.Run("property based test", func(t *testing.T) {
		hasher := bcrypt_hasher.NewBcryptHasher(5)
		assertion := func(pass string) bool {
			if hasher.Compare(pass, "fake string") {
				return false
			}
			if hasher.Compare("fake string", pass) {
				return false
			}
			hashedPass, err := hasher.Hash(pass)
			if err != nil {
				return false
			}
			if !hasher.Compare(pass, hashedPass) {
				return false
			}
			return true
		}

		if err := quick.Check(assertion, &quick.Config{MaxCount: 150}); err != nil {
			t.Error("failed checks", err)
		}
	})
}
