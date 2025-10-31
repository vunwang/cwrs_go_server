package service

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/cwrs_core/cwrs_jwt"
	"cwrs_go_server/src/cwrs_core/cwrs_redis"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_user/dao"
	"cwrs_go_server/src/server/sys_user/pojo"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var loginDaoImpl dao.LoginDao

type LoginService struct{}

// 登录
func (*LoginService) LoginUser(c *gin.Context, req *pojo.LoginReq) {
	//查询用户信息
	userResp, err := loginDaoImpl.LoginUser(req)
	if err != nil {
		cwrs_res.Waring(c, err, "用户不存在，请检查用户名或手机号是否正确！")
		return
	}
	//密码验证
	ok := cwrs_utils.CheckPasswordHash(req.Password, userResp.Password)
	userResp.Password = "" //密码清空
	if !ok {
		cwrs_res.Waring(c, nil, "请检查密码是否正确！")
		return
	}
	//发放token
	token, tokenErr := cwrs_jwt.GenJwtToken(
		userResp.UserId,
		req.RoleId,
		req.DeptId,
		req.RoleCode,
	)
	if tokenErr != nil {
		cwrs_res.Waring(c, tokenErr, "token发放失败，请检查.")
		return
	}
	tokenKey := fmt.Sprintf("%s%s:%s", cwrs_redis.KEY_SYS_USER_TOKEN, req.DeptId, userResp.UserId)
	err = cwrs_redis.GlobalRedis.Set(c.Request.Context(), tokenKey, token, cwrs_redis.KEY_SYS_USER_TOKEN_TIME).Err()
	if err != nil {
		cwrs_zap_logger.Error("设置token过期时间失败", zap.Error(err))
		cwrs_res.Waring(c, err, "登录失败，请重试")
		return
	}
	cwrs_res.SuccessData(c, "操作成功", token)
}

// 获取用户身份信息
func (*LoginService) GetUserIdentity(c *gin.Context, req *pojo.UserIdentityReq) {
	userResp, err := loginDaoImpl.GetUserIdentity(req)

	if err != nil {
		cwrs_res.Waring(c, err, "获取用户身份失败")
		return
	}
	cwrs_res.SuccessData(c, "操作成功", userResp)
}

// 获取用户信息
func (*LoginService) GetUserInfo(c *gin.Context) {
	//获取当前登录用户信息
	user := cwrs_utils.GetLoginUserInfo(c)
	userId := user.UserId

	//根据userId查询用户详细信息
	userResp, err := loginDaoImpl.GetUserInfo(userId)
	if err != nil {
		cwrs_res.Waring(c, err, "获取用户信息失败")
		return
	}
	permissions, err := loginDaoImpl.GetPermissions(&user)
	if err != nil {
		cwrs_res.Waring(c, err, "获取按钮标识失败")
		return
	}
	userResp.Permissions = permissions
	userResp.DeptId = user.DeptId
	userResp.RoleId = user.RoleId
	userResp.RoleCode = user.RoleCode
	cwrs_res.SuccessData(c, "操作成功", userResp)
}

// 退出登录
func (*LoginService) Logout(c *gin.Context) {
	//获取当前登录用户信息
	user := cwrs_utils.GetLoginUserInfo(c)
	//删除redis中的token
	tokenKey := fmt.Sprintf("%s%s:%s", cwrs_redis.KEY_SYS_USER_TOKEN, user.DeptId, user.UserId)
	err := cwrs_redis.GlobalRedis.Del(c.Request.Context(), tokenKey).Err()
	if err != nil {
		cwrs_zap_logger.Error("删除token失败", zap.Error(err))
		cwrs_res.Waring(c, err, "退出失败")
		return
	}
	//cwrs_res.InvalidToken(c, "退出成功")
	c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "msg": "退出成功"})
}
