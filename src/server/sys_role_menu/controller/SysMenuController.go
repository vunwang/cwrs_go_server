package controller

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/server/sys_role_menu/pojo"
	"cwrs_go_server/src/server/sys_role_menu/service"

	"github.com/gin-gonic/gin"
)

var sysMenuServiceImpl = service.SysMenuService{}

// @Tags 菜单管理【平台】
// @Summary 查询动态菜单树
// @Description 查询当前用户的动态菜单树
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Success 200 {object} cwrs_res.ResSuccessData{data=pojo.SysMenuTreeNode} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/getMenuTreeByRoleId [get]
func GetMenuTreeByRoleId(c *gin.Context) {
	sysMenuServiceImpl.GetMenuTreeByRoleId(c)
}

// @Tags 菜单管理【APP】
// @Summary 查询动态菜单树
// @Description 查询当前用户的动态菜单树
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Success 200 {object} cwrs_res.ResSuccessData{data=[]pojo.SysMenuAppTreeNode} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/getMenuAppTreeByRoleId [get]
func GetMenuAppTreeByRoleId(c *gin.Context) {
	sysMenuServiceImpl.GetMenuAppTreeByRoleId(c)
}

// @Tags 菜单管理【平台】
// @Summary 查询平台菜单树列表
// @Description 查询平台菜单树列表
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Success 200 {object} cwrs_res.ResSuccessData{data=[]pojo.SysMenuTreeNode} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/getMenuTreeList [get]
func GetMenuTreeList(c *gin.Context) {
	sysMenuServiceImpl.GetMenuTreeList(c)
}

// @Tags 菜单管理【平台】
// @Summary 查询APP菜单树列表
// @Description 查询APP菜单树列表
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Success 200 {object} cwrs_res.ResSuccessData{data=[]pojo.SysMenuAppTreeNode} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/getMenuAppTreeList [get]
func GetMenuAppTreeList(c *gin.Context) {
	sysMenuServiceImpl.GetMenuAppTreeList(c)
}

// @Tags 菜单管理【平台】
// @Summary 新增平台目录
// @Description 新增平台目录
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.AddSysMenuDirectoryReq true "新增目录参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/addDirectory [post]
func AddSysMenuDirectory(c *gin.Context) {
	var req pojo.AddSysMenuDirectoryReq
	req.ClientType = 1 //平台
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.AddSysMenu(c, &req)
}

// @Tags 菜单管理【平台】
// @Summary 新增平台菜单
// @Description 新增平台菜单
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.AddSysMenuMenuReq true "新增菜单参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/addMenu [post]
func AddSysMenuMenu(c *gin.Context) {
	var req pojo.AddSysMenuMenuReq
	req.ClientType = 1 //平台
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.AddSysMenu(c, &req)
}

// @Tags 菜单管理【平台】
// @Summary 新增平台按钮
// @Description 新增平台按钮
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.AddSysMenuButtonReq true "新增按钮参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/addButton [post]
func AddSysMenuButton(c *gin.Context) {
	var req pojo.AddSysMenuButtonReq
	req.ClientType = 1 //平台
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.AddSysMenu(c, &req)
}

// @Tags 菜单管理【平台】
// @Summary 新增APP目录
// @Description 新增APP目录
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.AddSysMenuAppDirectoryReq true "新增目录参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/addAppDirectory [post]
func AddSysMenuAppDirectory(c *gin.Context) {
	var req pojo.AddSysMenuAppDirectoryReq
	req.ClientType = 2 //APP
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.AddSysMenu(c, &req)
}

// @Tags 菜单管理【平台】
// @Summary 新增APP菜单
// @Description 新增APP菜单
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.AddSysMenuAppMenuReq true "新增菜单参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/addAppMenu [post]
func AddSysMenuAppMenu(c *gin.Context) {
	var req pojo.AddSysMenuAppMenuReq
	req.ClientType = 2 //APP
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.AddSysMenu(c, &req)
}

