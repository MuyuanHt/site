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

func TestIsZeroValue(t *testing.T) {
	var a int
	if !IsZeroValue(a) {
		t.Errorf("a = %v, but IsZeroValue return false.", a)
	}
	var b int = 2
	if IsZeroValue(b) {
		t.Errorf("b = %v, but IsZeroValue return true.", b)
	}
	var str1 string = ""
	if !IsZeroValue(str1) {
		t.Errorf("str1 = %v, but IsZeroValue return false.", str1)
	}
	var str2 string = "123"
	if IsZeroValue(str2) {
		t.Errorf("str2 = %v, but IsZeroValue return true.", str2)
	}
}

func TestCheckPwdRegexp(t *testing.T) {
	pwd1 := "a12"
	pwd2 := "1234567"
	pwd3 := "abcDef"
	pwd4 := "abc123"
	if CheckPwdRegexp(pwd1) {
		t.Errorf("pwd1 = %v, invalid but return true", pwd1)
	}
	if CheckPwdRegexp(pwd2) {
		t.Errorf("pwd2 = %v, invalid but return true", pwd2)
	}
	if CheckPwdRegexp(pwd3) {
		t.Errorf("pwd3 = %v, invalid but return ture", pwd3)
	}
	if !CheckPwdRegexp(pwd4) {
		t.Errorf("pwd4 = %v, valid but return false", pwd4)
	}
}
