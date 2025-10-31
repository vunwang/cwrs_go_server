package service

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_role_menu/dao"
	"cwrs_go_server/src/server/sys_role_menu/pojo"
	"go.uber.org/zap"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

var sysMenuDaoImpl dao.SysMenuDao
var sysRoleMenuDaoImpl dao.SysRoleMenuDao

type SysMenuService struct{}

// GetMenuTreeByRoleId 根据角色ID递归查询平台菜单树
func (*SysMenuService) GetMenuTreeByRoleId(c *gin.Context) {
	user := cwrs_utils.GetLoginUserInfo(c)
	var menuErr error
	var menus []pojo.SysMenuResp
	var req pojo.GetMenuTreeListReq
	req.MenuType = "menu"
	//判断是否为超管 超管角色标识sys_admin
	if user.RoleCode == "sys_admin" {
		menus, menuErr = sysMenuDaoImpl.GetMenusByIds(&req)
		if menuErr != nil {
			cwrs_res.Waring(c, menuErr, "获取菜单树失败！")
			return
		}
	} else {
		menuIds, err := sysRoleMenuDaoImpl.GetMenuIdsByRoleId(user.RoleId, user.DeptId)
		if err != nil {
			cwrs_res.Waring(c, err, "获取菜单ID列表失败！")
			return
		}
		req.MenuIds = menuIds
		menus, err = sysMenuDaoImpl.GetMenusByIds(&req)
		if err != nil {
			cwrs_res.Waring(c, err, "获取菜单树失败！")
			return
		}
	}
	//根节点parentId为"0"
	treeMenu := buildMenuTreeAuto(menus)
	cwrs_res.SuccessData(c, "获取菜单树成功！", treeMenu)
}

// GetMenuAppTreeByRoleId 根据角色ID递归查询APP菜单树
func (*SysMenuService) GetMenuAppTreeByRoleId(c *gin.Context) {
	user := cwrs_utils.GetLoginUserInfo(c)
	var menuErr error
	var menus []pojo.SysMenuAppResp
	var req pojo.GetMenuTreeListReq
	req.MenuType = "menu"
	//判断是否为超管 超管角色标识sys_admin
	if user.RoleCode == "sys_admin" {
		menus, menuErr = sysMenuDaoImpl.GetMenusAppByIds(&req)
		if menuErr != nil {
			cwrs_res.Waring(c, menuErr, "获取菜单树失败！")
			return
		}
	} else {
		menuIds, err := sysRoleMenuDaoImpl.GetMenuIdsByRoleId(user.RoleId, user.DeptId)
		if err != nil {
			cwrs_res.Waring(c, err, "获取菜单ID列表失败！")
			return
		}
		req.MenuIds = menuIds
		menus, err = sysMenuDaoImpl.GetMenusAppByIds(&req)
		if err != nil {
			cwrs_res.Waring(c, err, "获取菜单树失败！")
			return
		}
	}
	//根节点parentId为"0"
	treeMenu := buildMenuAppTreeAuto(menus)
	cwrs_res.SuccessData(c, "获取菜单树成功！", treeMenu)
}

func (*SysMenuService) GetMenuTreeList(c *gin.Context) {
	user := cwrs_utils.GetLoginUserInfo(c)
	var menuErr error
	var menus []pojo.SysMenuResp
	var req pojo.GetMenuTreeListReq
	req.MenuType = "all"
	//判断是否为超管 超管角色标识sys_admin
	if user.RoleCode == "sys_admin" {
		menus, menuErr = sysMenuDaoImpl.GetMenusByIds(&req)
		if menuErr != nil {
			cwrs_res.Waring(c, menuErr, "获取菜单树失败！")
			return
		}
	} else {
		menuIds, err := sysRoleMenuDaoImpl.GetMenuIdsByRoleId(user.RoleId, user.DeptId)
		if err != nil {
			cwrs_res.Waring(c, err, "获取菜单ID列表失败！")
			return
		}
		req.MenuIds = menuIds
		menus, err = sysMenuDaoImpl.GetMenusByIds(&req)
		if err != nil {
			cwrs_res.Waring(c, err, "获取菜单树失败！")
			return
		}
	}
	//根节点parentId为"0"
	treeMenu := buildMenuTreeAuto(menus)
	cwrs_res.SuccessData(c, "获取菜单树成功！", treeMenu)
}

