package pojo

// User 菜单
// @Description 菜单
type SysMenu struct {
	MenuId           string `gorm:"column:menu_id" json:"menuId"`                     //主键
	ParentId         string `gorm:"column:parent_id" json:"parentId"`                 //上级菜单
	Path             string `gorm:"column:path" json:"path"`                          //路由路径
	AndroidPath      string `gorm:"column:android_path" json:"androidPath"`           //安卓路由路径
	IosPath          string `gorm:"column:ios_path" json:"iosPath"`                   //IOS路由路径
	Component        string `gorm:"column:component" json:"component"`                //组件路径
	AndroidComponent string `gorm:"column:android_component" json:"androidComponent"` //安卓组件路径
	IosComponent     string `gorm:"column:ios_component" json:"iosComponent"`         //IOS组件路径
	Redirect         string `gorm:"column:redirect" json:"redirect"`                  //重定向
	Type             string `gorm:"column:type" json:"type"`                          //类型(字典-sys_menu_type)
	Title            string `gorm:"column:title" json:"title"`                        //菜单标题
	SvgIcon          string `gorm:"column:svg_icon" json:"svgIcon"`                   //自定义图标(优先显示)
	Icon             string `gorm:"column:icon" json:"icon"`                          //菜单图标
	Hidden           bool   `gorm:"column:hidden" json:"hidden"`                      //是否隐藏(1-true 0-false)
	KeepAlive        bool   `gorm:"column:keep_alive" json:"keepAlive"`               //是否缓存(1-true 0-false)
	Breadcrumb       bool   `gorm:"column:breadcrumb" json:"breadcrumb"`              //面包屑(1-true 0-false)
	AlwaysShow       bool   `gorm:"column:always_show" json:"alwaysShow"`             //目录-总是显示(1-true 0-false)
	ShowInTabs       bool   `gorm:"column:showIn_tabs" json:"showInTabs"`             //菜单-页签显示(1-true 0-false)
	Affix            bool   `gorm:"column:affix" json:"affix"`                        //菜单-是否固定在标签页(1-true 0-false)
	Status           string `gorm:"column:status" json:"status"`                      //状态(字典-sys_status)
	ActiveMenu       string `gorm:"column:active_menu" json:"activeMenu"`             //激活的菜单项
	Permission       string `gorm:"column:permission" json:"permission"`              //按钮-权限标识
	ClientType       int    `gorm:"column:client_type" json:"clientType"`             //客户端类型(1-平台 2-APP)
	Sort             int    `gorm:"column:sort" json:"sort"`                          //菜单排序
	CreatedUserId    string `gorm:"column:created_user_id" json:"createdUserId"`      //创建用户
	CreatedTime      string `gorm:"column:created_time" json:"createdTime"`           //创建时间
	UpdatedUserId    string `gorm:"column:updated_user_id" json:"updatedUserId"`      //更新用户
	UpdatedTime      string `gorm:"column:updated_time" json:"updatedTime"`           //更新时间
}

type SysMenuResp struct {
	MenuId     string `gorm:"column:menu_id" json:"menuId"`         //主键
	ParentId   string `gorm:"column:parent_id" json:"parentId"`     //上级菜单
	Path       string `gorm:"column:path" json:"path"`              //路由路径
	Component  string `gorm:"column:component" json:"component"`    //组件路径
	Redirect   string `gorm:"column:redirect" json:"redirect"`      //重定向
	Type       string `gorm:"column:type" json:"type"`              //类型(字典-sys_menu_type)
	Title      string `gorm:"column:title" json:"title"`            //菜单标题
	SvgIcon    string `gorm:"column:svg_icon" json:"svgIcon"`       //自定义图标(优先显示)
	Icon       string `gorm:"column:icon" json:"icon"`              //菜单图标
	Hidden     bool   `gorm:"column:hidden" json:"hidden"`          //是否隐藏(1-true 0-false)
	KeepAlive  bool   `gorm:"column:keep_alive" json:"keepAlive"`   //是否缓存(1-true 0-false)
	Breadcrumb bool   `gorm:"column:breadcrumb" json:"breadcrumb"`  //面包屑(1-true 0-false)
	AlwaysShow bool   `gorm:"column:always_show" json:"alwaysShow"` //目录-总是显示(1-true 0-false)
	ShowInTabs bool   `gorm:"column:showIn_tabs" json:"showInTabs"` //菜单-页签显示(1-true 0-false)
	Affix      bool   `gorm:"column:affix" json:"affix"`            //菜单-是否固定在标签页(1-true 0-false)
	Status     string `gorm:"column:status" json:"status"`          //状态(字典-sys_status)
	ActiveMenu string `gorm:"column:active_menu" json:"activeMenu"` //激活的菜单项
	Permission string `gorm:"column:permission" json:"permission"`  //按钮-权限标识
	Sort       int    `gorm:"column:sort" json:"sort"`              //菜单排序
}

