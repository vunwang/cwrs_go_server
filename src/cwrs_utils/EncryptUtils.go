package cwrs_utils

import (
	"encoding/base64"
	"fmt"
	"github.com/forgoer/openssl"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 密码加密
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// CheckPasswordHash 密码校验
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var key = []byte("0123456789abcdef") //这个key不能变, 否则加密解密数据就不对了
// AesECBEncrypt 数据加密(原文)
func AesECBEncrypt(origin string) string {
	src := []byte(origin)

	dst, _ := openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	fmt.Println("加密:", base64.StdEncoding.EncodeToString(dst))
	return base64.StdEncoding.EncodeToString(dst)
}

// AesECBDecrypt 数据解密(密文)
func AesECBDecrypt(ciphertext string) (string, error) {
	decodeString, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	dst, _ := openssl.AesECBDecrypt(decodeString, key, openssl.PKCS7_PADDING)

	fmt.Println("解密:", string(dst))
	return string(dst), err
}
