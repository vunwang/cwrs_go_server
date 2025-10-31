package pojo

type SysPost struct {
	PostId        string `gorm:"column:post_id" json:"postId"`                //主键
	PostName      string `gorm:"column:post_name" json:"postName"`            //岗位名称
	PostCode      string `gorm:"column:post_code" json:"postCode"`            //岗位编码
	PostStatus    string `gorm:"column:post_status" json:"postStatus"`        //状态(字典-sys_status)
	PostSort      int    `gorm:"column:post_sort" json:"postSort"`            //排序
	DeptId        string `gorm:"column:dept_id" json:"deptId"`                //所属组织
	Desc          string `gorm:"column:desc" json:"desc"`                     //描述
	CreatedUserId string `gorm:"column:created_user_id" json:"createdUserId"` //创建用户
	CreatedTime   string `gorm:"column:created_time" json:"createdTime"`      //创建时间
	UpdatedUserId string `gorm:"column:updated_user_id" json:"updatedUserId"` //更新用户
	UpdatedTime   string `gorm:"column:updated_time" json:"updatedTime"`      //更新时间
}

type SysPostResp struct {
	PostId          string `gorm:"column:post_id" json:"postId"`                    //主键
	PostName        string `gorm:"column:post_name" json:"postName"`                //岗位名称
	PostCode        string `gorm:"column:post_code" json:"postCode"`                //岗位编码
	PostStatus      string `gorm:"column:post_status" json:"postStatus"`            //状态(字典-sys_status)
	PostSort        int    `gorm:"column:post_sort" json:"postSort"`                //排序
	DeptId          string `gorm:"column:dept_id" json:"deptId"`                    //所属组织
	DeptName        string `gorm:"column:dept_name" json:"deptName"`                //所属组织名称
	Desc            string `gorm:"column:desc" json:"desc"`                         //描述
	CreatedUserId   string `gorm:"column:created_user_id" json:"createdUserId"`     //创建用户
	CreatedUserName string `gorm:"column:created_user_name" json:"createdUserName"` //创建用户名称
	CreatedTime     string `gorm:"column:created_time" json:"createdTime"`          //创建时间
	UpdatedUserId   string `gorm:"column:updated_user_id" json:"updatedUserId"`     //更新用户
	UpdatedTime     string `gorm:"column:updated_time" json:"updatedTime"`          //更新时间
}

// AddSysPostReq 新增岗位入参
// @Description 新增岗位
type AddSysPostReq struct {
	PostName   string `form:"postName" json:"postName" binding:"required"` // 岗位名称
	PostCode   string `form:"postCode" json:"postCode" binding:"required"` // 岗位编码
	PostStatus string `form:"postStatus" json:"postStatus"`                // 状态(字典-sys_status)
	PostSort   int    `form:"postSort" json:"postSort"`                    // 排序
	DeptId     string `form:"deptId" json:"deptId"`                        // 所属组织
	Desc       string `form:"desc" json:"desc"`                            // 描述
}

// EditSysPostReq 编辑岗位入参
// @Description 编辑岗位
type EditSysPostReq struct {
	PostId     string `form:"postId" json:"postId" binding:"required"`     // 主键
	PostName   string `form:"postName" json:"postName" binding:"required"` // 岗位名称
	PostCode   string `form:"postCode" json:"postCode" binding:"required"` // 岗位编码
	PostStatus string `form:"postStatus" json:"postStatus"`                // 状态(字典-sys_status)
	PostSort   int    `form:"postSort" json:"postSort"`                    // 排序
	DeptId     string `form:"deptId" json:"deptId"`                        // 所属组织
	Desc       string `form:"desc" json:"desc"`                            // 描述
}

// DelSysPostReq 删除岗位入参
// @Description 删除岗位
type DelSysPostReq struct {
	PostIds string `form:"postIds" json:"postIds" binding:"required"` // 逗号分隔
}

// GetSysPostDetailReq 查询详情入参
// @Description 查询岗位详情
type GetSysPostDetailReq struct {
	PostId string `form:"postId" json:"postId" binding:"required"`
}

// GetSysPostListReq 分页查询入参
// @Description 分页查询岗位
type GetSysPostListReq struct {
	PageNum    int    `form:"pageNum" json:"pageNum"`       // 页码
	PageSize   int    `form:"pageSize" json:"pageSize"`     // 每页显示条数
	DeptId     string `form:"deptId" json:"deptId"`         // 所属组织
	PostName   string `form:"postName" json:"postName"`     // 岗位名称(模糊)
	PostCode   string `form:"postCode" json:"postCode"`     // 岗位编码(模糊)
	PostStatus string `form:"postStatus" json:"postStatus"` // 状态(字典-sys_status)
}