type SysMenuTreeNode struct {
	SysMenuResp
	Children []SysMenuTreeNode `json:"children,omitempty"`
}

type SysMenuAppResp struct {
	MenuId           string `gorm:"column:menu_id" json:"menuId"`                     //主键
	ParentId         string `gorm:"column:parent_id" json:"parentId"`                 //上级菜单
	AndroidPath      string `gorm:"column:android_path" json:"androidPath"`           //安卓路由路径
	IosPath          string `gorm:"column:ios_path" json:"iosPath"`                   //IOS路由路径
	AndroidComponent string `gorm:"column:android_component" json:"androidComponent"` //安卓组件路径
	IosComponent     string `gorm:"column:ios_component" json:"iosComponent"`         //IOS组件路径
	Type             string `gorm:"column:type" json:"type"`                          //类型(字典-sys_menu_type)
	Title            string `gorm:"column:title" json:"title"`                        //菜单标题
	Icon             string `gorm:"column:icon" json:"icon"`                          //菜单图标
	Status           string `gorm:"column:status" json:"status"`                      //状态(字典-sys_status)
	Permission       string `gorm:"column:permission" json:"permission"`              //按钮-权限标识
	Sort             int    `gorm:"column:sort" json:"sort"`                          //菜单排序
}

type SysMenuAppTreeNode struct {
	SysMenuAppResp
	Children []SysMenuAppTreeNode `json:"children,omitempty"`
}

// AddSysMenuDirectoryReq 新增目录入参
// @Description 新增目录
type AddSysMenuDirectoryReq struct {
	Type       string `form:"type" json:"type" binding:"required,oneof='1' '2' '3'"` // 菜单类型(字典-sys_menu_type)
	ParentId   string `form:"parentId" json:"parentId"`                              // 上级菜单
	Title      string `form:"title" json:"title" binding:"required"`                 // 菜单标题
	Component  string `form:"component" json:"component"`                            // 组件路径
	Sort       int    `form:"sort" json:"sort"`                                      // 菜单排序
	SvgIcon    string `form:"svgIcon" json:"svgIcon"`                                // 自定义图标
	Icon       string `form:"icon" json:"icon"`                                      // 菜单图标
	Path       string `form:"path" json:"path"`                                      // 路由路径
	Redirect   string `form:"redirect" json:"redirect"`                              // 重定向
	Status     string `form:"status" json:"status"`                                  // 状态(字典-sys_status)
	Hidden     bool   `form:"hidden" json:"hidden"`                                  // 是否隐藏
	KeepAlive  bool   `form:"keepAlive" json:"keepAlive"`                            // 是否缓存
	Affix      bool   `form:"affix" json:"affix"`                                    // 是否固定在标签页
	Breadcrumb bool   `form:"breadcrumb" json:"breadcrumb"`                          // 面包屑
	AlwaysShow bool   `form:"alwaysShow" json:"alwaysShow"`                          // 总是显示
	ClientType int    `json:"clientType" swaggerignore:"true"`                       // 客户端类型(1-平台 2-APP)
}

