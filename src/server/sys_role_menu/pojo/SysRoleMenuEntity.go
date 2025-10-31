package pojo

// User 角色菜单关联关系
// @Description 角色菜单关联关系
type SysRoleMenu struct {
	RoleMenuId    string `gorm:"column:role_menu_id" json:"roleMenuId"`       //主键
	DeptId        string `gorm:"column:dept_id" json:"deptId"`                //部门id
	RoleId        string `gorm:"column:role_id" json:"roleId"`                //角色id
	MenuType      string `gorm:"column:menu_type" json:"menuType"`            //菜单类型
	MenuId        string `gorm:"column:menu_id" json:"menuId"`                //菜单id
	DataPurview   string `gorm:"column:data_purview" json:"dataPurview"`      //数据权限(字典-sys_data_purview)-按菜单控制
	CreatedUserId string `gorm:"column:created_user_id" json:"createdUserId"` //创建用户
	CreatedTime   string `gorm:"column:created_time" json:"createdTime"`      //创建时间
	UpdatedUserId string `gorm:"column:updated_user_id" json:"updatedUserId"` //更新用户
	UpdatedTime   string `gorm:"column:updated_time" json:"updatedTime"`      //更新时间
}

type RoleMenuResp struct {
	MenuId      string   `gorm:"column:menu_id" json:"menuId"`           //菜单id
	MenuType    string   `gorm:"column:menu_type" json:"menuType"`       //菜单类型
	DataPurview string   `gorm:"column:data_purview" json:"dataPurview"` //数据权限(字典-sys_data_purview)-按菜单控制
	DeptIds     []string `gorm:"-" json:"deptIds"`                       //自定义数据权限授权组织ids
}

type RoleMenuAddReq struct {
	MenuId      string   `form:"menuId" json:"menuId"`           //菜单id
	MenuType    string   `form:"menuType" json:"menuType"`       //菜单类型
	DataPurview string   `form:"dataPurview" json:"dataPurview"` //数据权限(字典-sys_data_purview)-按菜单控制
	DeptIds     []string `form:"deptIds" json:"deptIds"`         //自定义数据权限授权组织ids
}
