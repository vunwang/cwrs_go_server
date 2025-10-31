package pojo

type SysUser struct {
	UserId        string `gorm:"column:user_id" json:"userId"`                // 主键
	UserPhone     string `gorm:"column:user_phone" json:"userPhone"`          // 手机号
	UserName      string `gorm:"column:user_name" json:"userName"`            // 用户名
	NickName      string `gorm:"column:nick_name" json:"nickName"`            // 昵称
	Password      string `gorm:"column:password" json:"password"`             // 密码
	Avatar        string `gorm:"column:avatar" json:"avatar"`                 // 头像
	Email         string `gorm:"column:email" json:"email"`                   // 邮箱
	Gender        string `gorm:"column:gender" json:"gender"`                 // 性别(字典)
	Birth         string `gorm:"column:birth" json:"birth"`                   // 出生日期
	UserStatus    string `gorm:"column:user_status" json:"userStatus"`        // 状态(字典)
	Signature     string `gorm:"column:signature" json:"signature"`           // 个性签名
	Desc          string `gorm:"column:desc" json:"desc"`                     // 描述
	CreatedUserId string `gorm:"column:created_user_id" json:"createdUserId"` // 创建用户
	CreatedTime   string `gorm:"column:created_time" json:"createdTime"`      // 创建时间
	UpdatedUserId string `gorm:"column:updated_user_id" json:"updatedUserId"` // 更新用户
	UpdatedTime   string `gorm:"column:updated_time" json:"updatedTime"`      // 更新时间
}

type SysUserResp struct {
	UserId          string `gorm:"column:user_id" json:"userId" excel:"用户id"`                                        // 主键
	UserPhone       string `gorm:"column:user_phone" json:"userPhone" excel:"手机号"`                                  // 手机号
	UserName        string `gorm:"column:user_name" json:"userName" excel:"用户名"`                                    // 用户名
	NickName        string `gorm:"column:nick_name" json:"nickName" excel:"昵称"`                                      // 昵称
	Avatar          string `gorm:"column:avatar" json:"avatar" excel:"头像"`                                           // 头像
	Password        string `gorm:"column:password" json:"password"`                                                    // 密码
	Email           string `gorm:"column:email" json:"email" excel:"邮箱"`                                             // 邮箱
	Gender          string `gorm:"column:gender" json:"gender" excel:"性别"`                                           // 性别(字典-sys_gender)
	Birth           string `gorm:"column:birth" json:"birth" excel:"出生日期,format=yyyy-mm-dd"`                       // 出生日期
	UserStatus      string `gorm:"column:user_status" json:"userStatus" excel:"状态"`                                  // 状态(字典-sys_user_status)
	Signature       string `gorm:"column:signature" json:"signature"`                                                  // 个性签名
	CreatedUserId   string `gorm:"column:created_user_id" json:"createdUserId"`                                        // 创建用户
	UpdatedUserId   string `gorm:"column:updated_user_id" json:"updatedUserId"`                                        // 更新用户
	UpdatedTime     string `gorm:"column:updated_time" json:"updatedTime"`                                             // 更新时间
	DeptId          string `gorm:"column:dept_id" json:"deptId"`                                                       // 所属组织id
	DeptName        string `gorm:"column:dept_name" json:"deptName" excel:"所属组织"`                                  // 所属组织名称
	RoleId          string `gorm:"column:role_id" json:"roleId"`                                                       // 角色id
	RoleName        string `gorm:"column:role_name" json:"roleName" excel:"角色名称"`                                  // 角色名称
	PostId          string `gorm:"column:post_id" json:"postId"`                                                       // 岗位
	PostName        string `gorm:"column:post_name" json:"postName" excel:"岗位名称"`                                  // 岗位名称
	Desc            string `gorm:"column:desc" json:"desc" excel:"描述"`                                               // 描述
	CreatedUserName string `gorm:"column:created_user_name" json:"createdUserName" excel:"创建用户名称"`               // 创建用户名称
	CreatedTime     string `gorm:"column:created_time" json:"createdTime" excel:"创建时间,format=yyyy-mm-ss HH:mm:ss"` // 创建时间
}