// AddSysMenuMenuReq 新增菜单入参
// @Description 新增菜单
type AddSysMenuMenuReq struct {
	Type       string `form:"type" json:"type" binding:"required,oneof='1' '2' '3'"` // 菜单类型(字典-sys_menu_type)
	ParentId   string `form:"parentId" json:"parentId"`                              // 上级菜单
	Title      string `form:"title" json:"title" binding:"required"`                 // 菜单标题
	Sort       int    `form:"sort" json:"sort"`                                      // 菜单排序
	SvgIcon    string `form:"svgIcon" json:"svgIcon"`                                // 自定义图标
	Icon       string `form:"icon" json:"icon"`                                      // 菜单图标
	Path       string `form:"path" json:"path"`                                      // 路由路径
	Redirect   string `form:"redirect" json:"redirect"`                              // 重定向
	Component  string `form:"component" json:"component"`                            // 组件路径
	Status     string `form:"status" json:"status"`                                  // 状态(字典-sys_status)
	Affix      bool   `form:"affix" json:"affix"`                                    // 是否固定在标签页
	Hidden     bool   `form:"hidden" json:"hidden"`                                  // 是否隐藏
	KeepAlive  bool   `form:"keepAlive" json:"keepAlive"`                            // 是否缓存
	Breadcrumb bool   `form:"breadcrumb" json:"breadcrumb"`                          // 面包屑
	ShowInTabs bool   `form:"showInTabs" json:"showInTabs"`                          // 页签显示
	ClientType int    `json:"clientType" swaggerignore:"true"`                       // 客户端类型(1-平台 2-APP)
}

// AddSysMenuAppDirectoryReq 新增APP目录入参
// @Description 新增菜单
type AddSysMenuAppDirectoryReq struct {
	Type             string `form:"type" json:"type" binding:"required,oneof='1' '2' '3'"` // 菜单类型(字典-sys_menu_type)
	ParentId         string `form:"parentId" json:"parentId"`                              // 上级菜单
	Title            string `form:"title" json:"title" binding:"required"`                 // 菜单标题
	Sort             int    `form:"sort" json:"sort"`                                      // 菜单排序
	Icon             string `form:"icon" json:"icon"`                                      // 菜单图标(上传)
	AndroidPath      string `form:"androidPath" json:"androidPath"`                        // 安卓路由路径
	IosPath          string `form:"iosPath" json:"iosPath"`                                // IOS路由路径
	AndroidComponent string `form:"androidComponent" json:"androidComponent"`              // 安卓组件路径
	IosComponent     string `form:"iosComponent" json:"iosComponent"`                      // IOS组件路径
	Status           string `form:"status" json:"status" binding:"required"`               // 状态(字典-sys_status)
	ClientType       int    `json:"clientType" swaggerignore:"true"`                       // 客户端类型(1-平台 2-APP)
}

// AddSysMenuMenuReq 新增APP菜单入参
// @Description 新增菜单
type AddSysMenuAppMenuReq struct {
	Type             string `form:"type" json:"type" binding:"required,oneof='1' '2' '3'"` // 菜单类型(字典-sys_menu_type)
	ParentId         string `form:"parentId" json:"parentId"`                              // 上级菜单
	Title            string `form:"title" json:"title" binding:"required"`                 // 菜单标题
	Sort             int    `form:"sort" json:"sort"`                                      // 菜单排序
	Icon             string `form:"icon" json:"icon" binding:"required"`                   // 菜单图标(上传)
	AndroidPath      string `form:"androidPath" json:"androidPath" binding:"required"`     // 安卓路由路径
	IosPath          string `form:"iosPath" json:"iosPath" binding:"required"`             // IOS路由路径
	AndroidComponent string `form:"androidComponent" json:"androidComponent"`              // 安卓组件路径
	IosComponent     string `form:"iosComponent" json:"iosComponent"`                      // IOS组件路径
	Status           string `form:"status" json:"status" binding:"required"`               // 状态(字典-sys_status)
	ClientType       int    `json:"clientType" swaggerignore:"true"`                       // 客户端类型(1-平台 2-APP)
}

// AddSysMenuButtonReq 新增按钮入参
// @Description 新增按钮
type AddSysMenuButtonReq struct {
	Type       string `form:"type" json:"type" binding:"required,oneof='1' '2' '3'"` // 菜单类型(字典-sys_menu_type)
	ParentId   string `form:"parentId" json:"parentId"`                              // 上级菜单
	Title      string `form:"title" json:"title" binding:"required"`                 // 菜单标题
	Sort       int    `form:"sort" json:"sort"`                                      // 菜单排序
	Permission string `form:"permission" json:"permission" binding:"required"`       // 权限标识
	Status     string `form:"status" json:"status"`                                  // 状态(字典-sys_status)
	ClientType int    `json:"clientType" swaggerignore:"true"`                       // 客户端类型(1-平台 2-APP)
}

