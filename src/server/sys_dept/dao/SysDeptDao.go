package dao

import (
	"cwrs_go_server/src/cwrs_core/cwrs_gorm"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/server/sys_dept/pojo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var deptLog = cwrs_zap_logger.ZapLogger

const (
	tableSysDept = "sys_dept"
)

type SysDeptDao struct{}

// AddSysDept 添加组织
func (*SysDeptDao) AddSysDept(dept *pojo.SysDept) error {
	fields := make([]string, 0)
	if dept.DeptName != "" {
		fields = append(fields, "dept_name")
	}
	if dept.ParentId != "" {
		fields = append(fields, "parent_id")
	}
	if dept.DeptStatus != "" {
		fields = append(fields, "dept_status")
	}
	if dept.DeptSort != 0 {
		fields = append(fields, "dept_sort")
	}
	if dept.DeptLevel != "" {
		fields = append(fields, "dept_level")
	}
	fields = append(fields, "dept_id", "created_user_id", "created_time")
	if err := cwrs_gorm.GormDb.Table(tableSysDept).
		Select(fields).
		Create(&dept).Error; err != nil {
		deptLog.Error("AddSysDept Error", zap.Error(err))
		return err
	}
	return nil
}

// EditSysDept 修改组织
func (*SysDeptDao) EditSysDept(dept *pojo.SysDept) error {
	fields := make([]string, 0)
	if dept.DeptName != "" {
		fields = append(fields, "dept_name")
	}
	if dept.ParentId != "" {
		fields = append(fields, "parent_id")
	}
	if dept.DeptStatus != "" {
		fields = append(fields, "dept_status")
	}
	if dept.DeptSort != 0 {
		fields = append(fields, "dept_sort")
	}
	if dept.DeptLevel != "" {
		fields = append(fields, "dept_level")
	}
	fields = append(fields, "updated_user_id", "updated_time")
	if err := cwrs_gorm.GormDb.Table(tableSysDept).
		Select(fields).
		Where("dept_id = ?", dept.DeptId).Updates(&dept).Error; err != nil {
		deptLog.Error("EditSysDept Error", zap.Error(err))
		return err
	}
	return nil
}

// DelSysDept 删除组织
func (*SysDeptDao) DelSysDept(deptIds []string) error {
	if err := cwrs_gorm.GormDb.Table(tableSysDept).
		Where("dept_id IN (?)", deptIds).Delete(&pojo.SysDept{}).Error; err != nil {
		deptLog.Error("DelSysDept Error", zap.Error(err))
		return err
	}
	return nil
}

// GetSysDeptById 根据组织ID查询组织详情
func (*SysDeptDao) GetSysDeptById(deptId string) (*pojo.SysDept, error) {
	var dept pojo.SysDept
	if err := cwrs_gorm.GormDb.Table(tableSysDept).
		Select("dept_id", "dept_name", "parent_id", "dept_status", "dept_level", "dept_sort").
		Where("dept_id = ?", deptId).First(&dept).Error; err != nil {
		deptLog.Error("GetSysDeptById Error", zap.Error(err))
		return nil, err
	}
	return &dept, nil
}

// GetSysDeptList 查询组织列表 数据权限默认为本组织及以下
func (*SysDeptDao) GetSysDeptList(c *gin.Context, req *pojo.GetSysDeptTreeReq) ([]pojo.SysDept, error) {
	var list []pojo.SysDept
	db := cwrs_gorm.GormDb.Table(tableSysDept).Scopes(cwrs_gorm.WithDataScope(c, tableSysDept)).
		Select("dept_id", "dept_name", "parent_id", "dept_status", "dept_sort", "created_time")

	if req.DeptName != "" {
		db = db.Where("dept_name LIKE ?", "%"+req.DeptName+"%")
	}
	if req.DeptStatus != "" {
		db = db.Where("dept_status = ?", req.DeptStatus)
	}

	if err := db.Order("dept_sort").Find(&list).Error; err != nil {
		deptLog.Error("GetSysDeptList Error", zap.Error(err))
		return nil, err
	}
	return list, nil
}

// GetDeptCountByParentIds 查询组织下是否有子组织
func (*SysDeptDao) GetDeptCountByParentIds(parentIds []string) (int64, error) {
	var count int64
	if err := cwrs_gorm.GormDb.Table(tableSysDept).
		Where("parent_id IN (?)", parentIds).
		Count(&count).Error; err != nil {
		deptLog.Error("GetDeptCountByParentIds Error", zap.Error(err))
		return 0, err
	}
	return count, nil
}
