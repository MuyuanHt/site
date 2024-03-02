package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"math/rand"
	"site/conf"
	"site/protocol/shared"
	"strconv"
	"strings"
	"time"
)

// JwtWrapper 封装 jwt 对象 用于方法调用
type JwtWrapper struct {
	JwtIssuer       string           // 颁发者
	JwtGrantScope   string           // 授予范围
	JwtAudience     jwt.ClaimStrings // 授予对象
	ExpirationHours int              // 有效时间
}

// AuthJwtClaims 封装 jwt 信息 存储实际信息
type AuthJwtClaims struct {
	UserId     int64
	Account    string
	GrantScope string
	jwt.RegisteredClaims
}

var (
	priKey  []byte                                                                     // priKey 私钥
	pubKey  []byte                                                                     // pubKey 公钥
	letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") // letters 用于生成随机字符串
)

// InitLoadKey 初始化加载密钥
func InitLoadKey() error {
	pri, err := conf.LoadPrivateKey()
	if err != nil {
		return err
	}
	pub, err := conf.LoadPublicKey()
	if err != nil {
		return err
	}
	priKey = pri
	pubKey = pub
	return nil
}

// InitJwt 初始化 jwt
func InitJwt() (*JwtWrapper, error) {
	err := InitLoadKey()
	if err != nil {
		return nil, err
	}
	issuer, err := conf.GetConfigParam("jwtIssuer")
	if err != nil {
		return nil, err
	}
	grantScope, err := conf.GetConfigParam("jwtGrantScope")
	if err != nil {
		return nil, err
	}
	audience, err := conf.GetConfigParam("jwtAudience")
	if err != nil {
		return nil, err
	}
	hours, err := conf.GetConfigParam("jwtExpirationHours")
	if err != nil {
		return nil, err
	}
	expirationHours, err := strconv.Atoi(hours)
	if err != nil {
		return nil, err
	}

	jwtWrapper := &JwtWrapper{
		JwtIssuer:       issuer,
		JwtGrantScope:   grantScope,
		JwtAudience:     strings.Split(audience, "#"),
		ExpirationHours: expirationHours,
	}
	return jwtWrapper, nil
}

// randStr 随机生成字符串
func randStr(strLen int) string {
	randBytes := make([]rune, strLen)
	for i := range randBytes {
		randBytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(randBytes)
}

// parsePriKeyBytes 解析私钥
func parsePriKeyBytes(buf []byte) (*rsa.PrivateKey, error) {
	// p 是 PEM 格式用于存储密钥数据
	block, _ := pem.Decode(buf)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.ParseKeyError))
	}

	// 解析 RSA 私钥
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.ParseKeyError))
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.ParseKeyError))
	}
	return rsaKey, nil
}

// parsePubKeyBytes 解析公钥
func parsePubKeyBytes(buf []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(buf)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.ParseKeyError))
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.ParseKeyError))
	}

	rsaKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.ParseKeyError))
	}
	return rsaKey, nil
}

// ParseTokenRs256 解析 token
func (jw *JwtWrapper) ParseTokenRs256(signedToken string) (*AuthJwtClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&AuthJwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			pub, err := parsePubKeyBytes(pubKey)
			if err != nil {
				log.Printf("Error parsing token error: %v", err)
				return nil, err
			}
			return pub, nil
		},
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.TokenInvalid))
	}
	claims, ok := token.Claims.(*AuthJwtClaims)
	if !ok {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.TokenParsingError))
	}
	return claims, nil
}

// GenerateTokenUsingRS256 生成 token
func (jw *JwtWrapper) GenerateTokenUsingRS256(id int64, account string) (string, error) {
	issueTime := time.Now() // 颁发时间
	claim := AuthJwtClaims{
		UserId:     id,
		Account:    account,
		GrantScope: jw.JwtGrantScope,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    jw.JwtIssuer,                                                                     // 签发者
			Subject:   account,                                                                          // 签发对象
			Audience:  jw.JwtAudience,                                                                   // 签发受众
			ExpiresAt: jwt.NewNumericDate(issueTime.Add(time.Hour * time.Duration(jw.ExpirationHours))), // 过期时间
			NotBefore: jwt.NewNumericDate(issueTime.Add(time.Second)),                                   // 最早使用时间
			IssuedAt:  jwt.NewNumericDate(issueTime),                                                    // 签发时间
			ID:        randStr(32),                                                                      // jwt ID, 类似于盐值
		},
	}
	rsaPriKey, err := parsePriKeyBytes(priKey)
	if err != nil {
		return "", err
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claim).SignedString(rsaPriKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