type EditSysMenusReq struct {
	MenuIds  []string `form:"menuIds" json:"menuIds" binding:"required"` // 菜单id列表
	ParentId string   `form:"parentId" json:"parentId"`                  // 上级菜单
}

type EditSysMenuStatusReq struct {
	MenuId    string `form:"menuId" json:"menuId" binding:"required"` // 菜单id
	Type      string `form:"type" json:"type"`                        // 类型
	Status    string `form:"status" json:"status"`                    // 状态(字典-sys_status)
	Hidden    bool   `form:"hidden" json:"hidden"`                    // 是否隐藏
	KeepAlive bool   `form:"keepAlive" json:"keepAlive"`              // 是否缓存
}

// EditSysMenuDirectoryReq 修改目录入参
// @Description 修改目录
type EditSysMenuDirectoryReq struct {
	MenuId     string `form:"menuId" json:"menuId" binding:"required"`               // 主键
	Type       string `form:"type" json:"type" binding:"required,oneof='1' '2' '3'"` // 菜单类型(字典-sys_menu_type)
	ParentId   string `form:"parentId" json:"parentId"`                              // 上级菜单
	Title      string `form:"title" json:"title" binding:"required"`                 // 菜单标题
	Sort       int    `form:"sort" json:"sort"`                                      // 菜单排序
	SvgIcon    string `form:"svgIcon" json:"svgIcon"`                                // 自定义图标
	Icon       string `form:"icon" json:"icon"`                                      // 菜单图标
	Path       string `form:"path" json:"path"`                                      // 路由路径
	Redirect   string `form:"redirect" json:"redirect"`                              // 重定向
	Affix      bool   `form:"affix" json:"affix"`                                    // 是否固定在标签页
	Status     string `form:"status" json:"status"`                                  // 状态(字典-sys_status)
	Component  string `form:"component" json:"component"`                            // 组件路径
	Hidden     bool   `form:"hidden" json:"hidden"`                                  // 是否隐藏
	KeepAlive  bool   `form:"keepAlive" json:"keepAlive"`                            // 是否缓存
	Breadcrumb bool   `form:"breadcrumb" json:"breadcrumb"`                          // 面包屑
	AlwaysShow bool   `form:"alwaysShow" json:"alwaysShow"`                          // 总是显示
	ClientType int    `json:"clientType" swaggerignore:"true"`                       // 客户端类型(1-平台 2-APP)
}

// EditSysMenuMenuReq 修改菜单入参
// @Description 修改菜单
type EditSysMenuMenuReq struct {
	MenuId     string `form:"menuId" json:"menuId" binding:"required"`               // 主键
	Type       string `form:"type" json:"type" binding:"required,oneof='1' '2' '3'"` // 菜单类型(字典-sys_menu_type)
	ParentId   string `form:"parentId" json:"parentId"`                              // 上级菜单
	Title      string `form:"title" json:"title" binding:"required"`                 // 菜单标题
	Sort       int    `form:"sort" json:"sort"`                                      // 菜单排序
	SvgIcon    string `form:"svgIcon" json:"svgIcon"`                                // 自定义图标
	Icon       string `form:"icon" json:"icon"`                                      // 菜单图标
	Path       string `form:"path" json:"path"`                                      // 路由路径
	Redirect   string `form:"redirect" json:"redirect"`                              // 重定向
	Component  string `form:"component" json:"component"`                            // 组件路径
	Affix      bool   `form:"affix" json:"affix"`                                    // 是否固定在标签页
	Status     string `form:"status" json:"status"`                                  // 状态(字典-sys_status)
	Hidden     bool   `form:"hidden" json:"hidden"`                                  // 是否隐藏
	KeepAlive  bool   `form:"keepAlive" json:"keepAlive"`                            // 是否缓存
	Breadcrumb bool   `form:"breadcrumb" json:"breadcrumb"`                          // 面包屑
	ShowInTabs bool   `form:"showInTabs" json:"showInTabs"`                          // 页签显示
	ClientType int    `json:"clientType" swaggerignore:"true"`                       // 客户端类型(1-平台 2-APP)
}

