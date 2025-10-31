package cwrs_base

type UserAuth struct {
	UserId   string `json:"userId"`   // 用户id
	DeptId   string `json:"deptId"`   // 组织id
	RoleId   string `json:"roleId"`   // 角色id
	RoleCode string `json:"roleCode"` // 角色编码
}
