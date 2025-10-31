package dao

import (
	"cwrs_go_server/src/cwrs_common/cwrs_base"
	"cwrs_go_server/src/cwrs_core/cwrs_gorm"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/server/sys_user/pojo"
	"go.uber.org/zap"
)

type LoginDao struct{}

var loginLog = cwrs_zap_logger.ZapLogger

const (
	tableLoginSysUserSu = "sys_user su"
	tableLoginSysUser   = "sys_user"
	tableLoginSysMenu   = "sys_menu sm"
)

// LoginUser 登录校验
func (*LoginDao) LoginUser(req *pojo.LoginReq) (pojo.LoginResp, error) {
	var userRes pojo.LoginResp
	err := cwrs_gorm.GormDb.Table(tableLoginSysUserSu).Where("su.user_status = '1'").
		Where("su.user_name = ?", req.Account).Or("su.user_phone = ?", req.Account).
		Select("su.user_id", "su.user_name", "su.nick_name", "su.user_phone", "su.avatar", "su.password").First(&userRes).Error
	//打印并记录err信息
	if err != nil {
		loginLog.Error("LoginUser Error", zap.Error(err))
	}
	return userRes, err
}

// GetUserIdentity 获取用户身份信息
func (*LoginDao) GetUserIdentity(req *pojo.UserIdentityReq) ([]pojo.UserIdentityResp, error) {
	var identityRes []pojo.UserIdentityResp
	err := cwrs_gorm.GormDb.Table(tableLoginSysUserSu).
		Joins("inner join sys_user_dept_role sudr on su.user_id = sudr.user_id").
		Joins("left join sys_dept sd on sudr.dept_id = sd.dept_id").
		Joins("left join sys_role sr on sudr.role_id = sr.role_id").
		Where("su.user_status = '1'").Where("su.user_name = ?", req.Account).Or("su.user_phone = ?", req.Account).
		Select("CONCAT(sudr.dept_id, ',', sudr.role_id, ',', sr.role_code) value", "CONCAT(sd.dept_name, '(',sr.role_name,')') label").Find(&identityRes).Error
	//打印并记录err信息
	if err != nil {
		loginLog.Error("GetUserIdentity Error", zap.Error(err))
	}
	return identityRes, err
}

func (*LoginDao) GetUserInfo(userId string) (pojo.UserInfoResp, error) {
	var userInfo pojo.UserInfoResp
	err := cwrs_gorm.GormDb.Table(tableLoginSysUser).
		Where("user_id = ?", userId).
		Select("user_id", "user_name", "nick_name", "user_phone", "avatar", "email", "gender", "birth", "user_status", "desc", "signature").First(&userInfo).Error
	if err != nil {
		loginLog.Error("GetUserInfo Error", zap.Error(err))
	}
	return userInfo, err
}

// 根据角色查询按钮标识
func (*LoginDao) GetPermissions(user *cwrs_base.UserAuth) ([]string, error) {
	var permissions []string
	query := cwrs_gorm.GormDb.Table(tableLoginSysMenu)
	query.Where("sm.status = '1' AND sm.type = '3'")
	if user.RoleCode != "sys_admin" {
		query.Joins("inner join sys_role_menu srm on sm.menu_id = srm.menu_id")
		query.Where("srm.role_id = ?", user.RoleId)
	}
	err := query.Pluck("permission", &permissions).Error
	if err != nil {
		loginLog.Error("GetPermissions Error", zap.Error(err))
	}
	return permissions, err
}