type SysUserListResp struct {
	UserId          string `gorm:"column:user_id" json:"userId" excel:"用户id"`                                        // 主键
	UserPhone       string `gorm:"column:user_phone" json:"userPhone" excel:"手机号"`                                  // 手机号
	UserName        string `gorm:"column:user_name" json:"userName" excel:"用户名"`                                    // 用户名
	NickName        string `gorm:"column:nick_name" json:"nickName" excel:"昵称"`                                      // 昵称
	Avatar          string `gorm:"column:avatar" json:"avatar" excel:"头像"`                                           // 头像
	Email           string `gorm:"column:email" json:"email" excel:"邮箱"`                                             // 邮箱
	Gender          string `gorm:"column:gender" json:"gender" excel:"性别"`                                           // 性别(字典-sys_gender)
	Birth           string `gorm:"column:birth" json:"birth" excel:"出生日期,format=yyyy-mm-dd"`                       // 出生日期
	UserStatus      string `gorm:"column:user_status" json:"userStatus" excel:"状态"`                                  // 状态(字典-sys_user_status)
	CreatedUserId   string `gorm:"column:created_user_id" json:"createdUserId"`                                        // 创建用户
	DeptId          string `gorm:"column:dept_id" json:"deptId"`                                                       // 所属组织id
	DeptName        string `gorm:"column:dept_name" json:"deptName" excel:"所属组织"`                                  // 所属组织名称
	RoleId          string `gorm:"column:role_id" json:"roleId"`                                                       // 角色id
	RoleName        string `gorm:"column:role_name" json:"roleName" excel:"角色名称"`                                  // 角色名称
	PostId          string `gorm:"column:post_id" json:"postId"`                                                       // 岗位
	PostName        string `gorm:"column:post_name" json:"postName" excel:"岗位名称"`                                  // 岗位名称
	Desc            string `gorm:"column:desc" json:"desc" excel:"描述"`                                               // 描述
	CreatedUserName string `gorm:"column:created_user_name" json:"createdUserName" excel:"创建用户名称"`               // 创建用户名称
	CreatedTime     string `gorm:"column:created_time" json:"createdTime" excel:"创建时间,format=yyyy-mm-ss HH:mm:ss"` // 创建时间
}

// AddSysUserReq 新增用户入参
type AddSysUserReq struct {
	DeptId     string `form:"deptId" json:"deptId" binding:"required" excel:"所属组织"`                 // 所属组织id
	UserPhone  string `form:"userPhone" json:"userPhone" binding:"required" excel:"手机号"`             // 手机号
	UserName   string `form:"userName" json:"userName" binding:"required" excel:"用户名"`               // 用户名
	NickName   string `form:"nickName" json:"nickName" binding:"required" excel:"昵称"`                 // 昵称
	RoleId     string `form:"roleId" json:"roleId" binding:"required" excel:"角色"`                     // 角色id(先选组织根据组织查询角色)
	PostId     string `form:"postId" json:"postId" excel:"岗位"`                                        // 所属岗位(先选组织根据组织查询岗位)
	Gender     string `form:"gender" json:"gender" binding:"required" excel:"性别"`                     // 性别(字典-sys_gender)
	Birth      string `form:"birth" json:"birth" binding:"required" excel:"出生日期,format=yyyy-mm-dd"` // 出生日期
	Avatar     string `form:"avatar" json:"avatar" excel:"头像"`                                        // 头像
	Email      string `form:"email" json:"email" excel:"邮箱"`                                          // 邮箱
	UserStatus string `form:"userStatus" json:"userStatus" binding:"required" excel:"状态"`             // 状态(字典)
	Desc       string `form:"desc" json:"desc" excel:"描述"`                                            // 描述

}

// EditSysUserReq 编辑用户入参
type EditSysUserReq struct {
	UserId     string `form:"userId" json:"userId" binding:"required"`         // 主键
	OldDeptId  string `form:"oldDeptId" json:"oldDeptId" binding:"required"`   // 原组织id
	DeptId     string `form:"deptId" json:"deptId" binding:"required"`         // 目标组织id
	UserPhone  string `form:"userPhone" json:"userPhone" binding:"required"`   // 手机号
	UserName   string `form:"userName" json:"userName" binding:"required"`     // 用户名
	NickName   string `form:"nickName" json:"nickName" binding:"required"`     // 昵称
	RoleId     string `form:"roleId" json:"roleId" binding:"required"`         // 角色id(先选组织 根据组织查询角色)
	PostId     string `form:"postId" json:"postId"`                            // 所属岗位(先选组织根据组织查询岗位)
	Gender     string `form:"gender" json:"gender" binding:"required"`         // 性别(字典)
	Birth      string `form:"birth" json:"birth"`                              // 出生日期
	Avatar     string `form:"avatar" json:"avatar"`                            // 头像
	Email      string `form:"email" json:"email"`                              // 邮箱
	UserStatus string `form:"userStatus" json:"userStatus" binding:"required"` // 状态(字典)
	Desc       string `form:"desc" json:"desc"`                                // 描述
}

