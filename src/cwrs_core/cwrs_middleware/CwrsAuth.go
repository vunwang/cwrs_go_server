package cwrs_middleware

import (
	"cwrs_go_server/src/cwrs_common/cwrs_constants"
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/cwrs_core/cwrs_gorm"
	"cwrs_go_server/src/cwrs_core/cwrs_jwt"
	"cwrs_go_server/src/cwrs_core/cwrs_redis"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		menuId := c.GetHeader("Menu-Id")

		claims, err := cwrs_jwt.ParseToken(token)
		if err != nil {
			cwrs_res.InvalidToken(c, "无效的令牌")
			c.Abort()
			return
		}

		// 设置基础用户信息
		c.Set("menuId", menuId)
		c.Set("deptId", claims.DeptId)
		c.Set("roleId", claims.RoleId)
		c.Set("userId", claims.UserId)
		c.Set("roleCode", claims.RoleCode)

		// 超管不校验数据权限
		if claims.RoleCode != "sys_admin" && menuId != "" && menuId != "undefined" {
			// 查询数据权限
			var dataPurview string
			err = cwrs_gorm.GormDb.Table("sys_role_menu").Select("data_purview").
				Where("dept_id = ? AND role_id = ? AND menu_id = ?", claims.DeptId, claims.RoleId, menuId).
				Scan(&dataPurview).Error

			// 数据权限不存在或为空
			if err != nil || len(dataPurview) == 0 {
				c.Abort()
				return
			}
			c.Set("dataPurview", dataPurview)
			// 如果是本组织及以下权限，查对应组织
			if dataPurview == cwrs_constants.DataPurviewDeptAndChild {
				// 获取本组织层级
				var deptLevel string
				err = cwrs_gorm.GormDb.Table("sys_dept").Select("dept_level").
					Where("dept_id = ?", claims.DeptId).
					Scan(&deptLevel).Error
				if err != nil {
					cwrs_res.Waring(c, err, "查询本组织层级失败")
					c.Abort()
					return
				}
				// 获取本组织及子组织ID列表
				var deptIds []string
				err = cwrs_gorm.GormDb.Table("sys_dept").Select("dept_id").
					Where("dept_level LIKE ?", deptLevel+"%").
					Pluck("dept_id", &deptIds).Error
				if err != nil {
					cwrs_res.Waring(c, err, "查询本组织及子组织ID列表失败")
					c.Abort()
					return
				}
				c.Set("allowedDeptIds", deptIds)
			} else if dataPurview == cwrs_constants.DataPurviewCustom { // 如果是自定义权限，查授权组织
				var deptIds []string
				err = cwrs_gorm.GormDb.Table("sys_data_scope").Select("dept_id").
					Where("scope_type = 1 AND role_id = ? AND menu_id = ?", claims.RoleId, menuId).
					Pluck("dept_id", &deptIds).Error
				if err != nil {
					cwrs_res.Waring(c, err, "查询自定义权限失败")
					c.Abort()
					return
				}
				c.Set("allowedDeptIds", deptIds)
			}
		}

		//redis 校验token是否过期
		tokenKey := fmt.Sprintf("%s%s:%s", cwrs_redis.KEY_SYS_USER_TOKEN, claims.DeptId, claims.UserId)
		resErr := cwrs_redis.GlobalRedis.Get(c.Request.Context(), tokenKey).Err()
		if resErr != nil {
			cwrs_res.InvalidToken(c, "无效的令牌")
			c.Abort()
			return
		}

		// 重置token过期时间 redis
		err = cwrs_redis.GlobalRedis.Set(c.Request.Context(), tokenKey, token, cwrs_redis.KEY_SYS_USER_TOKEN_TIME).Err()
		if err != nil {
			cwrs_zap_logger.Error("重置token过期时间失败", zap.Error(err))
			return
		}
		c.Next()
	}
}
