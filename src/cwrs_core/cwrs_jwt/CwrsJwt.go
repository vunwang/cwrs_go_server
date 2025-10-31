package cwrs_jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	UserId   string `json:"userId"`
	RoleId   string `json:"roleId"`
	DeptId   string `json:"deptId"`
	RoleCode string `json:"roleCode"`
	jwt.StandardClaims
}

// TokenExpireDuration Token过期时间
const TokenExpireDuration = time.Hour * 24 * 7

// 密钥/盐
var mySecret = []byte("CwrsSecret")

// GenJwtToken 生成JWT
func GenJwtToken(userId, roleId, deptId, roleCode string) (string, error) {
	// 使用自定义token部分内容
	c := MyClaims{
		userId,   // 当前用户id
		roleId,   // 当前用户角色id
		deptId,   // 当前用户组织id
		roleCode, // 当前用户组织id
		//创建token
		jwt.StandardClaims{
			//ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 设置过期时间 注释掉怎生成不含过期时间的token
			Issuer: "system", // 设置签发人
		},
	}
	// 指定加密方式,创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret/盐 签名并获得完整的编码后的字符串token
	signedString, err := token.SignedString(mySecret)
	return signedString, err
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var claims = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