// DelSysUserReq 删除用户入参
type DelSysUserReq struct {
	UserIds string `form:"userIds" json:"userIds" binding:"required"` //id 多个英文逗号分隔
}

// GetSysUserDetailReq 查询详情入参
type GetSysUserDetailReq struct {
	UserId string `form:"userId" json:"userId" binding:"required"` //id
}

// GetSysUserListReq 分页查询入参
type GetSysUserListReq struct {
	PageNum   int    `form:"pageNum" json:"pageNum"`     // 页码
	PageSize  int    `form:"pageSize" json:"pageSize"`   // 每页显示条数
	DeptId    string `form:"deptId" json:"deptId"`       // 所属组织
	UserPhone string `form:"userPhone" json:"userPhone"` // 手机号
	UserName  string `form:"userName" json:"userName"`   // 用户名称
	UserType  string `form:"userType" json:"userType"`   // 用户类型(字典-sys_user_type)
	StartTime string `form:"startTime" json:"startTime"` // 创建时间开始
	EndTime   string `form:"endTime" json:"endTime"`     // 创建时间结束
}

// User 管理用户
// @Description 修改管理用户密码入参
type EditUserPwdReq struct {
	OldPwd string `form:"oldPwd" json:"oldPwd" binding:"required"` //旧密码
	NewPwd string `form:"newPwd" json:"newPwd" binding:"required"` //新密码
}

// User 管理用户
// @Description 用户登录入参
type LoginReq struct {
	Account  string `form:"account"  json:"account" binding:"required"`   //用户名/手机号
	Password string `form:"password"  json:"password" binding:"required"` //密码
	DeptId   string `form:"deptId" json:"deptId" binding:"required"`      //组织ID
	RoleId   string `form:"roleId" json:"roleId" binding:"required"`      //角色ID
	RoleCode string `form:"roleCode" json:"roleCode" binding:"required"`  //角色Code
}

// User 管理用户
// @Description 获取用户身份入参
type UserIdentityReq struct {
	Account string `form:"account"  json:"account" binding:"required"` //用户名/手机号
}

// User 管理用户
// @Description 用户登录出参
type LoginResp struct {
	UserId   string `gorm:"column:user_id" json:"userId"`    //用户Id
	DeptId   string `gorm:"column:dept_id" json:"deptId"`    //组织
	RoleId   string `gorm:"column:role_id" json:"roleId"`    //角色
	Password string `gorm:"column:password" json:"password"` //密码
}

// User 管理用户
// @Description 用户信息
type UserIdentityResp struct {
	Label string `gorm:"column:label" json:"label"`
	Value string `gorm:"column:value" json:"value"`
}

// User 管理用户
// @Description 用户信息
type UserInfoResp struct {
	UserId      string   `gorm:"column:user_id" json:"userId"`         //主键
	UserPhone   string   `gorm:"column:user_phone" json:"userPhone"`   //手机号
	UserName    string   `gorm:"column:user_name" json:"userName"`     //用户名
	NickName    string   `gorm:"column:nick_name" json:"nickName"`     //昵称
	Avatar      string   `gorm:"column:avatar" json:"avatar"`          //头像
	Email       string   `gorm:"column:email" json:"email"`            //邮箱
	Gender      string   `gorm:"column:gender" json:"gender"`          //性别(字典)
	Birth       string   `gorm:"column:birth" json:"birth"`            //出生日期
	UserStatus  string   `gorm:"column:user_status" json:"userStatus"` //状态(字典)
	Signature   string   `gorm:"column:signature" json:"signature"`    //个性签名
	Desc        string   `gorm:"column:desc" json:"desc"`              //描述
	DeptId      string   `gorm:"column:dept_id" json:"deptId"`         //组织Id
	RoleId      string   `gorm:"column:role_id" json:"roleId"`         //角色Id
	RoleCode    string   `gorm:"column:role_code" json:"roleCode"`     //角色Code
	Permissions []string `gorm:"-" json:"permissions"`                 //按钮标识
}
