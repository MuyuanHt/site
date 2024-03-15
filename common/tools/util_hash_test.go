package tools

import "testing"

func TestHashPassword(t *testing.T) {
	pwd := "abc123@.456"
	encodePwd, err := HashPassword(pwd)
	if err != nil {
		t.Errorf("failed to hash password: %v", err)
		return
	}
	ok := CheckHashPassword(encodePwd, pwd)
	if !ok {
		t.Errorf("failed to check hash password")
	}
}
