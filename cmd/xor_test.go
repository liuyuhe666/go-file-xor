package cmd

import "testing"

func TestXor(t *testing.T) {
	secret := "MHgxMzEyZGQ2"
	text := "Hello, World!"
	result := Xor([]byte(text), []byte(secret))
	result = Xor(result, []byte(secret))
	if text != string(result) {
		t.Fatalf("expect %s, but got %s", text, string(result))
	}
}
