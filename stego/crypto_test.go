package stego

import (
	"bytes"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	original := []byte("secret steganography payload 12345")
	password := "correct-horse-battery-staple"

	encrypted, err := Encrypt(original, password)
	if err != nil {
		t.Fatalf("encrypt failed: %v", err)
	}

	decrypted, err := Decrypt(encrypted, password)
	if err != nil {
		t.Fatalf("decrypt failed: %v", err)
	}

	if !bytes.Equal(decrypted, original) {
		t.Errorf("decrypted %s does not match original %s", string(decrypted), string(original))
	}
}

func TestDecryptWrongPassword(t *testing.T) {
	original := []byte("secret payload")
	encrypted, _ := Encrypt(original, "pass1")

	_, err := Decrypt(encrypted, "wrong-pass")
	if err == nil {
		t.Errorf("expected error when decrypting with wrong password")
	}
}
