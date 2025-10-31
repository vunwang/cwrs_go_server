package cwrs_utils

import (
	"cwrs_go_server/src/cwrs_common/cwrs_base"
	"github.com/gin-gonic/gin"
)

// GetLoginUserInfo 获取登录用户信息
func GetLoginUserInfo(c *gin.Context) (loginUser cwrs_base.UserAuth) {
	userId, _ := c.Get("userId")
	deptId, _ := c.Get("deptId")
	roleId, _ := c.Get("roleId")
	roleCode, _ := c.Get("roleCode")
	loginUser.UserId, _ = userId.(string)
	loginUser.DeptId, _ = deptId.(string)
	loginUser.RoleId, _ = roleId.(string)
	loginUser.RoleCode, _ = roleCode.(string)
	return
}
