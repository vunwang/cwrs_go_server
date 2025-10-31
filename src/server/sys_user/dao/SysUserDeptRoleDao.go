package dao

import (
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/server/sys_user/pojo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var sysUserDeptRoleLog = cwrs_zap_logger.ZapLogger

const tableSysUserDeptRole = "sys_user_dept_role"
const tableSelSysUserDeptRole = "sys_user_dept_role sudr"

type SysUserDeptRoleDao struct{}

// AddSysUserDeptRole 添加用户身份
func (*SysUserDeptRoleDao) AddSysUserDeptRole(tx *gorm.DB, item *pojo.SysUserDeptRole) error {
	fields := make([]string, 0)
	if item.UserId != "" {
		fields = append(fields, "user_id")
	}
	if item.DeptId != "" {
		fields = append(fields, "dept_id")
	}
	if item.RoleId != "" {
		fields = append(fields, "role_id")
	}
	if item.PostId != "" {
		fields = append(fields, "post_id")
	}
	if item.DataPurview != "" {
		fields = append(fields, "data_purview")
	}
	fields = append(fields, "user_dept_role_id", "created_user_id")
	if err := tx.Table(tableSysUserDeptRole).Select(fields).Create(&item).Error; err != nil {
		sysUserDeptRoleLog.Error("AddSysUserDeptRole Error", zap.Error(err))
		return err
	}
	return nil
}

// EditSysUserDeptRole 修改用户身份
func (*SysUserDeptRoleDao) EditSysUserDeptRole(tx *gorm.DB, item *pojo.SysUserDeptRole) error {
	fields := make([]string, 0)
	if item.UserId != "" {
		fields = append(fields, "user_id")
	}
	if item.DeptId != "" {
		fields = append(fields, "dept_id")
	}
	if item.RoleId != "" {
		fields = append(fields, "role_id")
	}
	if err := tx.Table(tableSysUserDeptRole).Select(fields, "post_id", "data_purview").
		Where("dept_id = ? AND user_id = ?", item.OldDeptId, item.UserId).Updates(&item).Error; err != nil {
		sysUserDeptRoleLog.Error("EditSysUserDeptRole Error", zap.Error(err))
		return err
	}
	return nil
}

// DelSysUserDeptRole 删除用户身份
func (*SysUserDeptRoleDao) DelSysUserDeptRole(tx *gorm.DB, userIds []string) error {
	if err := tx.Table(tableSysUserDeptRole).
		Where("user_id IN (?)", userIds).Delete(&pojo.SysUserDeptRole{}).Error; err != nil {
		sysUserDeptRoleLog.Error("DelSysUserDeptRole Error", zap.Error(err))
		return err
	}
	return nil
}
