package cwrs_routes

import (
	sysDeptController "cwrs_go_server/src/server/sys_dept/controller"
	sysDictController "cwrs_go_server/src/server/sys_dict/controller"
	sysOssController "cwrs_go_server/src/server/sys_oss/controller"
	sysPostController "cwrs_go_server/src/server/sys_post/controller"
	sysRoleController "cwrs_go_server/src/server/sys_role/controller"
	sysMenuController "cwrs_go_server/src/server/sys_role_menu/controller"
	sysUserController "cwrs_go_server/src/server/sys_user/controller"
	"github.com/gin-gonic/gin"

	sysOperLogController "cwrs_go_server/src/server/sys_oper_log/controller"
	sysParamController "cwrs_go_server/src/server/sys_param/controller"
	sysTaskController "cwrs_go_server/src/server/sys_task/controller"
)

// NotAuthRoutes 不用鉴权接口
func NotAuthRoutes(router *gin.RouterGroup) {
	router.GET("/sysLogin/login", sysUserController.LoginUser)
	router.GET("/sysLogin/getUserIdentity", sysUserController.GetUserIdentity)

	// 前端初始化数据字典
	router.GET("/sysDict/allMap", sysDictController.GetAllDictMap)
}

// AuthRoutes 鉴权接口
func AuthRoutes(router *gin.RouterGroup) {
	//文件管理
	router.GET("/sysOss/upload", sysOssController.AddSysOssUpload)
	//router.GET("/sysOss/fileList", sysOssController.GetSysOssList)
	router.GET("/sysOss/dirList", sysOssController.GetSysOssDirList)
	router.DELETE("/sysOss/del", sysOssController.DeleteSysOss)

	//退出登录
	router.GET("/sysLogin/logout", sysUserController.Logout)
	router.GET("/sysLogin/getUserInfo", sysUserController.GetUserInfo)

	router.GET("/sysMenu/getMenuTreeByRoleId", sysMenuController.GetMenuTreeByRoleId)
	router.GET("/sysMenu/getMenuTreeList", sysMenuController.GetMenuTreeList)
	router.GET("/sysMenu/getMenuAppTreeByRoleId", sysMenuController.GetMenuAppTreeByRoleId)
	router.GET("/sysMenu/getMenuAppTreeList", sysMenuController.GetMenuAppTreeList)

	// 角色管理相关接口
	router.POST("/sysRole/add", sysRoleController.AddSysRole)
	router.PUT("/sysRole/edit", sysRoleController.EditSysRole)
	router.DELETE("/sysRole/del", sysRoleController.DelSysRole)
	router.GET("/sysRole/detail", sysRoleController.GetSysRoleDetail)
	router.GET("/sysRole/list", sysRoleController.GetSysRoleList)

	// 菜单管理相关接口
	router.GET("/sysMenu/getSysMenuById", sysMenuController.GetSysMenuById)
	router.POST("/sysMenu/addDirectory", sysMenuController.AddSysMenuDirectory)
	router.POST("/sysMenu/addMenu", sysMenuController.AddSysMenuMenu)
	router.POST("/sysMenu/addButton", sysMenuController.AddSysMenuButton)
	router.POST("/sysMenu/addAppDirectory", sysMenuController.AddSysMenuAppDirectory)
	router.POST("/sysMenu/addAppMenu", sysMenuController.AddSysMenuAppMenu)
	router.POST("/sysMenu/addAppButton", sysMenuController.AddSysMenuAppButton)
	router.PUT("/sysMenu/editMenus", sysMenuController.EditSysMenus)
	router.PUT("/sysMenu/editMenuStatus", sysMenuController.EditSysMenuStatus)
	router.PUT("/sysMenu/editDirectory", sysMenuController.EditSysMenuDirectory)
	router.PUT("/sysMenu/editMenu", sysMenuController.EditSysMenuMenu)
	router.PUT("/sysMenu/editButton", sysMenuController.EditSysMenuButton)
	router.PUT("/sysMenu/editAppDirectory", sysMenuController.EditSysMenuAppDirectory)
	router.PUT("/sysMenu/editAppMenu", sysMenuController.EditSysMenuAppMenu)
	router.PUT("/sysMenu/editAppButton", sysMenuController.EditSysMenuAppButton)
	router.DELETE("/sysMenu/del", sysMenuController.DelSysMenu)

	// 组织管理相关接口
	router.POST("/sysDept/add", sysDeptController.AddSysDept)
	router.PUT("/sysDept/edit", sysDeptController.EditSysDept)
	router.DELETE("/sysDept/del", sysDeptController.DelSysDept)
	router.GET("/sysDept/detail", sysDeptController.GetSysDeptDetail)
	router.GET("/sysDept/tree", sysDeptController.GetSysDeptTree)

	// 数据字典相关接口
	router.POST("/sysDict/add", sysDictController.AddSysDict)
	router.PUT("/sysDict/edit", sysDictController.EditSysDict)
	router.DELETE("/sysDict/del", sysDictController.DelSysDict)
	router.GET("/sysDict/detail", sysDictController.GetSysDictDetail)
	router.GET("/sysDict/list", sysDictController.GetSysDictList)

	// 数据字典项相关接口
	router.POST("/sysDictItem/add", sysDictController.AddSysDictItem)
	router.PUT("/sysDictItem/edit", sysDictController.EditSysDictItem)
	router.DELETE("/sysDictItem/del", sysDictController.DelSysDictItem)
	router.GET("/sysDictItem/list", sysDictController.GetSysDictItemList)
	router.GET("/sysDictItem/detail", sysDictController.GetSysDictItemDetail)

	// 岗位管理相关接口
	router.POST("/sysPost/add", sysPostController.AddSysPost)
	router.PUT("/sysPost/edit", sysPostController.EditSysPost)
	router.DELETE("/sysPost/del", sysPostController.DelSysPost)
	router.GET("/sysPost/detail", sysPostController.GetSysPostDetail)
	router.GET("/sysPost/list", sysPostController.GetSysPostList)

	// 系统参数相关接口
	router.POST("/sysParam/add", sysParamController.AddSysParam)
	router.PUT("/sysParam/edit", sysParamController.EditSysParam)
	router.DELETE("/sysParam/del", sysParamController.DelSysParam)
	router.GET("/sysParam/detail", sysParamController.GetSysParamDetail)
	router.GET("/sysParam/list", sysParamController.GetSysParamList)
	router.GET("/sysParam/dept", sysParamController.GetSysParamByDeptId)

	// 定时任务相关接口
	router.POST("/sysTask/add", sysTaskController.AddSysTask)
	router.PUT("/sysTask/edit", sysTaskController.EditSysTask)
	router.DELETE("/sysTask/del", sysTaskController.DelSysTask)
	router.GET("/sysTask/detail", sysTaskController.GetSysTaskDetail)
	router.GET("/sysTask/list", sysTaskController.GetSysTaskList)
	router.GET("/sysTask/start", sysTaskController.StartTask)
	router.GET("/sysTask/stop", sysTaskController.StopTask)

	// 操作日志相关接口
	router.GET("/sysOperLog/detail", sysOperLogController.GetSysOperLogDetail)
	router.GET("/sysOperLog/list", sysOperLogController.GetSysOperLogList)

	// 用户相关接口
	router.POST("/sysUser/add", sysUserController.AddSysUser)
	router.PUT("/sysUser/edit", sysUserController.EditSysUser)
	router.PUT("/sysUser/editPassword", sysUserController.EditUserPassword)
	router.PUT("/sysUser/resetPassword", sysUserController.ResetUserPassword)
	router.DELETE("/sysUser/del", sysUserController.DelSysUser)
	router.GET("/sysUser/detail", sysUserController.GetSysUserDetail)
	router.GET("/sysUser/list", sysUserController.GetSysUserList)
}

// 对外提供接口: 无需鉴权
func ExternalNotAuthApi(router *gin.RouterGroup) {
	//router.POST("/test", controller.AddUser)
}

// 对外提供接口: 需要鉴权
func ExternalAuthApi(router *gin.RouterGroup) {
}
