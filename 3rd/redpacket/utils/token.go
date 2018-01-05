package utils

import (
	"time"
	"crypto/md5"
	"fmt"
	"encoding/base64"
	"math/rand"
)

type Token struct {
	Token string
	ExpiredAt int64
}
//这里用token做简单校验，为了简单，直接用全局变量存储token
type TokenMapper map[uint]Token

//func (mapper *TokenMapper) checkToken(uid int,token string )  {
//
//}

func (t *Token) Validate(token string) bool {
	if t.ExpiredAt < time.Now().Unix() || token != t.Token{
		return false
	}
	return true
}

func (t *Token) Generate()  {
	ut := time.Now().Unix()

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(fmt.Sprintf("%d%d",rand.Int(),ut)))

	cipherStr := md5Ctx.Sum(nil)
	t.Token = base64.StdEncoding.EncodeToString(cipherStr)
	t.ExpiredAt = ut + 3600*24
}