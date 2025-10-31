package cwrs_utils

import (
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"fmt"
	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"math/rand"
	"strings"
)

var logUtils = cwrs_zap_logger.ZapLogger

// LowerFirst 将字符串首字母转为小写
func LowerFirst(s string) string {
	if s == "" {
		return ""
	}
	if len(s) == 1 {
		return strings.ToLower(s)
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// 复制结构体
func CopyStruct(src, dst interface{}, tagName string) {
	// 先把 src 转成 map
	srcMap := structs.Map(src)

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  dst,
		TagName: tagName, // 比如 "json" 或 "mapstructure"
	})

	err := decoder.Decode(srcMap)
	if err != nil {
		fmt.Println("结构体复制失败:", err)
		logUtils.Error("结构体复制失败", zap.Error(err))
	}
}

// 获取32位uuid
func CreateUuid() string {
	// 生成一个新的UUID v4
	newUUID, err := uuid.NewRandom()
	if err != nil {
		fmt.Println("获取UUID失败:", err)
		return ""
	}

	// 转换为字符串，并移除连字符
	uuidStr := strings.ReplaceAll(newUUID.String(), "-", "")

	fmt.Println("Generated UUID (32 characters):", uuidStr)

	return uuidStr
}

var CHARS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

/*
RandAllString  生成随机字符串([a~zA~Z0~9])
lenNum 长度
*/
func RandAllString(lenNum int) string {
	str := strings.Builder{}
	length := len(CHARS)
	for i := 0; i < lenNum; i++ {
		l := CHARS[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

/*
RandNumString  生成随机数字字符串([0~9])

	lenNum 长度
*/
func RandNumString(lenNum int) string {
	str := strings.Builder{}
	length := 10
	for i := 0; i < lenNum; i++ {
		str.WriteString(CHARS[52+rand.Intn(length)])
	}
	return str.String()
}

/*
RandString  生成随机字符串(a~zA~Z])

	lenNum 长度
*/
func RandString(lenNum int) string {
	str := strings.Builder{}
	length := 52
	for i := 0; i < lenNum; i++ {
		str.WriteString(CHARS[rand.Intn(length)])
	}
	return str.String()
}
