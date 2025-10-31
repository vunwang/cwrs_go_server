package dao

import (
	"cwrs_go_server/src/cwrs_core/cwrs_gorm"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_role/pojo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var roleLog = cwrs_zap_logger.ZapLogger

const tableRole = "sys_role"
const tableSysRole = "sys_role sr"

type SysRoleDao struct{}

// AddSysRole 添加角色
func (*SysRoleDao) AddSysRole(role *pojo.SysRole) error {
	fields := make([]string, 0)
	if role.RoleName != "" {
		fields = append(fields, "role_name")
	}
	if role.RoleCode != "" {
		fields = append(fields, "role_code")
	}
	if role.RoleStatus != "" {
		fields = append(fields, "role_status")
	}
	if role.RoleSort != 0 {
		fields = append(fields, "role_sort")
	}
	if role.DeptId != "" {
		fields = append(fields, "dept_id")
	}
	if role.Desc != "" {
		fields = append(fields, "desc")
	}
	fields = append(fields, "role_id", "parent_id", "role_level", "is_builtin", "created_user_id", "created_time")

	if err := cwrs_gorm.GormDb.Table(tableRole).Select(fields).
		Create(&role).Error; err != nil {
		roleLog.Error("AddSysRole Error", zap.Error(err))
		return err
	}
	return nil
}

// EditSysRole 修改角色
func (*SysRoleDao) EditSysRole(role *pojo.SysRole) error {
	fields := make([]string, 0)
	if role.RoleName != "" {
		fields = append(fields, "role_name")
	}
	if role.RoleCode != "" {
		fields = append(fields, "role_code")
	}
	if role.RoleStatus != "" {
		fields = append(fields, "role_status")
	}
	if role.RoleSort != 0 {
		fields = append(fields, "role_sort")
	}
	if role.DeptId != "" {
		fields = append(fields, "dept_id")
	}
	fields = append(fields, "parent_id", "role_level", "is_builtin", "desc", "updated_user_id", "updated_time")
	if err := cwrs_gorm.GormDb.Table(tableRole).Select(fields).
		Where("role_id = ?", role.RoleId).Updates(&role).Error; err != nil {
		roleLog.Error("EditSysRole Error", zap.Error(err))
		return err
	}
	return nil
}

// DelSysRole 批量删除角色
func (*SysRoleDao) DelSysRole(roleId string) error {
	if err := cwrs_gorm.GormDb.Table(tableRole).
		Where("is_builtin = 0 AND role_id = ?", roleId).Delete(&pojo.SysRole{}).Error; err != nil {
		roleLog.Error("DelSysRole Error", zap.Error(err))
		return err
	}
	return nil
}

// GetSysRoleById 查询角色详情
func (*SysRoleDao) GetSysRoleById(roleId string) (*pojo.SysRoleDetailResp, error) {
	var role pojo.SysRoleDetailResp
	if err := cwrs_gorm.GormDb.Table(tableSysRole).
		Select("sr.role_id", "sr.role_name", "sr.role_code", "sr.role_status", "sr.role_sort", "sr.dept_id", "sd.dept_name", "sr.desc", "sr.is_builtin", "sr.created_user_id", "su.user_name as created_user_name", "sr.created_time").
		Joins("LEFT JOIN sys_dept sd ON sr.dept_id = sd.dept_id").
		Joins("LEFT JOIN sys_user su ON sr.created_user_id = su.user_id").
		Where("sr.role_id = ?", roleId).First(&role).Error; err != nil {
		roleLog.Error("GetSysRoleById Error", zap.Error(err))
		return nil, err
	}
	return &role, nil
}

// GetSysRoleList 分页查询角色列表
func (*SysRoleDao) GetSysRoleList(c *gin.Context, req *pojo.GetSysRoleListReq) ([]pojo.SysRoleResp, int64, error) {
	var list []pojo.SysRoleResp
	var total int64
	db := cwrs_gorm.GormDb.Table(tableSysRole).Scopes(cwrs_gorm.WithDataScope(c, "sr")).
		Select("sr.role_id", "sr.role_name", "sr.role_code", "sr.role_status", "sr.role_sort", "sr.dept_id", "sd.dept_name", "sr.desc", "sr.is_builtin", "sr.created_user_id", "su.user_name as created_user_name", "sr.created_time").
		Joins("LEFT JOIN sys_dept sd ON sr.dept_id = sd.dept_id").
		Joins("LEFT JOIN sys_user su ON sr.created_user_id = su.user_id")
	if req.DeptId != "" {
		db = db.Where("sr.dept_id = ?", req.DeptId)
	}
	if req.RoleName != "" {
		db = db.Where("sr.role_name LIKE ?", "%"+req.RoleName+"%")
	}
	if req.RoleCode != "" {
		db = db.Where("sr.role_code LIKE ?", "%"+req.RoleCode+"%")
	}
	if req.RoleStatus != "" {
		db = db.Where("sr.role_status = ?", req.RoleStatus)
	}
	db.Count(&total)
	if req.PageNum > 0 && req.PageSize > 0 {
		offset := cwrs_utils.CalcOffset(req.PageNum, req.PageSize)
		db = db.Offset(offset).Limit(req.PageSize)
	}
	if err := db.Order("sr.role_sort").Order("sr.role_sort asc").Find(&list).Error; err != nil {
		roleLog.Error("GetSysRoleList Error", zap.Error(err))
		return nil, 0, err
	}
	return list, total, nil
}

// GetRoleCountByDeptIds 获取组织下角色数量
func (*SysRoleDao) GetRoleCountByDeptIds(deptIds []string) (int64, error) {
	var count int64
	if err := cwrs_gorm.GormDb.Table(tableRole).
		Where("dept_id IN (?)", deptIds).
		Count(&count).Error; err != nil {
		roleLog.Error("GetRoleCountByDeptIds Error", zap.Error(err))
	}
	return count, nil
}
