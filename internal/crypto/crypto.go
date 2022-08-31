package crypto

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	mathrand "math/rand"

	"golang.org/x/crypto/pbkdf2"
)

const (
	saltMinLen = 8
	saltMaxLen = 32
	iter       = 20000
	keyLen     = 32
)

// Generate key
func EncryptPwd(pwd string) (*string, error) {

	// Generating random salt
	salt, err := randSalt()
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	// Generating Hash
	en := encryptPwdWithSalt([]byte(pwd), salt)
	en = append(en, salt...)

	// Generating final string
	encrypt := base64.StdEncoding.EncodeToString(en)

	return &encrypt, nil
}

// Generating salt
func randSalt() ([]byte, error) {

	// Generating a random salt of 8-32 bit
	salt := make([]byte, mathrand.Intn(saltMaxLen-saltMinLen)+saltMinLen)
	_, err := rand.Read(salt)

	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Add salt to password then hash through n iterations of pbkdf2
func encryptPwdWithSalt(pwd, salt []byte) []byte {
	pwd = append(pwd, salt...)
	pwdEn := pbkdf2.Key(pwd, salt, iter, keyLen, sha512.New)
	return pwdEn
}

// Check if encryption match password
func CheckEncryptPwdMatch(pwd, encrypt string) bool {

	if len(encrypt) == 0 {
		return false
	}

	enDecode, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return false
	}

	// Get salt
	salt := enDecode[keyLen:]

	// Comparison
	enBase64 := base64.StdEncoding.EncodeToString(enDecode[0:keyLen])
	pwdEnBase64 := base64.StdEncoding.EncodeToString(encryptPwdWithSalt([]byte(pwd), salt))

	return enBase64 == pwdEnBase64
}
