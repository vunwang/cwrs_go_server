package dao

import (
	"cwrs_go_server/src/cwrs_core/cwrs_gorm"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/server/sys_data_scope/pojo"
	"go.uber.org/zap"
)

var dataScopeLog = cwrs_zap_logger.ZapLogger

const tableDataScope = "sys_data_scope"

type SysDataScopeDao struct{}

// AddSysDataScope 添加自定义数据权限
func (*SysDataScopeDao) AddSysDataScope(req []pojo.SysDataScope) error {
	if err := cwrs_gorm.GormDb.Table(tableDataScope).
		Select("scope_id", "scope_type", "menu_id", "dept_id", "role_id", "created_user_id", "created_time", "updated_user_id", "updated_time").
		Create(&req).Error; err != nil {
		dataScopeLog.Error("AddSysDataScope Error", zap.Error(err))
		return err
	}
	return nil
}

func (*SysDataScopeDao) DelSysDataScope(roleId string) error {
	if err := cwrs_gorm.GormDb.Table(tableDataScope).
		Where("role_id = ?", roleId).Delete(&pojo.SysDataScope{}).Error; err != nil {
		dataScopeLog.Error("DelSysDataScope Error", zap.Error(err))
		return err
	}
	return nil
}

// GetSysDataScopeDeptIds 查询自定义数据权限授权组织
func (*SysDataScopeDao) GetSysDataScopeDeptIds(req *pojo.GetSysDataScopeReq) ([]string, error) {
	var deptIds []string
	db := cwrs_gorm.GormDb.Table(tableDataScope).
		Select("dept_id dept_ids")
	if err := db.Where("scope_type = ? and menu_id = ? and role_id = ?", req.ScopeType, req.MenuId, req.RoleId).Find(&deptIds).Error; err != nil {
		dataScopeLog.Error("GetSysDataScopeDeptIds Error", zap.Error(err))
		return nil, err
	}
	return deptIds, nil
}
