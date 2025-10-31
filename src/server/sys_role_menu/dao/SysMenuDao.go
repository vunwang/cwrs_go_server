package dao

import (
	"cwrs_go_server/src/cwrs_core/cwrs_gorm"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/server/sys_role_menu/pojo"

	"go.uber.org/zap"
)

type SysMenuDao struct{}

var logSysMenu = cwrs_zap_logger.ZapLogger

const (
	tableMenu = "sys_menu"
)

// GetMenusByIds 根据菜单ID列表查询平台菜单详情
func (*SysMenuDao) GetMenusByIds(req *pojo.GetMenuTreeListReq) ([]pojo.SysMenuResp, error) {
	var menus []pojo.SysMenuResp
	query := cwrs_gorm.GormDb.Table(tableMenu)
	if req.MenuType != "all" {
		//菜单树只查状态正常的目录和菜单 菜单管理列表查所有
		query = query.Where("type IN ('1','2')")
		query = query.Where("status = '1'")
	}
	if len(req.MenuIds) != 0 {
		query = query.Where("menu_id IN (?)", req.MenuIds)
	}
	err := query.Where("client_type = 1").Order("sort").Find(&menus).Error
	if err != nil {
		logSysMenu.Error("GetMenusByIds Error", zap.Error(err))
	}
	return menus, err
}

// GetMenusByIds 根据菜单ID列表查询APP菜单详情
func (*SysMenuDao) GetMenusAppByIds(req *pojo.GetMenuTreeListReq) ([]pojo.SysMenuAppResp, error) {
	var menus []pojo.SysMenuAppResp
	query := cwrs_gorm.GormDb.Table(tableMenu)
	if req.MenuType != "all" {
		//菜单树只查状态正常的目录和菜单 菜单管理列表查所有
		query = query.Where("type IN ('1','2')")
		query = query.Where("status = '1'")
	}
	if len(req.MenuIds) != 0 {
		query = query.Where("menu_id IN (?)", req.MenuIds)
	}
	err := query.Where("client_type = 2").Order("sort").Find(&menus).Error
	if err != nil {
		logSysMenu.Error("GetMenusAppByIds Error", zap.Error(err))
	}
	return menus, err
}

// AddSysMenu 添加菜单
func (*SysMenuDao) AddSysMenu(menu *pojo.SysMenu) error {
	tx := cwrs_gorm.GormDb.Table(tableMenu)
	//1平台 2App
	if menu.ClientType == 1 {
		if menu.Type == "1" {
			tx.Select("menu_id", "client_type", "parent_id", "path", "component", "redirect", "type", "title", "svg_icon", "icon", "is_external", "hidden", "keep_alive", "breadcrumb", "always_show", "affix", "status", "sort", "created_user_id", "created_time")
		} else if menu.Type == "2" {
			tx.Select("menu_id", "client_type", "parent_id", "path", "component", "redirect", "type", "title", "svg_icon", "icon", "is_external", "hidden", "keep_alive", "breadcrumb", "showIn_tabs", "affix", "status", "sort", "created_user_id", "created_time")
		} else if menu.Type == "3" {
			tx.Select("menu_id", "client_type", "parent_id", "type", "title", "permission", "status", "sort", "created_user_id", "created_time")
		}
	} else if menu.ClientType == 2 {
		if menu.Type == "3" {
			tx.Select("menu_id", "client_type", "parent_id", "type", "title", "permission", "status", "sort", "created_user_id", "created_time")
		} else {
			tx.Select("menu_id", "client_type", "parent_id", "android_path", "ios_path", "android_component", "ios_component", "type", "title", "icon", "status", "sort", "created_user_id", "created_time")
		}
	}

	err := tx.Create(&menu).Error
	if err != nil {
		logSysMenu.Error("AddSysMenu Error", zap.Error(err))
		return err
	}
	return nil
}

