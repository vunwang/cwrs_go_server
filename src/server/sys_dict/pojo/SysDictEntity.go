package pojo

// User 数据字典
// @Description 数据字典
type SysDict struct {
	DictId        string `gorm:"column:dict_id" json:"dictId"`                //主键
	DictName      string `gorm:"column:dict_name" json:"dictName"`            //名称
	DictCode      string `gorm:"column:dict_code" json:"dictCode"`            //编码
	DictStatus    string `gorm:"column:dict_status" json:"dictStatus"`        //状态(字典-sys_status)
	DictSort      int    `gorm:"column:dict_sort" json:"dictSort"`            //排序
	Desc          string `gorm:"column:desc" json:"desc"`                     //描述
	CreatedUserId string `gorm:"column:created_user_id" json:"createdUserId"` //创建用户;所有表中都要有该字段，数据权限需要使用
	CreatedTime   string `gorm:"column:created_time" json:"createdTime"`      //创建时间
	UpdatedUserId string `gorm:"column:updated_user_id" json:"updatedUserId"` //更新用户
	UpdatedTime   string `gorm:"column:updated_time" json:"updatedTime"`      //更新时间
}

// AddSysDictReq 新增数据字典入参
// @Description 新增数据字典
type AddSysDictReq struct {
	DictName   string `form:"dictName" json:"dictName" binding:"required"`      // 名称
	DictCode   string `form:"dictCode" json:"dictCode" binding:"required"`      // 编码
	DictStatus string `form:"dictStatus" json:"dictStatus" binding:"oneof=0 1"` // 状态(字典-sys_status)
	DictSort   int    `form:"dictSort" json:"dictSort"`                         // 排序
	Desc       string `form:"desc" json:"desc"`                                 // 描述
}

// EditSysDictReq 修改数据字典入参
// @Description 修改数据字典
type EditSysDictReq struct {
	DictId     string `form:"dictId" json:"dictId" binding:"required"`          // 主键
	DictName   string `form:"dictName" json:"dictName" binding:"required"`      // 名称
	DictCode   string `form:"dictCode" json:"dictCode" binding:"required"`      // 编码
	DictStatus string `form:"dictStatus" json:"dictStatus" binding:"oneof=0 1"` // 状态(字典-sys_status)
	DictSort   int    `form:"dictSort" json:"dictSort"`                         // 排序
	Desc       string `form:"desc" json:"desc"`                                 // 描述
}

// DelSysDictReq 删除数据字典入参
// @Description 删除数据字典
type DelSysDictReq struct {
	DictIds string `form:"dictIds" json:"dictIds" binding:"required"` // 主键
}

// GetSysDictDetailReq 查询数据字典详情入参
// @Description 查询数据字典详情
type GetSysDictDetailReq struct {
	DictId string `form:"dictId" json:"dictId" binding:"required"` // 主键
}

// GetSysDictListReq 分页查询数据字典入参
// @Description 分页查询数据字典
type GetSysDictListReq struct {
	PageNum    int    `form:"pageNum" json:"pageNum"`       // 页码
	PageSize   int    `form:"pageSize" json:"pageSize"`     // 每页显示调数
	DictName   string `form:"dictName" json:"dictName"`     // 名称(模糊查)
	DictStatus string `form:"dictStatus" json:"dictStatus"` // 状态(字典-sys_status)
}
