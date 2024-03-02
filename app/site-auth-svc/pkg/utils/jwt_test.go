package utils

import (
	"testing"
	"time"
)

func TestJwt(t *testing.T) {
	userId := int64(20)
	account := "18037283986"
	jw := JwtWrapper{
		JwtIssuer:       "read_user_info",       // 颁发者
		JwtGrantScope:   "Auth_Server",          // 授予范围
		JwtAudience:     []string{"web", "app"}, // 授予对象
		ExpirationHours: 24 * 30,                // 有效时间
	}

	err := InitLoadKey()
	if err != nil {
		t.Errorf("InitLoadKey error: %v", err)
		return
	}

	token, err := jw.GenerateTokenUsingRS256(userId, account)
	if err != nil {
		t.Errorf("GenerateTokenUsingRS256 error: %v", err)
		return
	}
	t.Logf("token: %v", token)
	time.Sleep(time.Second)
	claims, err := jw.ParseTokenRs256(token)
	if err != nil {
		t.Errorf("ParseTokenRs256 error: %v", err)
		return
	}
	t.Logf("claims: %v", claims)
	if claims.UserId != userId || claims.Account != account {
		t.Errorf("valid token failed")
	}
}