// EditSysMenuAppDirectoryReq 修改APP目录入参
// @Description 修改APP目录
type EditSysMenuAppDirectoryReq struct {
	MenuId           string `form:"menuId" json:"menuId" binding:"required"`               // 主键
	Type             string `form:"type" json:"type" binding:"required,oneof='1' '2' '3'"` // 菜单类型(字典-sys_menu_type)
	ParentId         string `form:"parentId" json:"parentId"`                              // 上级菜单
	Title            string `form:"title" json:"title" binding:"required"`                 // 菜单标题
	Sort             int    `form:"sort" json:"sort"`                                      // 菜单排序
	Icon             string `form:"icon" json:"icon"`                                      // 菜单图标(上传)
	AndroidPath      string `form:"androidPath" json:"androidPath"`                        // 安卓路由路径
	IosPath          string `form:"iosPath" json:"iosPath"`                                // IOS路由路径
	AndroidComponent string `form:"androidComponent" json:"androidComponent"`              // 安卓组件路径
	IosComponent     string `form:"iosComponent" json:"iosComponent"`                      // IOS组件路径
	Status           string `form:"status" json:"status" binding:"required"`               // 状态(字典-sys_status)
	ClientType       int    `json:"clientType" swaggerignore:"true"`                       // 客户端类型(1-平台 2-APP)
}

// EditSysMenuAppMenuReq 修改APP菜单入参
// @Description 修改APP菜单
type EditSysMenuAppMenuReq struct {
	MenuId           string `form:"menuId" json:"menuId" binding:"required"`               // 主键
	Type             string `form:"type" json:"type" binding:"required,oneof='1' '2' '3'"` // 菜单类型(字典-sys_menu_type)
	ParentId         string `form:"parentId" json:"parentId"`                              // 上级菜单
	Title            string `form:"title" json:"title" binding:"required"`                 // 菜单标题
	Sort             int    `form:"sort" json:"sort"`                                      // 菜单排序
	Icon             string `form:"icon" json:"icon" binding:"required"`                   // 菜单图标(上传)
	AndroidPath      string `form:"androidPath" json:"androidPath" binding:"required"`     // 安卓路由路径
	IosPath          string `form:"iosPath" json:"iosPath" binding:"required"`             // IOS路由路径
	AndroidComponent string `form:"androidComponent" json:"androidComponent"`              // 安卓组件路径
	IosComponent     string `form:"iosComponent" json:"iosComponent"`                      // IOS组件路径
	Status           string `form:"status" json:"status" binding:"required"`               // 状态(字典-sys_status)
	ClientType       int    `json:"clientType" swaggerignore:"true"`                       // 客户端类型(1-平台 2-APP)
}

// EditSysMenuButtonReq 修改按钮入参
// @Description 修改按钮
type EditSysMenuButtonReq struct {
	MenuId     string `form:"menuId" json:"menuId" binding:"required"`               // 主键
	Type       string `form:"type" json:"type" binding:"required,oneof='1' '2' '3'"` // 菜单类型(字典-sys_menu_type)
	ParentId   string `form:"parentId" json:"parentId"`                              // 上级菜单
	Title      string `form:"title" json:"title" binding:"required"`                 // 菜单标题
	Sort       int    `form:"sort" json:"sort"`                                      // 菜单排序
	Permission string `form:"permission" json:"permission" binding:"required"`       // 权限标识
	Status     string `form:"status" json:"status"`                                  // 状态(字典-sys_status)
	ClientType int    `json:"clientType" swaggerignore:"true"`                       // 客户端类型(1-平台 2-APP)
}

// DelSysMenuReq 删除菜单入参
// @Description 删除菜单
type DelSysMenuReq struct {
	MenuIds string `form:"menuIds" json:"menuIds" binding:"required"` // 主键
}

// GetSysMenuByIdReq 菜单详情入参
// @Description 菜单详情
type GetSysMenuByIdReq struct {
	MenuId string `form:"menuId" json:"menuId"` // 主键
}

// GetMenuTreeListReq 菜单查询入参
// @Description 菜单查询
type GetMenuTreeListReq struct {
	MenuType string   `json:"menuType"` // 菜单类型 all查询(1目录2菜单3按钮) menu查询(1目录2菜单)
	MenuIds  []string `json:"menuIds"`  // 菜单ID列表
}
