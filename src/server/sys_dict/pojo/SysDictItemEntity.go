package pojo

// User 数据字典
// @Description 数据字典项
type SysDictItem struct {
	DictItemId    string `gorm:"column:dict_item_id" json:"dictItemId"`       //主键
	DictCode      string `gorm:"column:dict_code" json:"dictCode"`            //字典编码
	ItemName      string `gorm:"column:item_name" json:"itemName"`            //名称
	ItemValue     string `gorm:"column:item_value" json:"itemValue"`          //字典值
	ItemColor     string `gorm:"column:item_color" json:"itemColor"`          //展示颜色
	ItemStatus    string `gorm:"column:item_status" json:"itemStatus"`        //状态(字典-sys_status)
	ItemSelect    string `gorm:"column:item_select" json:"itemSelect"`        //下拉菜单是否显示(1是 0否)
	ItemSort      int    `gorm:"column:item_sort" json:"itemSort"`            //排序
	CreatedUserId string `gorm:"column:created_user_id" json:"createdUserId"` //创建用户;所有表中都要有该字段，数据权限需要使用
	CreatedTime   string `gorm:"column:created_time" json:"createdTime"`      //创建时间
	UpdatedUserId string `gorm:"column:updated_user_id" json:"updatedUserId"` //更新用户
	UpdatedTime   string `gorm:"column:updated_time" json:"updatedTime"`      //更新时间
}

// AddSysDictItemReq 新增数据字典项入参
// @Description 新增数据字典项
type AddSysDictItemReq struct {
	DictCode   string `form:"dictCode" json:"dictCode" binding:"required"`      // 字典编码(必填)
	ItemName   string `form:"itemName" json:"itemName" binding:"required"`      // 项名称
	ItemValue  string `form:"itemValue" json:"itemValue" binding:"required"`    // 项值
	ItemColor  string `form:"itemColor" json:"itemColor"`                       // 展示颜色
	ItemSort   int    `form:"itemSort" json:"itemSort"`                         // 排序
	ItemStatus string `form:"itemStatus" json:"itemStatus" binding:"oneof=0 1"` // 状态(字典-sys_status)
	ItemSelect string `form:"itemSelect" json:"itemSelect" binding:"oneof=0 1"` // 下拉菜单是否显示(1是 0否)
	Desc       string `form:"desc" json:"desc"`                                 // 描述
}

// EditSysDictItemReq 修改数据字典项入参
// @Description 修改数据字典项
type EditSysDictItemReq struct {
	DictItemId string `form:"dictItemId" json:"dictItemId" binding:"required"` // 主键(必填)
	DictCode   string `form:"dictCode" json:"dictCode"`                        // 字典编码(必填)
	ItemName   string `form:"itemName" json:"itemName"`                        // 项名称
	ItemValue  string `form:"itemValue" json:"itemValue"`                      // 项值
	ItemColor  string `form:"itemColor" json:"itemColor"`                      // 展示颜色
	ItemSort   int    `form:"itemSort" json:"itemSort"`                        // 排序
	ItemStatus string `form:"itemStatus" json:"itemStatus"`                    // 状态(字典-sys_status)
	ItemSelect string `form:"itemSelect" json:"itemSelect"`                    // 下拉菜单是否显示(1是 0否)
	Desc       string `form:"desc" json:"desc"`                                // 描述
}

// DelSysDictItemReq 查询数据字典项详情入参
// @Description 查询数据字典项详情入参
type GetSysDictItemDetailReq struct {
	DictItemId string `form:"dictItemId" json:"dictItemId" binding:"required"` // 主键(必填)
}

// DelSysDictItemReq 删除数据字典项入参
// @Description 删除数据字典项
type DelSysDictItemReq struct {
	DictItemIds string `form:"dictItemIds" json:"dictItemIds" binding:"required"` // 主键(必填)
}

// GetSysDictItemListReq 分页查询数据字典项入参
// @Description 分页查询数据字典项
type GetSysDictItemListReq struct {
	PageNum  int    `form:"pageNum" json:"pageNum"`                      // 页码
	PageSize int    `form:"pageSize" json:"pageSize"`                    // 每页显示条数
	DictCode string `form:"dictCode" json:"dictCode" binding:"required"` // 字典编码(必填)
	ItemName string `form:"itemName" json:"itemName"`                    // 项名称(模糊查)
}