func (*SysMenuService) GetMenuAppTreeList(c *gin.Context) {
	user := cwrs_utils.GetLoginUserInfo(c)
	var menuErr error
	var menus []pojo.SysMenuAppResp
	var req pojo.GetMenuTreeListReq
	req.MenuType = "all"
	//判断是否为超管 超管角色标识sys_admin
	if user.RoleCode == "sys_admin" {
		menus, menuErr = sysMenuDaoImpl.GetMenusAppByIds(&req)
		if menuErr != nil {
			cwrs_res.Waring(c, menuErr, "获取菜单树失败！")
			return
		}
	} else {
		menuIds, err := sysRoleMenuDaoImpl.GetMenuIdsByRoleId(user.RoleId, user.DeptId)
		if err != nil {
			cwrs_res.Waring(c, err, "获取菜单ID列表失败！")
			return
		}
		req.MenuIds = menuIds
		menus, err = sysMenuDaoImpl.GetMenusAppByIds(&req)
		if err != nil {
			cwrs_res.Waring(c, err, "获取菜单树失败！")
			return
		}
	}
	//根节点parentId为"0"
	treeMenu := buildMenuAppTreeAuto(menus)
	cwrs_res.SuccessData(c, "获取菜单树成功！", treeMenu)
}

// buildMenuTreeAuto 递归组装平台菜单树
func buildMenuTreeAuto(menus []pojo.SysMenuResp) []pojo.SysMenuTreeNode {
	menuMap := make(map[string]pojo.SysMenuResp)
	for _, menu := range menus {
		menuMap[menu.MenuId] = menu
	}

	var roots []pojo.SysMenuResp
	for _, menu := range menus {
		parentId := menu.ParentId
		// 判断是否为根节点
		if parentId == "" || parentId == "0" || menuMap[parentId].MenuId == "" {
			roots = append(roots, menu)
		}
	}

	// 对根节点按 Sort 排序
	sort.Slice(roots, func(i, j int) bool {
		return roots[i].Sort < roots[j].Sort
	})

	// 为每个根节点递归构建子树
	var tree []pojo.SysMenuTreeNode
	for _, root := range roots {
		node := pojo.SysMenuTreeNode{SysMenuResp: root}
		node.Children = buildChildren(menuMap, root.MenuId)
		tree = append(tree, node)
	}

	return tree
}

// buildChildren 递归构建子节点
func buildChildren(menuMap map[string]pojo.SysMenuResp, parentId string) []pojo.SysMenuTreeNode {
	var children []pojo.SysMenuTreeNode
	for _, menu := range menuMap {
		if menu.ParentId == parentId {
			child := pojo.SysMenuTreeNode{SysMenuResp: menu}
			child.Children = buildChildren(menuMap, menu.MenuId)
			children = append(children, child)
		}
	}

	// 对当前层级的子节点按 Sort 排序
	sort.Slice(children, func(i, j int) bool {
		return children[i].SysMenuResp.Sort < children[j].SysMenuResp.Sort
	})
	return children
}

// buildMenuAppTreeAuto 递归组装APP菜单树
func buildMenuAppTreeAuto(menus []pojo.SysMenuAppResp) []pojo.SysMenuAppTreeNode {
	menuMap := make(map[string]pojo.SysMenuAppResp)
	for _, menu := range menus {
		menuMap[menu.MenuId] = menu
	}

	var roots []pojo.SysMenuAppResp
	for _, menu := range menus {
		parentId := menu.ParentId
		// 判断是否为根节点
		if parentId == "" || parentId == "0" || menuMap[parentId].MenuId == "" {
			roots = append(roots, menu)
		}
	}

	// 对根节点按 Sort 排序
	sort.Slice(roots, func(i, j int) bool {
		return roots[i].Sort < roots[j].Sort
	})

	// 为每个根节点递归构建子树
	var tree []pojo.SysMenuAppTreeNode
	for _, root := range roots {
		node := pojo.SysMenuAppTreeNode{SysMenuAppResp: root}
		node.Children = buildAppChildren(menuMap, root.MenuId)
		tree = append(tree, node)
	}

	return tree
}

// buildAppChildren 递归构建子节点
func buildAppChildren(menuMap map[string]pojo.SysMenuAppResp, parentId string) []pojo.SysMenuAppTreeNode {
	var children []pojo.SysMenuAppTreeNode
	for _, menu := range menuMap {
		if menu.ParentId == parentId {
			child := pojo.SysMenuAppTreeNode{SysMenuAppResp: menu}
			child.Children = buildAppChildren(menuMap, menu.MenuId)
			children = append(children, child)
		}
	}

	// 对当前层级的子节点按 Sort 排序
	sort.Slice(children, func(i, j int) bool {
		return children[i].SysMenuAppResp.Sort < children[j].SysMenuAppResp.Sort
	})
	return children
}

