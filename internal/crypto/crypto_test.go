package crypto

import (
	"testing"
)

func TestCheckEncryptPwdMatch(t *testing.T) {
	hashstring1, _ := EncryptPwd("thisisamockpassworddontlaugh")

	canDecrypt1 := CheckEncryptPwdMatch("thisisamockpassworddontlaugh", *hashstring1)

	if !canDecrypt1 {
		t.Errorf("Result was incorrect")
	}

	hashstring2, _ := EncryptPwd("ChUJTgvzp8QTuLzeTfY62EnR")
	canDecrypt2 := CheckEncryptPwdMatch("ChUJTgvzp8QTuLzeTfY62Enr", *hashstring2)

	if canDecrypt2 {
		t.Errorf("Result was incorrect")
	}

	hashstring3, _ := EncryptPwd("thisisamockpassworddontlaugh")

	if hashstring1 == hashstring3 {
		t.Errorf("Result was incorrect")
	}
}
