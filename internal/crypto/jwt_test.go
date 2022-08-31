package crypto

import (
	"testing"
)

func TestJWT(t *testing.T) {

	// First case: Can verify
	testUser := "arthur"

	tokenString1, _, _ := GenerateJWT(testUser)
	canDecrypt1, any1, _ := ParseJWT(tokenString1)

	if !canDecrypt1 || any1["username"] != "arthur" {
		t.Errorf("expected to be able to decrypt, but %v", any1)
	}
}
