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
			if hasher.Equals(pass, "fake string") {
				return false
			}
			if hasher.Equals("fake string", pass) {
				return false
			}
			hashedPass := hasher.Hash(pass)
			if !hasher.Equals(pass, hashedPass) {
				return false
			}
			return true
		}

		if err := quick.Check(assertion, &quick.Config{MaxCount: 150}); err != nil {
			t.Error("failed checks", err)
		}
	})
}
