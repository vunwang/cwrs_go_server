package cwrs_gorm

import (
	"cwrs_go_server/src/cwrs_common/cwrs_constants"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// WithDataScope 自动注入数据权限条件
func WithDataScope(c *gin.Context, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		userId, _ := c.Get("userId")
		deptId, _ := c.Get("deptId")
		roleCode, _ := c.Get("roleCode")
		dataPurview, _ := c.Get("dataPurview")
		if roleCode == "sys_admin" {
			return db // 超管不限制
		}
		// 自动注入数据权限条件
		switch dataPurview {
		//全部
		case cwrs_constants.DataPurviewAll:
			return db // 无限制
		//仅自己
		case cwrs_constants.DataPurviewSelf:
			return db.Where(tableName+".created_user_id = ?", userId)
		//本组织
		case cwrs_constants.DataPurviewOwnDept:
			return db.Where(tableName+".dept_id = ?", deptId)
		//本组织及以下 || 自定义
		case cwrs_constants.DataPurviewDeptAndChild, cwrs_constants.DataPurviewCustom:
			if allowed, ok := c.Get("allowedDeptIds"); ok {
				deptIds := allowed.([]string)
				if len(deptIds) == 0 {
					return db.Where("1 = 0") // 无权限
				}
				return db.Where(tableName+".dept_id IN (?)", deptIds) // 自定义权限
			}
			return db.Where("1 = 0") // 无权限

		default:
			return db.Where("1 = 0") // 默认无权限
		}
	}
}
