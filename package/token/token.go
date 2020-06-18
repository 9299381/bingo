package token

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/9299381/bingo/package/config"
	"github.com/9299381/bingo/package/util"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	ErrNoToken   string = "9020::缺少authToken"
	ErrTokenFmt  string = "9021::Token格式错误"
	ErrTokenExp  string = "9022::Token过期,请重新登录"
	ErrTokenSign string = "9023::签名错误,请重新登录"
)

var token *Token
var once sync.Once

func instance() *Token {
	once.Do(func() {
		token = initToken()
	})
	return token
}

func GetToken(claims *Claims) (string, error) {
	return instance().getToken(claims)
}

func CheckToken(token string) (*Claims, error) {
	return instance().checkToken(token)
}

type Token struct {
	Key string
	Exp int64
}

func initToken() *Token {
	t := &Token{}
	t.Key = config.EnvString("token.key", "EHKHHP54PXKYTS2E")
	t.Exp = int64(config.EnvInt("token.exp", 2592000))
	return t
}

func (t *Token) getToken(claims *Claims) (string, error) {
	iat := time.Now().Unix()
	claims.Iat = iat
	claims.Exp = iat + t.Exp
	jsonClaim, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	payload := base64.StdEncoding.EncodeToString(jsonClaim)
	sign := t.getSign(claims)
	ret := strings.Join([]string{payload, sign}, ".")
	return ret, err
}

func (t *Token) checkToken(token string) (*Claims, error) {
	m := strings.Split(token, ".")
	if len(m) < 1 {
		return nil, errors.New(ErrTokenFmt)
	}
	jsonClaim, decodeErr := base64.StdEncoding.DecodeString(m[0])
	if decodeErr != nil {
		return nil, decodeErr
	}
	claims := &Claims{}
	jsonErr := json.Unmarshal(jsonClaim, claims)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if claims.Exp < time.Now().Unix() {
		return nil, errors.New(ErrTokenExp)
	}

	if m[1] != t.getSign(claims) {
		return nil, errors.New(ErrTokenSign)
	}

	return claims, nil
}

func (t *Token) getSign(claims *Claims) string {
	keyPlain := claims.Id + strconv.Itoa(int(claims.Iat)) + t.Key
	return util.Md5(keyPlain)
}
