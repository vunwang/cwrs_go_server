package pojo

type SysUserDeptRole struct {
	UserDeptRoleId string `gorm:"column:user_dept_role_id" json:"userDeptRoleId"` // 主键
	UserId         string `gorm:"column:user_id" json:"userId"`                   // 用户id
	OldDeptId      string `gorm:"-" json:"oldDeptId"`                             // 原组织id 修改用
	DeptId         string `gorm:"column:dept_id" json:"deptId"`                   // 所属组织id
	RoleId         string `gorm:"column:role_id" json:"roleId"`                   // 角色id
	PostId         string `gorm:"column:post_id" json:"postId"`                   // 所属岗位
	DataPurview    string `gorm:"column:data_purview" json:"dataPurview"`         // 用户-数据权限(字典)-预留暂时未用
	CreatedUserId  string `gorm:"column:created_user_id" json:"createdUserId"`    // 创建用户
}