// AddSysMenu 新增菜单
func (*SysMenuService) AddSysMenu(c *gin.Context, req interface{}) {
	var entity pojo.SysMenu
	cwrs_utils.CopyStruct(req, &entity, "json")
	entity.MenuId = cwrs_utils.CreateUuid()
	entity.CreatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.CreatedTime = cwrs_utils.GetNowDateTime()

	if err := sysMenuDaoImpl.AddSysMenu(&entity); err != nil {
		if strings.Contains(err.Error(), "only_permission") {
			cwrs_res.Waring(c, err, "权限标识已存在，请重新输入")
			return
		}
		cwrs_res.Waring(c, err, "新增菜单失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// EditSysMenu 修改菜单
func (*SysMenuService) EditSysMenu(c *gin.Context, req interface{}) {
	var entity pojo.SysMenu
	cwrs_utils.CopyStruct(req, &entity, "json")
	entity.UpdatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.UpdatedTime = cwrs_utils.GetNowDateTime()

	if err := sysMenuDaoImpl.EditSysMenu(&entity); err != nil {
		if strings.Contains(err.Error(), "only_permission") {
			cwrs_res.Waring(c, err, "权限标识已存在，请重新输入")
			return
		}
		cwrs_res.Waring(c, err, "修改菜单失败")
		return
	}

	// 如果是app菜单 判断是否删除oss中图片
	resObj, err := sysMenuDaoImpl.GetSysMenuById(entity.MenuId)
	if err != nil {
		cwrs_zap_logger.Error("EditSysMenu-删除oss时获取菜单失败", zap.Error(err))
	} else {
		if resObj.ClientType == 2 && resObj.Icon != entity.Icon {
			err = cwrs_utils.AliyunOssDeleteObject(resObj.Icon)
			if err != nil {
				cwrs_zap_logger.Error("EditSysMenu-删除oss文件失败", zap.Error(err))
			}
		}
	}
	cwrs_res.Success(c, "操作成功")
}

// EditSysMenus 批量修改菜单父节点
func (*SysMenuService) EditSysMenus(c *gin.Context, req pojo.EditSysMenusReq) {
	if err := sysMenuDaoImpl.EditSysMenus(&req); err != nil {
		cwrs_res.Waring(c, err, "修改菜单失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// EditSysMenuStatus 修改菜单状态
func (*SysMenuService) EditSysMenuStatus(c *gin.Context, req pojo.EditSysMenuStatusReq) {
	if err := sysMenuDaoImpl.EditSysMenuStatus(&req); err != nil {
		cwrs_res.Waring(c, err, "修改失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// DelSysMenu 删除菜单
func (*SysMenuService) DelSysMenu(c *gin.Context, req *pojo.DelSysMenuReq) {
	menuIds := strings.Split(req.MenuIds, ",")
	//查询该菜单下是否有子菜单
	menuCount, err := sysMenuDaoImpl.GetMenuCountByParentIds(menuIds)
	if err != nil {
		cwrs_res.Waring(c, err, "操作失败")
		return
	}
	if menuCount > 0 {
		cwrs_res.Forbidden(c, nil, "请先删除子菜单")
		return
	}
	if err = sysMenuDaoImpl.DelSysMenu(menuIds); err != nil {
		cwrs_res.Waring(c, err, "删除菜单失败")
		return
	}
	// 如果是app菜单 判断是否删除oss中图片
	for _, menuId := range menuIds {
		resObj, errRes := sysMenuDaoImpl.GetSysMenuById(menuId)
		if errRes != nil {
			cwrs_zap_logger.Error("DelSysMenu-删除oss文件时获取菜单失败", zap.Error(errRes))
		} else {
			if resObj.ClientType == 2 && resObj.Icon != "" {
				err = cwrs_utils.AliyunOssDeleteObject(resObj.Icon)
				if err != nil {
					cwrs_zap_logger.Error("DelSysMenu-删除oss文件失败", zap.Error(err))
				}
			}
		}
	}
	cwrs_res.Success(c, "操作成功")
}

// GetSysMenuById 根据菜单ID查询菜单详情
func (*SysMenuService) GetSysMenuById(c *gin.Context, req *pojo.GetSysMenuByIdReq) {
	menu, err := sysMenuDaoImpl.GetSysMenuById(req.MenuId)
	if err != nil {
		cwrs_res.Waring(c, err, "获取菜单详情失败")
		return
	}
	cwrs_res.SuccessData(c, "操作成功", menu)
}
