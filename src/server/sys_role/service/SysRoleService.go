package service

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/cwrs_utils"
	dataScopeDao "cwrs_go_server/src/server/sys_data_scope/dao"
	dataScopePojo "cwrs_go_server/src/server/sys_data_scope/pojo"
	"cwrs_go_server/src/server/sys_role/dao"
	"cwrs_go_server/src/server/sys_role/pojo"
	roleMenuDao "cwrs_go_server/src/server/sys_role_menu/dao"
	roleMenuPojo "cwrs_go_server/src/server/sys_role_menu/pojo"
	userDao "cwrs_go_server/src/server/sys_user/dao"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

var sysRoleDaoImpl = dao.SysRoleDao{}
var sysUserDaoImpl = userDao.SysUserDao{}
var sysRoleMenuDaoImpl = roleMenuDao.SysRoleMenuDao{}
var sysDataScopeDaoImpl = dataScopeDao.SysDataScopeDao{}

type SysRoleService struct{}

// AddSysRole 新增角色
func (*SysRoleService) AddSysRole(c *gin.Context, req *pojo.AddSysRoleReq) {
	var entity pojo.SysRole
	cwrs_utils.CopyStruct(req, &entity, "json")
	entity.IsBuiltin = 0 // 非内置角色
	entity.RoleId = cwrs_utils.CreateUuid()
	entity.CreatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.CreatedTime = cwrs_utils.GetNowDateTime()
	if err := sysRoleDaoImpl.AddSysRole(&entity); err != nil {
		if strings.Contains(err.Error(), "only_role_code") {
			cwrs_res.Waring(c, err, "角色编码已存在，请重新输入")
			return
		}
		if strings.Contains(err.Error(), "only_dept_role") {
			cwrs_res.Waring(c, err, "角色名称已存在，请重新输入")
			return
		}
		cwrs_res.Waring(c, err, "新增角色失败")
		return
	}
	//添加角色菜单关联关系
	if addRoleMenus(c, "add", req.RoleMenus, req.RoleMenusApp, entity) {
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// addRoleMenus 添加角色菜单关联关系
// 返回值：true表示新增失败，false表示新增成功
func addRoleMenus(c *gin.Context, addOrEdit string, roleMenus []roleMenuPojo.RoleMenuAddReq, roleMenusApp []roleMenuPojo.RoleMenuAddReq, entity pojo.SysRole) bool {
	if addOrEdit == "edit" {
		//删除角色菜单关联关系
		if err := sysRoleMenuDaoImpl.DelSysRoleMenu(entity.RoleId); err != nil {
			cwrs_res.Waring(c, err, "删除角色菜单关联关系失败")
			return true
		}
		//删除自定义数据权限
		if err := sysDataScopeDaoImpl.DelSysDataScope(entity.RoleId); err != nil {
			cwrs_res.Waring(c, err, "删除自定义数据权限失败")
			return true
		}
	}
	var roleMenuReq []roleMenuPojo.SysRoleMenu
	for _, v := range roleMenus {
		var roleMenu roleMenuPojo.SysRoleMenu
		roleMenu.MenuId = v.MenuId
		roleMenu.MenuType = v.MenuType
		roleMenu.RoleId = entity.RoleId
		roleMenu.DeptId = entity.DeptId
		roleMenu.CreatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
		roleMenu.CreatedTime = cwrs_utils.GetNowDateTime()
		roleMenu.UpdatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
		roleMenu.UpdatedTime = cwrs_utils.GetNowDateTime()
		roleMenu.RoleMenuId = cwrs_utils.CreateUuid()
		if v.MenuType != "3" {
			roleMenu.DataPurview = v.DataPurview
			//自定义数据权限处理
			if addDataScope(c, roleMenu, v, entity) {
				return true
			}
		}
		roleMenuReq = append(roleMenuReq, roleMenu)
	}
	for _, v := range roleMenusApp {
		var roleMenuApp roleMenuPojo.SysRoleMenu
		roleMenuApp.MenuId = v.MenuId
		roleMenuApp.MenuType = v.MenuType
		roleMenuApp.RoleId = entity.RoleId
		roleMenuApp.DeptId = entity.DeptId
		roleMenuApp.CreatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
		roleMenuApp.CreatedTime = cwrs_utils.GetNowDateTime()
		roleMenuApp.UpdatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
		roleMenuApp.UpdatedTime = cwrs_utils.GetNowDateTime()
		roleMenuApp.RoleMenuId = cwrs_utils.CreateUuid()
		if v.MenuType != "3" {
			roleMenuApp.DataPurview = v.DataPurview
			//自定义数据权限处理
			if addDataScope(c, roleMenuApp, v, entity) {
				return true
			}
		}
		roleMenuReq = append(roleMenuReq, roleMenuApp)
	}
	if len(roleMenuReq) > 0 {
		if err := sysRoleMenuDaoImpl.AddSysRoleMenu(roleMenuReq); err != nil {
			cwrs_res.Waring(c, err, "新增角色菜单关联关系失败")
			return true
		}
	}
	return false
}

// 新增自定义数据权限
// 返回值：true表示新增失败，false表示新增成功
func addDataScope(c *gin.Context, roleMenuApp roleMenuPojo.SysRoleMenu, v roleMenuPojo.RoleMenuAddReq, entity pojo.SysRole) bool {
	if roleMenuApp.DataPurview == "custom" {
		//添加自定义数据权限
		var dataScopeReq []dataScopePojo.SysDataScope
		for _, deptIdRes := range v.DeptIds {
			var dataScopeRes dataScopePojo.SysDataScope
			dataScopeRes.ScopeId = cwrs_utils.CreateUuid()
			dataScopeRes.ScopeType = "1"
			dataScopeRes.MenuId = v.MenuId
			dataScopeRes.DeptId = deptIdRes
			dataScopeRes.RoleId = entity.RoleId
			dataScopeRes.CreatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
			dataScopeRes.CreatedTime = cwrs_utils.GetNowDateTime()
			dataScopeRes.UpdatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
			dataScopeRes.UpdatedTime = cwrs_utils.GetNowDateTime()
			dataScopeReq = append(dataScopeReq, dataScopeRes)
		}
		if len(dataScopeReq) > 0 {
			if err := sysDataScopeDaoImpl.AddSysDataScope(dataScopeReq); err != nil {
				cwrs_res.Waring(c, err, "新增自定义数据权限失败")
				return true
			}
		}
	}
	return false
}

// EditSysRole 修改角色
func (*SysRoleService) EditSysRole(c *gin.Context, req *pojo.EditSysRoleReq) {
	var entity pojo.SysRole
	cwrs_utils.CopyStruct(req, &entity, "json")
	entity.IsBuiltin = 0 // 非内置角色
	entity.UpdatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.UpdatedTime = cwrs_utils.GetNowDateTime()
	if err := sysRoleDaoImpl.EditSysRole(&entity); err != nil {
		if strings.Contains(err.Error(), "only_role_code") {
			cwrs_res.Waring(c, err, "角色编码已存在，请重新输入")
			return
		}
		if strings.Contains(err.Error(), "only_dept_role") {
			cwrs_res.Waring(c, err, "角色名称已存在，请重新输入")
			return
		}
		cwrs_res.Waring(c, err, "修改角色失败")
		return
	}
	//添加角色菜单关联关系
	if addRoleMenus(c, "edit", req.RoleMenus, req.RoleMenusApp, entity) {
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// DelSysRole 删除角色
func (*SysRoleService) DelSysRole(c *gin.Context, req *pojo.DelSysRoleReq) {
	roleIds := strings.Split(req.RoleIds, ",")
	var err error
	if len(roleIds) == 1 {
		if count, _ := sysUserDaoImpl.GetUserCountByRoleIds(roleIds[0]); count > 0 {
			cwrs_res.Waring(c, err, "该角色已被使用，无法删除")
			return
		}
		err = sysRoleDaoImpl.DelSysRole(roleIds[0])
		if err != nil {
			cwrs_res.Waring(c, err, "删除角色失败")
			return
		}
		cwrs_res.Success(c, "操作成功")
	} else {
		var ids []string
		for _, roleId := range roleIds {
			count, _ := sysUserDaoImpl.GetUserCountByRoleIds(roleId)
			if count > 0 {
				ids = append(ids, roleId)
				continue
			}
			//删除角色主表
			err = sysRoleDaoImpl.DelSysRole(roleId)
			//删除角色菜单关联关系
			err = sysRoleMenuDaoImpl.DelSysRoleMenu(roleId)
			//删除自定义数据权限
			err = sysDataScopeDaoImpl.DelSysDataScope(roleId)
		}
		if err != nil {
			cwrs_res.Waring(c, err, "删除角色失败")
			return
		}
		cwrs_res.Success(c, fmt.Sprintf("已删除 %d 个角色；有%d个已被使用，无法删除！", len(roleIds)-len(ids), len(ids)))
	}
}

// GetSysRoleDetail 查询角色详情
func (*SysRoleService) GetSysRoleDetail(c *gin.Context, req *pojo.GetSysRoleDetailReq) {
	role, err := sysRoleDaoImpl.GetSysRoleById(req.RoleId)
	if err != nil {
		cwrs_res.Waring(c, err, "查询角色详情失败")
		return
	}
	role.RoleMenus, err = sysRoleMenuDaoImpl.GetSysRoleMenuDetail(req.RoleId, 1)
	if err != nil {
		cwrs_res.Waring(c, err, "查询角色平台菜单失败")
		return
	}
	//查询自定义数据权限
	for key, menu := range role.RoleMenus {
		isCustom, deptIds := getCustomDataScope(c, req, menu)
		if isCustom {
			return
		}
		role.RoleMenus[key].DeptIds = deptIds
	}
	role.RoleMenusApp, err = sysRoleMenuDaoImpl.GetSysRoleMenuDetail(req.RoleId, 2)
	if err != nil {
		cwrs_res.Waring(c, err, "查询角色APP菜单失败")
		return
	}
	//查询自定义数据权限
	for keyApp, menuApp := range role.RoleMenusApp {
		getCustomDataScope(c, req, menuApp)
		isCustom, deptIds := getCustomDataScope(c, req, menuApp)
		if isCustom {
			return
		}
		role.RoleMenusApp[keyApp].DeptIds = deptIds
	}
	cwrs_res.SuccessData(c, "操作成功", role)
}

// 获取自定义数据权限
func getCustomDataScope(c *gin.Context, req *pojo.GetSysRoleDetailReq, menu roleMenuPojo.RoleMenuResp) (bool, []string) {
	//判断菜单类型为菜单并且数据权限为自定义
	if menu.MenuType == "2" && menu.DataPurview == "custom" {
		var dataScopeReq dataScopePojo.GetSysDataScopeReq
		dataScopeReq.ScopeType = "1"
		dataScopeReq.RoleId = req.RoleId
		dataScopeReq.MenuId = menu.MenuId

		deptIds, err := sysDataScopeDaoImpl.GetSysDataScopeDeptIds(&dataScopeReq)
		if err != nil {
			cwrs_res.Waring(c, err, "查询自定义数据权限失败")
			return true, nil
		}
		return false, deptIds
	}
	return false, nil
}

// GetSysRoleList 分页查询角色列表
func (*SysRoleService) GetSysRoleList(c *gin.Context, req *pojo.GetSysRoleListReq) {
	list, total, err := sysRoleDaoImpl.GetSysRoleList(c, req)
	if err != nil {
		cwrs_res.Waring(c, err, "查询角色列表失败")
		return
	}
	cwrs_res.SuccessDataList(c, "操作成功", list, total)
}
