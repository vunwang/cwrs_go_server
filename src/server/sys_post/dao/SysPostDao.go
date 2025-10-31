package dao

import (
	"cwrs_go_server/src/cwrs_core/cwrs_gorm"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_post/pojo"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

var postLog = cwrs_zap_logger.ZapLogger

const tablePost = "sys_post"
const tableSysPost = "sys_post sp"

type SysPostDao struct{}

// AddSysPost 添加岗位
func (*SysPostDao) AddSysPost(post *pojo.SysPost) error {
	fields := make([]string, 0)
	if post.PostName != "" {
		fields = append(fields, "post_name")
	}
	if post.PostCode != "" {
		fields = append(fields, "post_code")
	}
	if post.PostStatus != "" {
		fields = append(fields, "post_status")
	}
	if post.PostSort != 0 {
		fields = append(fields, "post_sort")
	}
	if post.DeptId != "" {
		fields = append(fields, "dept_id")
	}
	if post.Desc != "" {
		fields = append(fields, "desc")
	}
	fields = append(fields, "post_id", "created_user_id", "created_time")

	if err := cwrs_gorm.GormDb.Table(tablePost).Select(fields).
		Create(&post).Error; err != nil {
		postLog.Error("AddSysPost Error", zap.Error(err))
		return err
	}
	return nil
}

// EditSysPost 修改岗位
func (*SysPostDao) EditSysPost(post *pojo.SysPost) error {
	fields := make([]string, 0)
	if post.PostName != "" {
		fields = append(fields, "post_name")
	}
	if post.PostCode != "" {
		fields = append(fields, "post_code")
	}
	if post.PostStatus != "" {
		fields = append(fields, "post_status")
	}
	if post.PostSort != 0 {
		fields = append(fields, "post_sort")
	}
	if post.DeptId != "" {
		fields = append(fields, "dept_id")
	}
	fields = append(fields, "desc", "updated_user_id", "updated_time")
	if err := cwrs_gorm.GormDb.Table(tablePost).Select(fields).
		Where("post_id = ?", post.PostId).Updates(&post).Error; err != nil {
		postLog.Error("EditSysPost Error", zap.Error(err))
		return err
	}
	return nil
}

// DelSysPost 批量删除岗位
func (*SysPostDao) DelSysPost(postId string) error {
	if err := cwrs_gorm.GormDb.Table(tablePost).
		Where("post_id = ?", postId).Delete(&pojo.SysPost{}).Error; err != nil {
		postLog.Error("DelSysPost Error", zap.Error(err))
		return err
	}
	return nil
}

// GetSysPostById 查询岗位详情
func (*SysPostDao) GetSysPostById(postId string) (*pojo.SysPostResp, error) {
	var post pojo.SysPostResp
	if err := cwrs_gorm.GormDb.Table(tableSysPost).
		Select("sp.post_id", "sp.post_name", "sp.post_code", "sp.post_status", "sp.post_sort", "sp.dept_id", "sd.dept_name", "sp.desc", "su.user_name as created_user_name", "sp.created_time").
		Joins("LEFT JOIN sys_dept sd ON sp.dept_id = sd.dept_id").
		Joins("LEFT JOIN sys_user su ON sp.created_user_id = su.user_id").
		Where("sp.post_id = ?", postId).First(&post).Error; err != nil {
		postLog.Error("GetSysPostById Error", zap.Error(err))
		return nil, err
	}
	return &post, nil
}

// GetSysPostList 分页查询岗位列表
func (*SysPostDao) GetSysPostList(c *gin.Context, req *pojo.GetSysPostListReq) ([]pojo.SysPostResp, int64, error) {
	var list []pojo.SysPostResp
	var total int64
	db := cwrs_gorm.GormDb.Table(tableSysPost).Scopes(cwrs_gorm.WithDataScope(c, "sp")).
		Select("sp.post_id", "sp.post_name", "sp.post_code", "sp.post_status", "sp.post_sort", "sp.dept_id", "sd.dept_name", "sp.desc", "sp.created_time").
		Joins("LEFT JOIN sys_dept sd ON sp.dept_id = sd.dept_id")
	if req.DeptId != "" {
		db = db.Where("sp.dept_id = ?", req.DeptId)
	}
	if req.PostName != "" {
		db = db.Where("sp.post_name LIKE ?", "%"+req.PostName+"%")
	}
	if req.PostCode != "" {
		db = db.Where("sp.post_code LIKE ?", "%"+req.PostCode+"%")
	}
	if req.PostStatus != "" {
		db = db.Where("sp.post_status = ?", req.PostStatus)
	}
	db.Count(&total)
	if req.PageNum > 0 && req.PageSize > 0 {
		offset := cwrs_utils.CalcOffset(req.PageNum, req.PageSize)
		db = db.Offset(offset).Limit(req.PageSize)
	}
	if err := db.Order("sp.post_sort").Find(&list).Error; err != nil {
		postLog.Error("GetSysPostList Error", zap.Error(err))
		return nil, 0, err
	}
	return list, total, nil
}

// GetPostCountByDeptIds 获取组织下岗位数量
func (*SysPostDao) GetPostCountByDeptIds(deptIds []string) (int64, error) {
	var count int64
	if err := cwrs_gorm.GormDb.Table(tablePost).
		Where("dept_id IN (?)", deptIds).Count(&count).Error; err != nil {
		postLog.Error("GetPostCountByDeptIds Error", zap.Error(err))
		return 0, err
	}
	return count, nil
}
