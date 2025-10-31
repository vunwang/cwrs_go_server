package cwrs_res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResSuccess struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ResSuccessData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type ResSuccessDataList struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}

type ResError struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	ErrMsg string `json:"errMsg"`
}

type ResErrorData struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
}

func Success(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, ResSuccess{http.StatusOK, msg})
}

func SuccessData(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, ResSuccessData{http.StatusOK, msg, data})
}
func SuccessDataList(c *gin.Context, msg string, data interface{}, total int64) {
	c.JSON(http.StatusOK, ResSuccessDataList{http.StatusOK, msg, data, total})
}

func Error(c *gin.Context, err error, msg string) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	c.JSON(http.StatusInternalServerError, ResError{http.StatusInternalServerError, msg, errMsg})
}

func Parameter(c *gin.Context, err error, msg string) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	c.JSON(http.StatusPreconditionFailed, ResError{http.StatusPreconditionFailed, msg, errMsg})
}

func Forbidden(c *gin.Context, err error, msg string) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	c.JSON(http.StatusForbidden, ResError{http.StatusForbidden, msg, errMsg})
}

func Waring(c *gin.Context, err error, msg string) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	c.JSON(http.StatusUnprocessableEntity, ResError{http.StatusUnprocessableEntity, msg, errMsg})
}

func WaringData(c *gin.Context, err error, msg string, data interface{}) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	c.JSON(http.StatusUnprocessableEntity, ResErrorData{http.StatusUnprocessableEntity, msg, errMsg, data})
}

func InvalidToken(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": msg})
}
