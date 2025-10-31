package controller

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/server/sys_user/pojo"
	"cwrs_go_server/src/server/sys_user/service"
	"github.com/gin-gonic/gin"
)

var loginServiceImpl = service.LoginService{}

// @Tags 管理端登录【平台】
// @Summary 获取用户身份
// @Description 获取用户身份
// @Accept json
// @Produce json
// @Param req query pojo.UserIdentityReq true "用户信息"
// @Success 200 {object} cwrs_res.ResSuccessData{data=[]pojo.UserIdentityResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysLogin/getUserIdentity [get]
func GetUserIdentity(c *gin.Context) {
	var req pojo.UserIdentityReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	loginServiceImpl.GetUserIdentity(c, &req)
}

// @Tags 管理端登录【平台】
// @Summary 用户登录
// @Description 用户登录
// @Accept json
// @Produce json
// @Param req query pojo.LoginReq true "用户信息"
// @Success 200 {object} cwrs_res.ResSuccessData{data=string} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysLogin/login [get]
func LoginUser(c *gin.Context) {
	var req pojo.LoginReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	loginServiceImpl.LoginUser(c, &req)
}

// @Tags 管理端登录【平台】
// @Summary 获取用户信息
// @Description 获取用户信息
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Success 200 {object} cwrs_res.ResSuccessData{data=pojo.UserInfoResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysLogin/getUserInfo [get]
func GetUserInfo(c *gin.Context) {
	loginServiceImpl.GetUserInfo(c)
}

// @Tags 管理端登录【平台】
// @Summary 退出登录
// @Description 退出登录
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Success 200 {object} cwrs_res.ResSuccessData{data=pojo.UserInfoResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysLogin/logout [get]
func Logout(c *gin.Context) {
	loginServiceImpl.Logout(c)
}
