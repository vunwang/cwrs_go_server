package dao

import (
	"cwrs_go_server/src/cwrs_core/cwrs_gorm"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/server/sys_role_menu/pojo"
	"go.uber.org/zap"
)

type SysRoleMenuDao struct{}

var logSysRoleMenu = cwrs_zap_logger.ZapLogger

const (
	tableRoleMenu    = "sys_role_menu"
	tableSysRoleMenu = "sys_role_menu srm"
)

// AddSysRoleMenu 添加角色菜单关系
func (*SysRoleMenuDao) AddSysRoleMenu(req []pojo.SysRoleMenu) error {
	//根据角色id删除角色菜单关系
	if err := cwrs_gorm.GormDb.Table(tableRoleMenu).Where("role_id = ?", req[0].RoleId).Delete(&pojo.SysRoleMenu{}).Error; err != nil {
		logSysRoleMenu.Error("AddSysRoleMenu Error", zap.Error(err))
		return err
	}
	if err := cwrs_gorm.GormDb.Table(tableRoleMenu).
		Select("role_menu_id", "dept_id", "role_id", "menu_type", "menu_id", "data_purview", "created_user_id", "created_time").
		Create(&req).Error; err != nil {
		logSysRoleMenu.Error("AddSysRoleMenu Error", zap.Error(err))
		return err
	}
	return nil
}

// DelSysRoleMenu 根据角色ID删除角色菜单关系
func (*SysRoleMenuDao) DelSysRoleMenu(roleId string) error {
	if err := cwrs_gorm.GormDb.Table(tableRoleMenu).Where("role_id = ?", roleId).Delete(&pojo.SysRoleMenu{}).Error; err != nil {
		logSysRoleMenu.Error("DelSysRoleMenu Error", zap.Error(err))
		return err
	}
	return nil
}

// GetMenuIdsByRoleId 根据角色ID查询菜单ID列表
func (*SysRoleMenuDao) GetMenuIdsByRoleId(roleId, deptId string) ([]string, error) {
	var menuIds []string
	err := cwrs_gorm.GormDb.Table(tableRoleMenu).Where("role_id = ? and dept_id = ?", roleId, deptId).Select("menu_id").Find(&menuIds).Error
	if err != nil {
		logSysRoleMenu.Error("GetMenuIdsByRoleId Error", zap.Error(err))
	}
	return menuIds, err
}

/**
 * @Description: 查询角色权限详情
 * @param roleId
 * @param clientType
 * @return []pojo.RoleMenuAddResp
 * @return error
 */
func (*SysRoleMenuDao) GetSysRoleMenuDetail(roleId string, clientType int) ([]pojo.RoleMenuResp, error) {
	var list []pojo.RoleMenuResp
	db := cwrs_gorm.GormDb.Table(tableSysRoleMenu).
		Select("srm.menu_id", "srm.menu_type", "srm.data_purview").
		Joins("left join sys_menu sm on srm.menu_id = sm.menu_id")
	if err := db.Where("srm.role_id = ? and sm.client_type = ?", roleId, clientType).Find(&list).Error; err != nil {
		logSysRoleMenu.Error("GetSysRoleMenuDetail Error", zap.Error(err))
		return nil, err
	}
	return list, nil
}