// @Tags 菜单管理【平台】
// @Summary 新增APP按钮
// @Description 新增APP按钮
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.AddSysMenuButtonReq true "新增按钮参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/addAppButton [post]
func AddSysMenuAppButton(c *gin.Context) {
	var req pojo.AddSysMenuButtonReq
	req.ClientType = 2 //APP
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.AddSysMenu(c, &req)
}

// @Tags 菜单管理【平台】
// @Summary 批量修改菜单父节点
// @Description 批量修改菜单父节点
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysMenusReq true "修改菜单参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/editMenus [put]
func EditSysMenus(c *gin.Context) {
	var req pojo.EditSysMenusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.EditSysMenus(c, req)
}

// @Tags 菜单管理【平台】
// @Summary 修改状态/是否缓存/是否隐藏
// @Description 修改状态/是否缓存/是否隐藏
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysMenuStatusReq true "修改菜单参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/editMenuStatus [put]
func EditSysMenuStatus(c *gin.Context) {
	var req pojo.EditSysMenuStatusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.EditSysMenuStatus(c, req)
}

// @Tags 菜单管理【平台】
// @Summary 修改平台目录
// @Description 修改平台目录
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysMenuDirectoryReq true "修改目录参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/editDirectory [put]
func EditSysMenuDirectory(c *gin.Context) {
	var req pojo.EditSysMenuDirectoryReq
	req.ClientType = 1 //平台
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.EditSysMenu(c, &req)
}

// @Tags 菜单管理【平台】
// @Summary 修改平台菜单
// @Description 修改平台菜单
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysMenuMenuReq true "修改菜单参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/editMenu [put]
func EditSysMenuMenu(c *gin.Context) {
	var req pojo.EditSysMenuMenuReq
	req.ClientType = 1 //平台
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.EditSysMenu(c, &req)
}

// @Tags 菜单管理【平台】
// @Summary 修改平台按钮
// @Description 修改平台按钮
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysMenuButtonReq true "修改按钮参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/editButton [put]
func EditSysMenuButton(c *gin.Context) {
	var req pojo.EditSysMenuButtonReq
	req.ClientType = 1 //平台
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.EditSysMenu(c, &req)
}

// @Tags 菜单管理【平台】
// @Summary 修改APP目录
// @Description 修改APP目录
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysMenuAppDirectoryReq true "修改目录参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/editAppDirectory [put]
func EditSysMenuAppDirectory(c *gin.Context) {
	var req pojo.EditSysMenuAppDirectoryReq
	req.ClientType = 2 //APP
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.EditSysMenu(c, &req)
}

// @Tags 菜单管理【平台】
// @Summary 修改APP菜单
// @Description 修改APP菜单
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysMenuAppMenuReq true "修改菜单参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/editAppMenu [put]
func EditSysMenuAppMenu(c *gin.Context) {
	var req pojo.EditSysMenuAppMenuReq
	req.ClientType = 2 //APP
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.EditSysMenu(c, &req)
}

// @Tags 菜单管理【平台】
// @Summary 修改APP按钮
// @Description 修改APP按钮
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysMenuButtonReq true "修改按钮参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/editAppButton [put]
func EditSysMenuAppButton(c *gin.Context) {
	var req pojo.EditSysMenuButtonReq
	req.ClientType = 2 //APP
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.EditSysMenu(c, &req)
}

// @Tags 菜单管理【平台】
// @Summary 删除菜单(支持批量)
// @Description 根据菜单主键删除菜单
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.DelSysMenuReq true "删除菜单参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/del [delete]
func DelSysMenu(c *gin.Context) {
	var req pojo.DelSysMenuReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.DelSysMenu(c, &req)
}

// @Tags 菜单管理【平台】
// @Summary 查询菜单详情
// @Description 根据菜单主键查询菜单详情
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysMenuByIdReq true "查询菜单详情参数"
// @Success 200 {object} cwrs_res.ResSuccess{data=pojo.SysMenu} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysMenu/getSysMenuById [get]
func GetSysMenuById(c *gin.Context) {
	var req pojo.GetSysMenuByIdReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysMenuServiceImpl.GetSysMenuById(c, &req)
}
