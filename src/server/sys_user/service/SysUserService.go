package service

import (
	"cwrs_go_server/src/cwrs_common/cwrs_constants"
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/cwrs_core/cwrs_gorm"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_user/dao"
	"cwrs_go_server/src/server/sys_user/pojo"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

var sysUserDaoImpl = dao.SysUserDao{}
var sysUserDeptRoleDaoImpl = dao.SysUserDeptRoleDao{}

type SysUserService struct{}

// AddSysUser 新增用户
func (*SysUserService) AddSysUser(c *gin.Context, req *pojo.AddSysUserReq) {
	createdUserId := cwrs_utils.GetLoginUserInfo(c).UserId
	err := addSysUserData(createdUserId, req)

	if err != nil {
		if strings.Contains(err.Error(), "only_user_phone") {
			cwrs_res.Waring(c, err, "手机号已存在，不能重复添加！")
		} else if strings.Contains(err.Error(), "only_user_name") {
			cwrs_res.Waring(c, err, "用户名已存在，不能重复添加！")
		} else if strings.Contains(err.Error(), "uniq_user_dept") {
			cwrs_res.Waring(c, err, "该组织下已存在该用户，不能重复添加！")
		} else {
			cwrs_res.Waring(c, err, "新增用户失败！")
		}
		return
	}
	cwrs_res.Success(c, "操作成功")
}

func addSysUserData(createdUserId string, req *pojo.AddSysUserReq) error {
	var sysUser pojo.SysUser
	//设置密码 密码默认为123456
	sysUser.Password = cwrs_utils.HashPassword(cwrs_constants.DEFAULT_PWD)
	//生成uuid
	sysUser.UserId = cwrs_utils.CreateUuid()
	sysUser.UserPhone = req.UserPhone
	sysUser.UserName = req.UserName
	sysUser.NickName = req.NickName
	sysUser.Avatar = req.Avatar
	sysUser.Email = req.Email
	sysUser.Gender = req.Gender
	sysUser.Birth = req.Birth
	sysUser.UserStatus = req.UserStatus
	sysUser.Signature = "工程师 | 值班中"
	sysUser.Desc = req.Desc
	//创建用户id
	sysUser.CreatedUserId = createdUserId
	//设置创建时间
	sysUser.CreatedTime = cwrs_utils.GetNowDateTime()

	//添加用户组织角色关联信息数据处理
	var userDeptRole pojo.SysUserDeptRole
	userDeptRole.UserDeptRoleId = cwrs_utils.CreateUuid()
	userDeptRole.UserId = sysUser.UserId
	userDeptRole.RoleId = req.RoleId
	userDeptRole.DeptId = req.DeptId
	userDeptRole.PostId = req.PostId
	userDeptRole.CreatedUserId = createdUserId

	err := addUserData(sysUser, userDeptRole)
	return err
}

// 添加用户相关数据到数据库
func addUserData(sysUser pojo.SysUser, userDeptRole pojo.SysUserDeptRole) error {
	return cwrs_gorm.GormDb.Transaction(func(tx *gorm.DB) error {
		//添加用户基本信息
		if err := sysUserDaoImpl.AddSysUser(tx, &sysUser); err != nil {
			return err
		}
		//添加用户组织角色关联信息
		if err := sysUserDeptRoleDaoImpl.AddSysUserDeptRole(tx, &userDeptRole); err != nil {
			return err
		}
		return nil
	})
}

// 重置密码
func (*SysUserService) ResetUserPassword(c *gin.Context, req *pojo.GetSysUserDetailReq) {
	var sysUser pojo.SysUser
	sysUser.UserId = req.UserId
	//设置密码 密码默认为123456
	sysUser.Password = cwrs_utils.HashPassword(cwrs_constants.DEFAULT_PWD)
	err := sysUserDaoImpl.EditUserPassword(sysUser)
	if err != nil {
		cwrs_res.Waring(c, err, "操作失败！")
		return
	}
	cwrs_res.Success(c, fmt.Sprintf("密码已重置为：%s)", cwrs_constants.DEFAULT_PWD))
}

// 修改用户密码
func (*SysUserService) EditUserPassword(c *gin.Context, req *pojo.EditUserPwdReq) {
	var sysUser pojo.SysUser
	sysUser.UserId = cwrs_utils.GetLoginUserInfo(c).UserId
	sysUser.Password = cwrs_utils.HashPassword(req.NewPwd)
	//查询旧密码是否正确
	userResp, err := sysUserDaoImpl.GetSysUserById(sysUser.UserId)
	if err != nil {
		cwrs_res.Waring(c, err, "获取用户信息失败！")
		return
	}
	//密码验证
	ok := cwrs_utils.CheckPasswordHash(req.OldPwd, userResp.Password)
	if !ok {
		cwrs_res.Waring(c, nil, "旧密码错误！")
		return
	}
	err = sysUserDaoImpl.EditUserPassword(sysUser)
	if err != nil {
		cwrs_res.Waring(c, err, "操作失败！")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// EditSysUser 修改用户
func (*SysUserService) EditSysUser(c *gin.Context, req *pojo.EditSysUserReq) {
	var sysUser pojo.SysUser
	sysUser.UserId = req.UserId
	sysUser.UserPhone = req.UserPhone
	sysUser.UserName = req.UserName
	sysUser.NickName = req.NickName
	sysUser.Avatar = req.Avatar
	sysUser.Email = req.Email
	sysUser.Gender = req.Gender
	sysUser.Birth = req.Birth
	sysUser.UserStatus = req.UserStatus
	sysUser.Signature = "工程师 | 值班中"
	sysUser.Desc = req.Desc
	sysUser.UpdatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	sysUser.UpdatedTime = cwrs_utils.GetNowDateTime()

	//修改用户组织角色关联信息数据处理
	var userDeptRole pojo.SysUserDeptRole
	userDeptRole.UserDeptRoleId = cwrs_utils.CreateUuid()
	userDeptRole.UserId = sysUser.UserId
	userDeptRole.RoleId = req.RoleId
	userDeptRole.OldDeptId = req.OldDeptId
	userDeptRole.DeptId = req.DeptId
	userDeptRole.PostId = req.PostId

	if err := editUserData(sysUser, userDeptRole); err != nil {
		if strings.Contains(err.Error(), "only_user_phone") {
			cwrs_res.Waring(c, err, "手机号已存在，不能重复添加！")
		} else if strings.Contains(err.Error(), "only_user_name") {
			cwrs_res.Waring(c, err, "用户名已存在，不能重复添加！")
		} else if strings.Contains(err.Error(), "uniq_user_dept") {
			cwrs_res.Waring(c, err, "该组织下已存在该用户，不能重复添加！")
		} else {
			cwrs_res.Waring(c, err, "修改用户失败！")
		}
		return
	}

	cwrs_res.Success(c, "操作成功")
}

// 添加用户相关数据到数据库
func editUserData(sysUser pojo.SysUser, userDeptRole pojo.SysUserDeptRole) error {
	return cwrs_gorm.GormDb.Transaction(func(tx *gorm.DB) error {
		//修改用户基本信息
		if err := sysUserDaoImpl.EditSysUser(tx, &sysUser); err != nil {
			return err
		}
		//修改用户组织角色关联信息
		if err := sysUserDeptRoleDaoImpl.EditSysUserDeptRole(tx, &userDeptRole); err != nil {
			return err
		}
		return nil
	})
}

// DelSysUser 删除用户
func (*SysUserService) DelSysUser(c *gin.Context, req *pojo.DelSysUserReq) {
	userIds := strings.Split(req.UserIds, ",")
	err := cwrs_gorm.GormDb.Transaction(func(tx *gorm.DB) error {
		//删除用户
		if err := sysUserDaoImpl.DelSysUser(tx, userIds); err != nil {
			return err
		}
		//删除用户组织角色关联信息
		if err := sysUserDeptRoleDaoImpl.DelSysUserDeptRole(tx, userIds); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		cwrs_res.Waring(c, err, "删除用户失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// GetSysUserDetail 查询用户详情
func (*SysUserService) GetSysUserDetail(c *gin.Context, req *pojo.GetSysUserDetailReq) {
	item, err := sysUserDaoImpl.GetSysUserById(req.UserId)
	if err != nil {
		cwrs_res.Waring(c, err, "查询用户详情失败")
		return
	}
	cwrs_res.SuccessData(c, "操作成功", item)
}

// GetSysUserList 分页查询用户列表
func (*SysUserService) GetSysUserList(c *gin.Context, req *pojo.GetSysUserListReq) {
	list, total, err := sysUserDaoImpl.GetSysUserList(c, req)
	if err != nil {
		cwrs_res.Waring(c, err, "查询用户列表失败")
		return
	}
	cwrs_res.SuccessDataList(c, "操作成功", list, total)
}
