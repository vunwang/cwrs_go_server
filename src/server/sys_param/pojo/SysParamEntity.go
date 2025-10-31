package pojo

type SysParam struct {
	ParamId          string `gorm:"column:param_id" json:"paramId"`                    // 主键
	SysLogo          string `gorm:"column:sys_logo" json:"sysLogo"`                    // logo(上传)
	SysTitle         string `gorm:"column:sys_title" json:"sysTitle"`                  // 平台名称
	LargeScreenTitle string `gorm:"column:large_screen_title" json:"largeScreenTitle"` // 大屏标题
	DeptId           string `gorm:"column:dept_id" json:"deptId"`                      // 所属组织（all-系统默认）
	CreatedUserId    string `gorm:"column:created_user_id" json:"createdUserId"`       // 创建用户
	CreatedTime      string `gorm:"column:created_time" json:"createdTime"`            // 创建时间
	UpdatedUserId    string `gorm:"column:updated_user_id" json:"updatedUserId"`       // 更新用户
	UpdatedTime      string `gorm:"column:updated_time" json:"updatedTime"`            // 更新时间
}

type SysParamResp struct {
	ParamId          string `gorm:"column:param_id" json:"paramId"`                    // 主键
	SysLogo          string `gorm:"column:sys_logo" json:"sysLogo"`                    // logo(上传)
	SysTitle         string `gorm:"column:sys_title" json:"sysTitle"`                  // 平台名称
	LargeScreenTitle string `gorm:"column:large_screen_title" json:"largeScreenTitle"` // 大屏标题
	DeptId           string `gorm:"column:dept_id" json:"deptId"`                      // 所属组织（all-系统默认）
	CreatedUserId    string `gorm:"column:created_user_id" json:"createdUserId"`       // 创建用户
	CreatedTime      string `gorm:"column:created_time" json:"createdTime"`            // 创建时间
	DeptName         string `gorm:"column:dept_name" json:"deptName"`                  // 所属组织名称
	CreatedUserName  string `gorm:"column:created_user_name" json:"createdUserName"`   // 创建用户名称
}

type SysParamDeptResp struct {
	SysLogo          string `gorm:"column:sys_logo" json:"sysLogo"`                    // logo(上传)
	SysTitle         string `gorm:"column:sys_title" json:"sysTitle"`                  // 平台名称
	LargeScreenTitle string `gorm:"column:large_screen_title" json:"largeScreenTitle"` // 大屏标题
	DeptId           string `gorm:"column:dept_id" json:"deptId"`                      // 所属组织（all-系统默认）
}

// AddSysParamReq 新增系统参数入参
type AddSysParamReq struct {
	SysLogo          string `form:"sysLogo" json:"sysLogo" binding:"required"`                   // logo(上传)
	SysTitle         string `form:"sysTitle" json:"sysTitle" binding:"required"`                 // 平台名称
	LargeScreenTitle string `form:"largeScreenTitle" json:"largeScreenTitle" binding:"required"` // 大屏标题
	DeptId           string `form:"deptId" json:"deptId" binding:"required"`                     // 所属组织
}

// EditSysParamReq 编辑系统参数入参
type EditSysParamReq struct {
	ParamId          string `form:"paramId" json:"paramId" binding:"required"`                   // 主键
	SysLogo          string `form:"sysLogo" json:"sysLogo" binding:"required"`                   // logo(上传)
	SysTitle         string `form:"sysTitle" json:"sysTitle" binding:"required"`                 // 平台名称
	LargeScreenTitle string `form:"largeScreenTitle" json:"largeScreenTitle" binding:"required"` // 大屏标题
	DeptId           string `form:"deptId" json:"deptId" binding:"required"`                     // 所属组织（不可修改 页面禁用）
}

// DelSysParamReq 删除系统参数入参
type DelSysParamReq struct {
	ParamIds string `form:"paramIds" json:"paramIds" binding:"required"` //ids 多个英文逗号分隔
}

// GetSysParamDetailReq 查询详情入参
type GetSysParamDetailReq struct {
	ParamId string `form:"paramId" json:"paramId" binding:"required"` //id
}

// GetSysParamListReq 分页查询入参
type GetSysParamListReq struct {
	PageNum         int    `form:"pageNum" json:"pageNum"`                 // 页码
	PageSize        int    `form:"pageSize" json:"pageSize"`               // 每页显示条数
	DeptId          string `form:"deptId" json:"deptId"`                   // 所属组织
	CreatedUserName string `form:"createdUserName" json:"createdUserName"` // 创建用户名称
}
