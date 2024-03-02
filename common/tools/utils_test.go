package tools

import (
	consts "site/protocol/shared"
	"testing"
)

func TestGetCurrentDir(t *testing.T) {
	path := GetCurrentPath()
	t.Logf("resource path: %v\n", path)
	want := "tools"
	got := GetCurrentDir(path)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetAccountType(t *testing.T) {
	phone1 := "18293840245"
	phone2 := "12345"
	email1 := "123@example.com"
	email2 := "123@example@123"
	if GetAccountType(phone1) != consts.AccountTypePhone {
		t.Errorf("phone1: %v is a phone but got is false", phone1)
	}
	if GetAccountType(phone2) == consts.AccountTypePhone {
		t.Errorf("phone2: %v is not a phone but got is true", phone2)
	}
	if GetAccountType(email1) != consts.AccountTypeEmail {
		t.Errorf("email1: %v is an email but got is false", email1)
	}
	if GetAccountType(email2) == consts.AccountTypeEmail {
		t.Errorf("email2: %v is not an email but got is true", email2)
	}
}
