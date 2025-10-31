package pojo

type SysDept struct {
	DeptId        string `gorm:"column:dept_id" json:"deptId"`                //主键
	DeptName      string `gorm:"column:dept_name" json:"deptName"`            //组织名称
	ParentId      string `gorm:"column:parent_id" json:"parentId"`            //上级组织
	DeptLevel     string `gorm:"column:dept_level" json:"deptLevel"`          //组织层级
	DeptStatus    string `gorm:"column:dept_status" json:"deptStatus"`        //状态(字典-sys_status)
	DeptSort      int    `gorm:"column:dept_sort" json:"deptSort"`            //排序
	CreatedUserId string `gorm:"column:created_user_id" json:"createdUserId"` //创建用户
	CreatedTime   string `gorm:"column:created_time" json:"createdTime"`      //创建时间
	UpdatedUserId string `gorm:"column:updated_user_id" json:"updatedUserId"` //更新用户
	UpdatedTime   string `gorm:"column:updated_time" json:"updatedTime"`      //更新时间
}

// AddSysDeptReq 新增组织入参
// @Description 新增组织
type AddSysDeptReq struct {
	DeptName   string `form:"deptName" json:"deptName" binding:"required"` // 组织名称
	ParentId   string `form:"parentId" json:"parentId"`                    // 上级组织
	DeptStatus string `form:"deptStatus" json:"deptStatus"`                // 状态(字典-sys_status)
	DeptSort   int    `form:"deptSort" json:"deptSort"`                    // 排序
}

// EditSysDeptReq 编辑组织入参
// @Description 编辑组织
type EditSysDeptReq struct {
	DeptId     string `form:"deptId" json:"deptId" binding:"required"`     // 主键
	DeptName   string `form:"deptName" json:"deptName" binding:"required"` // 组织名称
	ParentId   string `form:"parentId" json:"parentId"`                    // 上级组织
	DeptStatus string `form:"deptStatus" json:"deptStatus"`                // 状态(字典-sys_status)
	DeptSort   int    `form:"deptSort" json:"deptSort"`                    // 排序
}

// DelSysDeptReq 删除组织入参
// @Description 删除组织
type DelSysDeptReq struct {
	DeptIds string `form:"deptIds" json:"deptIds" binding:"required"` // 主键列表(逗号分隔)
}

// GetSysDeptDetailReq 查询组织详情入参
// @Description 查询组织详情
type GetSysDeptDetailReq struct {
	DeptId string `form:"deptId" json:"deptId" binding:"required"` // 主键
}

// GetSysDeptTreeReq 查询组织树入参
// @Description 查询组织树
type GetSysDeptTreeReq struct {
	DeptName   string `form:"deptName" json:"deptName"`     // 组织名称(模糊查询)
	DeptStatus string `form:"deptStatus" json:"deptStatus"` // 状态(字典-sys_status)
}

// SysDeptTreeNode 组织树节点结构体
// @Description 组织树节点
type SysDeptTreeNode struct {
	SysDept
	Children []SysDeptTreeNode `json:"children,omitempty"`
}
