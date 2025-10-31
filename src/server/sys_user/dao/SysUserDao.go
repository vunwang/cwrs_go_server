package dao

import (
	"cwrs_go_server/src/cwrs_core/cwrs_gorm"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_user/pojo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var sysUserLog = cwrs_zap_logger.ZapLogger

const tableSysUser = "sys_user"
const tableSelSysUser = "sys_user su"

type SysUserDao struct{}

func (*SysUserDao) AddSysUser(tx *gorm.DB, sysUser *pojo.SysUser) error {
	fields := make([]string, 0)
	if sysUser.UserPhone != "" {
		fields = append(fields, "user_phone")
	}
	if sysUser.UserName != "" {
		fields = append(fields, "user_name")
	}
	if sysUser.NickName != "" {
		fields = append(fields, "nick_name")
	}
	if sysUser.Password != "" {
		fields = append(fields, "password")
	}
	if sysUser.Avatar != "" {
		fields = append(fields, "avatar")
	}
	if sysUser.Email != "" {
		fields = append(fields, "email")
	}
	if sysUser.Gender != "" {
		fields = append(fields, "gender")
	}
	if sysUser.Birth != "" {
		fields = append(fields, "birth")
	}
	if sysUser.UserStatus != "" {
		fields = append(fields, "user_status")
	}
	if sysUser.Signature != "" {
		fields = append(fields, "signature")
	}
	if sysUser.Desc != "" {
		fields = append(fields, "desc")
	}
	fields = append(fields, "user_id", "created_user_id", "created_time")
	if err := tx.Table(tableSysUser).Select(fields).Create(&sysUser).Error; err != nil {
		sysUserLog.Error("AddSysUser Error", zap.Error(err))
		return err
	}
	return nil
}

func (*SysUserDao) EditUserPassword(sysUser pojo.SysUser) error {
	err := cwrs_gorm.GormDb.Table(tableSysUser).
		Select("password").
		Where("user_id = ?", sysUser.UserId).
		Updates(sysUser).Error
	if err != nil {
		sysUserLog.Error("EditUserPassword Error", zap.Error(err))
	}
	return err
}

// EditSysUser 修改用户
func (*SysUserDao) EditSysUser(tx *gorm.DB, item *pojo.SysUser) error {
	fields := make([]string, 0)
	if item.UserPhone != "" {
		fields = append(fields, "user_phone")
	}
	if item.UserName != "" {
		fields = append(fields, "user_name")
	}
	if item.NickName != "" {
		fields = append(fields, "nick_name")
	}
	if item.Password != "" {
		fields = append(fields, "password")
	}
	if item.Avatar != "" {
		fields = append(fields, "avatar")
	}
	if item.Gender != "" {
		fields = append(fields, "gender")
	}
	if item.Birth != "" {
		fields = append(fields, "birth")
	}
	if item.UserStatus != "" {
		fields = append(fields, "user_status")
	}
	if item.Signature != "" {
		fields = append(fields, "signature")
	}
	fields = append(fields, "desc")
	fields = append(fields, "email", "updated_user_id", "updated_time")
	if err := tx.Table(tableSysUser).Select(fields).
		Where("user_id = ?", item.UserId).Updates(&item).Error; err != nil {
		sysUserLog.Error("EditSysUser Error", zap.Error(err))
		return err
	}
	return nil
}

// DelSysUser 删除用户
func (*SysUserDao) DelSysUser(tx *gorm.DB, userIds []string) error {
	if err := tx.Table(tableSysUser).
		Where("user_id IN (?)", userIds).Delete(&pojo.SysUser{}).Error; err != nil {
		sysUserLog.Error("DelSysUser Error", zap.Error(err))
		return err
	}
	return nil
}

// GetSysUserById 查询用户详情
func (*SysUserDao) GetSysUserById(userId string) (*pojo.SysUserResp, error) {
	var item pojo.SysUserResp
	if err := cwrs_gorm.GormDb.Table(tableSelSysUser).
		Select("su.user_id", "su.user_phone", "su.user_name", "su.nick_name", "su.avatar", "su.email", "su.gender", "su.birth", "su.user_status", "su.desc",
			"su.created_user_id", "su.created_time", "su1.user_name as created_user_name", "sudr.dept_id", "sd.dept_name", "sudr.post_id", "sp.post_name", "sudr.role_id", "sr.role_name").
		Joins("LEFT JOIN sys_user su1 ON su.user_id = su1.user_id").
		Joins("LEFT JOIN sys_user_dept_role sudr ON su.user_id = sudr.user_id").
		Joins("LEFT JOIN sys_dept sd ON sudr.dept_id = sd.dept_id").
		Joins("LEFT JOIN sys_post sp ON sudr.post_id = sp.post_id").
		Joins("LEFT JOIN sys_role sr ON sudr.role_id = sr.role_id").
		Where("su.user_id = ?", userId).First(&item).Error; err != nil {
		sysUserLog.Error("GetSysUserById Error", zap.Error(err))
		return nil, err
	}
	return &item, nil
}

// GetSysUserList 分页查询用户列表
func (*SysUserDao) GetSysUserList(c *gin.Context, req *pojo.GetSysUserListReq) ([]pojo.SysUserListResp, int64, error) {
	var list []pojo.SysUserListResp
	var total int64
	db := cwrs_gorm.GormDb.Table(tableSelSysUser).Scopes(cwrs_gorm.WithDataScope(c, "sudr")).
		Select("su.user_id", "su.user_phone", "su.user_name", "su.nick_name", "su.avatar", "su.email", "su.gender", "su.birth", "su.user_status", "su.desc",
			"su.created_user_id", "su.created_time", "su1.user_name as created_user_name", "sudr.dept_id", "sd.dept_name", "sudr.post_id", "sp.post_name", "sudr.role_id", "sr.role_name").
		Joins("LEFT JOIN sys_user su1 ON su.created_user_id = su1.user_id").
		Joins("LEFT JOIN sys_user_dept_role sudr ON su.user_id = sudr.user_id").
		Joins("LEFT JOIN sys_dept sd ON sudr.dept_id = sd.dept_id").
		Joins("LEFT JOIN sys_post sp ON sudr.post_id = sp.post_id").
		Joins("LEFT JOIN sys_role sr ON sudr.role_id = sr.role_id")
	if req.DeptId != "" {
		db = db.Where("sudr.dept_id = ?", req.DeptId)
	}
	if req.UserPhone != "" {
		db = db.Where("su.user_phone LIKE ?", "%"+req.UserPhone+"%")
	}
	if req.UserName != "" {
		db = db.Where("su.user_name LIKE ?", "%"+req.UserName+"%")
	}
	if req.UserType != "" {
		db = db.Where("FIND_IN_SET(?,sud.user_type) > 0", req.UserType)
	}
	if req.StartTime != "" && req.EndTime != "" {
		db = db.Where("su.created_time BETWEEN ? AND ?", req.StartTime, req.EndTime)
	}
	db.Count(&total)
	if req.PageNum > 0 && req.PageSize > 0 {
		offset := cwrs_utils.CalcOffset(req.PageNum, req.PageSize)
		db = db.Offset(offset).Limit(req.PageSize)
	}
	if err := db.Order("su.created_time desc").Find(&list).Error; err != nil {
		sysUserLog.Error("GetSysUserList Error", zap.Error(err))
		return nil, 0, err
	}
	return list, total, nil
}

// GetUserCountByDeptIds 获取组织下用户数量
func (*SysUserDao) GetUserCountByDeptIds(deptIds []string) (int64, error) {
	var count int64
	if err := cwrs_gorm.GormDb.Table(tableSysUser).
		Joins("LEFT JOIN sys_user_dept_role sudr ON sys_user.user_id = sudr.user_id").
		Where("sudr.dept_id IN (?)", deptIds).
		Count(&count).Error; err != nil {
		sysUserLog.Error("GetUserCountByDeptIds Error", zap.Error(err))
	}
	return count, nil
}

// GetUserCountByPostIds 获取岗位下用户数量
func (*SysUserDao) GetUserCountByPostIds(postId string) (int64, error) {
	var count int64
	if err := cwrs_gorm.GormDb.Table(tableSysUser).
		Joins("LEFT JOIN sys_user_dept_role sudr ON sys_user.user_id = sudr.user_id").
		Where("sudr.post_id = ?", postId).
		Count(&count).Error; err != nil {
		sysUserLog.Error("GetUserCountByPostIds Error", zap.Error(err))
	}
	return count, nil
}

// GetUserCountByRoleIds 获取角色下用户数量
func (*SysUserDao) GetUserCountByRoleIds(roleId string) (int64, error) {
	var count int64
	if err := cwrs_gorm.GormDb.Table(tableSysUser).
		Joins("LEFT JOIN sys_user_dept_role sudr ON sys_user.user_id = sudr.user_id").
		Where("sudr.role_id = ?", roleId).
		Count(&count).Error; err != nil {
		sysUserLog.Error("GetUserCountByRoleIds Error", zap.Error(err))
	}
	return count, nil
}