// EditSysMenu 修改菜单
func (*SysMenuDao) EditSysMenu(menu *pojo.SysMenu) error {
	tx := cwrs_gorm.GormDb.Table(tableMenu)
	//1平台 2App
	if menu.ClientType == 1 {
		if menu.Type == "1" {
			tx.Select("menu_id", "client_type", "parent_id", "path", "component", "redirect", "type", "title", "svg_icon", "icon", "is_external", "hidden", "keep_alive", "breadcrumb", "always_show", "affix", "status", "sort", "updated_user_id", "updated_time")
		} else if menu.Type == "2" {
			tx.Select("menu_id", "client_type", "parent_id", "path", "component", "redirect", "type", "title", "svg_icon", "icon", "is_external", "hidden", "keep_alive", "breadcrumb", "showIn_tabs", "affix", "status", "sort", "updated_user_id", "updated_time")
		} else if menu.Type == "3" {
			tx.Select("menu_id", "client_type", "parent_id", "type", "title", "permission", "status", "sort", "updated_user_id", "updated_time")
		}
	} else if menu.ClientType == 2 {
		if menu.Type == "3" {
			tx.Select("menu_id", "client_type", "parent_id", "type", "title", "permission", "status", "sort", "updated_user_id", "updated_time")
		} else {
			tx.Select("menu_id", "client_type", "parent_id", "android_path", "ios_path", "android_component", "ios_component", "type", "title", "icon", "status", "sort", "updated_user_id", "updated_time")
		}
	}

	err := tx.Where("menu_id = ?", menu.MenuId).Updates(&menu).Error
	if err != nil {
		logSysMenu.Error("EditSysMenu Error", zap.Error(err))
		return err
	}
	return nil
}

// EditSysMenus 批量修改菜单父节点
func (*SysMenuDao) EditSysMenus(menu *pojo.EditSysMenusReq) error {
	tx := cwrs_gorm.GormDb.Table(tableMenu)
	tx.Select("parent_id")
	err := tx.Where("menu_id IN (?)", menu.MenuIds).Updates(&menu).Error
	if err != nil {
		logSysMenu.Error("EditSysMenus Error", zap.Error(err))
		return err
	}
	return nil
}

// EditSysMenuStatus 修改菜单状态
func (*SysMenuDao) EditSysMenuStatus(menu *pojo.EditSysMenuStatusReq) error {
	tx := cwrs_gorm.GormDb.Table(tableMenu)
	if menu.Type == "status" {
		tx.Select("status")
	} else if menu.Type == "hidden" {
		tx.Select("hidden")
	} else if menu.Type == "keepAlive" {
		tx.Select("keep_alive")
	}
	err := tx.Where("menu_id = ?", menu.MenuId).Updates(&menu).Error
	if err != nil {
		logSysMenu.Error("EditSysMenuStatus Error", zap.Error(err))
		return err
	}
	return nil
}

// DelSysMenu 删除菜单
func (*SysMenuDao) DelSysMenu(menuId []string) error {
	if err := cwrs_gorm.GormDb.Table(tableMenu).
		Where("menu_id IN (?)", menuId).Delete(&pojo.SysMenu{}).Error; err != nil {
		logSysMenu.Error("DelSysMenu Error", zap.Error(err))
		return err
	}
	return nil
}

// GetMenuCountByParentIds 查询菜单下是否有子菜单
func (*SysMenuDao) GetMenuCountByParentIds(parentIds []string) (int64, error) {
	var count int64
	if err := cwrs_gorm.GormDb.Table(tableMenu).
		Where("parent_id IN (?)", parentIds).
		Count(&count).Error; err != nil {
		logSysMenu.Error("GetMenuCountByParentIds Error", zap.Error(err))
		return 0, err
	}
	return count, nil
}

// GetSysMenuById 根据菜单ID查询菜单详情
func (*SysMenuDao) GetSysMenuById(menuId string) (*pojo.SysMenu, error) {
	var menu pojo.SysMenu
	if err := cwrs_gorm.GormDb.Table(tableMenu).
		Where("menu_id = ?", menuId).First(&menu).Error; err != nil {
		logSysMenu.Error("GetSysMenuById Error", zap.Error(err))
		return nil, err
	}
	return &menu, nil
}
