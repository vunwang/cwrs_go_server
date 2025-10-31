package pojo

type SysDataScope struct {
	ScopeId       string `gorm:"column:scope_id" json:"scopeId"`              //主键
	ScopeType     string `gorm:"column:scope_type" json:"scopeType"`          //数据权限控制方式(字典 1菜单)
	MenuId        string `gorm:"column:menu_id" json:"menuId"`                //菜单id
	DeptId        string `gorm:"column:dept_id" json:"deptId"`                //组织
	RoleId        string `gorm:"column:role_id" json:"roleId"`                //角色id
	CreatedUserId string `gorm:"column:created_user_id" json:"createdUserId"` //创建用户
	CreatedTime   string `gorm:"column:created_time" json:"createdTime"`      //创建时间
	UpdatedUserId string `gorm:"column:updated_user_id" json:"updatedUserId"` //更新用户
	UpdatedTime   string `gorm:"column:updated_time" json:"updatedTime"`      //更新时间
}

// GetSysDataScopeDetailReq 查询详情入参
// @Description 查询自定义数据权限详情
type GetSysDataScopeReq struct {
	ScopeType string `form:"scopeType" json:"scopeType" binding:"required"` //数据权限控制方式(字典 1菜单)
	MenuId    string `form:"menuId" json:"menuId" binding:"required"`       //菜单id
	RoleId    string `form:"roleId" json:"roleId" binding:"required"`       //角色id
}
