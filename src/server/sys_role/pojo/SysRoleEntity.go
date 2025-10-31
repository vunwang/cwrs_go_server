package pojo

import (
	"cwrs_go_server/src/server/sys_role_menu/pojo"
)

// SysRole  角色。
// 表名:sys_role
type SysRole struct {
	RoleId        string `gorm:"column:role_id" json:"roleId"`                //主键
	RoleName      string `gorm:"column:role_name" json:"roleName"`            //角色名称
	RoleCode      string `gorm:"column:role_code" json:"roleCode"`            //角色编码
	RoleStatus    string `gorm:"column:role_status" json:"roleStatus"`        //状态(字典-sys_status)
	RoleSort      int    `gorm:"column:role_sort" json:"roleSort"`            //排序 从1开始
	IsBuiltin     int    `gorm:"column:is_builtin" json:"isBuiltin"`          //是否内置1是 0否
	Desc          string `gorm:"column:desc" json:"desc"`                     //描述
	DeptId        string `gorm:"column:dept_id" json:"deptId"`                //所属组织
	ParentId      string `gorm:"column:parent_id" json:"parentId"`            //父级id
	RoleLevel     string `gorm:"column:role_level" json:"roleLevel"`          //角色层级
	CreatedUserId string `gorm:"column:created_user_id" json:"createdUserId"` //创建用户
	CreatedTime   string `gorm:"column:created_time" json:"createdTime"`      //创建时间
	UpdatedUserId string `gorm:"column:updated_user_id" json:"updatedUserId"` //更新人
	UpdatedTime   string `gorm:"column:updated_time" json:"updatedTime"`      //更新时间
}

type SysRoleResp struct {
	RoleId          string `gorm:"column:role_id" json:"roleId"`                    //主键
	RoleName        string `gorm:"column:role_name" json:"roleName"`                //角色名称
	RoleCode        string `gorm:"column:role_code" json:"roleCode"`                //角色编码
	RoleStatus      string `gorm:"column:role_status" json:"roleStatus"`            //状态(字典-sys_status)
	RoleSort        int    `gorm:"column:role_sort" json:"roleSort"`                //排序 从1开始
	IsBuiltin       int    `gorm:"column:is_builtin" json:"isBuiltin"`              //是否内置1是 0否
	Desc            string `gorm:"column:desc" json:"desc"`                         //描述
	DeptId          string `gorm:"column:dept_id" json:"deptId"`                    //所属组织
	DeptName        string `gorm:"column:dept_name" json:"deptName"`                //所属组织名称
	CreatedUserId   string `gorm:"column:created_user_id" json:"createdUserId"`     //创建用户
	CreatedUserName string `gorm:"column:created_user_name" json:"createdUserName"` //创建用户名称
	CreatedTime     string `gorm:"column:created_time" json:"createdTime"`          //创建时间
}

type SysRoleDetailResp struct {
	RoleId          string              `gorm:"column:role_id" json:"roleId"`                    //主键
	RoleName        string              `gorm:"column:role_name" json:"roleName"`                //角色名称
	RoleCode        string              `gorm:"column:role_code" json:"roleCode"`                //角色编码
	RoleStatus      string              `gorm:"column:role_status" json:"roleStatus"`            //状态(字典-sys_status)
	RoleSort        int                 `gorm:"column:role_sort" json:"roleSort"`                //排序 从1开始
	IsBuiltin       int                 `gorm:"column:is_builtin" json:"isBuiltin"`              //是否内置1是 0否
	Desc            string              `gorm:"column:desc" json:"desc"`                         //描述
	DeptId          string              `gorm:"column:dept_id" json:"deptId"`                    //所属组织
	DeptName        string              `gorm:"column:dept_name" json:"deptName"`                //所属组织名称
	CreatedUserId   string              `gorm:"column:created_user_id" json:"createdUserId"`     //创建用户
	CreatedUserName string              `gorm:"column:created_user_name" json:"createdUserName"` //创建用户名称
	CreatedTime     string              `gorm:"column:created_time" json:"createdTime"`          //创建时间
	RoleMenus       []pojo.RoleMenuResp `gorm:"-" json:"roleMenus"`                              //平台菜单
	RoleMenusApp    []pojo.RoleMenuResp `gorm:"-" json:"roleMenusApp"`                           //APP菜单
}

// AddSysRoleReq 新增角色入参
// @Description 新增角色
type AddSysRoleReq struct {
	RoleName     string                `form:"roleName" json:"roleName" binding:"required"`     //角色名称
	RoleCode     string                `form:"roleCode" json:"roleCode" binding:"required"`     //角色编码
	RoleStatus   string                `form:"roleStatus" json:"roleStatus" binding:"required"` //状态(字典-sys_status)
	RoleSort     int                   `form:"roleSort" json:"roleSort" binding:"required"`     //排序
	Desc         string                `form:"desc" json:"desc"`                                //描述
	DeptId       string                `form:"deptId" json:"deptId" binding:"required"`         //所属组织
	RoleMenus    []pojo.RoleMenuAddReq `form:"roleMenus" json:"roleMenus"`                      //平台菜单
	RoleMenusApp []pojo.RoleMenuAddReq `form:"roleMenusApp" json:"roleMenusApp"`                //APP菜单
}

// EditSysRoleReq 修改角色入参
// @Description 修改角色
type EditSysRoleReq struct {
	RoleId       string                `form:"roleId" json:"roleId" binding:"required"`         //角色id
	RoleName     string                `form:"roleName" json:"roleName" binding:"required"`     //角色名称
	RoleCode     string                `form:"roleCode" json:"roleCode" binding:"required"`     //角色编码
	RoleStatus   string                `form:"roleStatus" json:"roleStatus" binding:"required"` //状态(字典-sys_status)
	RoleSort     int                   `form:"roleSort" json:"roleSort" binding:"required"`     //排序
	Desc         string                `form:"desc" json:"desc"`                                //描述
	DeptId       string                `form:"deptId" json:"deptId" binding:"required"`         //所属组织
	RoleMenus    []pojo.RoleMenuAddReq `form:"roleMenus" json:"roleMenus"`                      //平台菜单
	RoleMenusApp []pojo.RoleMenuAddReq `form:"roleMenusApp" json:"roleMenusApp"`                //APP菜单
}

// DelSysRoleReq 删除角色入参
// @Description 删除角色
type DelSysRoleReq struct {
	RoleIds string `form:"roleIds" json:"roleIds" binding:"required"` //角色id 多个英文逗号分隔
}

// DelSysRoleReq 查询角色详情入参
// @Description 查询角色详情
type GetSysRoleDetailReq struct {
	RoleId string `form:"roleId" json:"roleId" binding:"required"` //角色id
}

// GetSysRoleListReq 分页查询入参
// @Description 分页查询岗位
type GetSysRoleListReq struct {
	PageNum    int    `form:"pageNum" json:"pageNum"`       // 页码
	PageSize   int    `form:"pageSize" json:"pageSize"`     // 每页显示条数
	DeptId     string `form:"deptId" json:"deptId"`         // 所属组织
	RoleName   string `form:"roleName" json:"roleName"`     // 角色名称(模糊)
	RoleCode   string `form:"roleCode" json:"roleCode"`     // 角色编码(模糊)
	RoleStatus string `form:"roleStatus" json:"roleStatus"` // 状态
}
