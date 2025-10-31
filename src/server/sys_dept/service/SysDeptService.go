package service

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_dept/dao"
	"cwrs_go_server/src/server/sys_dept/pojo"
	paramDao "cwrs_go_server/src/server/sys_param/dao"
	postDao "cwrs_go_server/src/server/sys_post/dao"
	roleDao "cwrs_go_server/src/server/sys_role/dao"
	userDao "cwrs_go_server/src/server/sys_user/dao"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

var sysDeptDaoImpl = dao.SysDeptDao{}
var sysUserDaoImpl = userDao.SysUserDao{}
var sysRoleDaoImpl = roleDao.SysRoleDao{}
var sysPostDaoImpl = postDao.SysPostDao{}
var sysParamDaoImpl = paramDao.SysParamDao{}

type SysDeptService struct{}

// AddSysDept 新增组织
func (*SysDeptService) AddSysDept(c *gin.Context, req *pojo.AddSysDeptReq) {
	var entity pojo.SysDept
	cwrs_utils.CopyStruct(req, &entity, "json")
	entity.DeptId = cwrs_utils.CreateUuid()
	entity.CreatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.CreatedTime = cwrs_utils.GetNowDateTime()

	// 计算组织层级
	entity.DeptLevel = calculateDeptLevel(req.ParentId, entity.DeptId)

	if err := sysDeptDaoImpl.AddSysDept(&entity); err != nil {
		cwrs_res.Waring(c, err, "新增组织失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// EditSysDept 修改组织
func (*SysDeptService) EditSysDept(c *gin.Context, req *pojo.EditSysDeptReq) {
	var entity pojo.SysDept
	cwrs_utils.CopyStruct(req, &entity, "json")
	entity.UpdatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.UpdatedTime = cwrs_utils.GetNowDateTime()

	// 计算组织层级
	entity.DeptLevel = calculateDeptLevel(req.ParentId, entity.DeptId)

	if err := sysDeptDaoImpl.EditSysDept(&entity); err != nil {
		cwrs_res.Waring(c, err, "修改组织失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// DelSysDept 删除组织
func (*SysDeptService) DelSysDept(c *gin.Context, req *pojo.DelSysDeptReq) {
	deptIds := strings.Split(req.DeptIds, ",")
	// 检查是否有子组织
	count, _ := sysDeptDaoImpl.GetDeptCountByParentIds(deptIds)
	userCount, _ := sysUserDaoImpl.GetUserCountByDeptIds(deptIds)
	roleCount, _ := sysRoleDaoImpl.GetRoleCountByDeptIds(deptIds)
	postCount, _ := sysPostDaoImpl.GetPostCountByDeptIds(deptIds)
	paramCount, _ := sysParamDaoImpl.GetParamCountByDeptIds(deptIds)
	if count > 0 || userCount > 0 || roleCount > 0 || postCount > 0 || paramCount > 0 {
		cwrs_res.Waring(c, nil, "该组织已被使用，无法删除")
		return
	}

	if err := sysDeptDaoImpl.DelSysDept(deptIds); err != nil {
		cwrs_res.Waring(c, err, "删除组织失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// GetSysDeptDetail 查询组织详情
func (*SysDeptService) GetSysDeptDetail(c *gin.Context, req *pojo.GetSysDeptDetailReq) {
	dept, err := sysDeptDaoImpl.GetSysDeptById(req.DeptId)
	if err != nil {
		cwrs_res.Waring(c, err, "查询组织详情失败")
		return
	}
	cwrs_res.SuccessData(c, "操作成功", dept)
}

// GetSysDeptTree 查询组织树
func (*SysDeptService) GetSysDeptTree(c *gin.Context, req *pojo.GetSysDeptTreeReq) {
	list, err := sysDeptDaoImpl.GetSysDeptList(c, req)
	if err != nil {
		cwrs_res.Waring(c, err, "查询组织列表失败")
		return
	}

	// 构建组织树（自动识别根节点）
	tree := buildDeptTreeAuto(list)
	cwrs_res.SuccessData(c, "操作成功", tree)
}

// buildDeptTreeAuto 自动识别根节点并构建树
func buildDeptTreeAuto(depts []pojo.SysDept) []pojo.SysDeptTreeNode {
	// 1. 构建 deptId -> dept 的映射，方便快速查找
	deptMap := make(map[string]pojo.SysDept)
	for _, dept := range depts {
		deptMap[dept.DeptId] = dept
	}

	// 2. 找出所有根节点：ParentId 为空、"0"，或 ParentId 不在 deptMap 中
	var roots []pojo.SysDept
	for _, dept := range depts {
		parentId := dept.ParentId
		// 判断是否为根节点
		if parentId == "" || parentId == "0" || deptMap[parentId].DeptId == "" {
			roots = append(roots, dept)
		}
	}

	// 对根节点按 Sort 排序
	sort.Slice(roots, func(i, j int) bool {
		return roots[i].DeptSort < roots[j].DeptSort
	})

	// 3. 为每个根节点递归构建子树
	var tree []pojo.SysDeptTreeNode
	for _, root := range roots {
		node := pojo.SysDeptTreeNode{SysDept: root}
		node.Children = buildChildren(deptMap, root.DeptId)
		tree = append(tree, node)
	}

	return tree
}

// buildChildren 递归构建子节点
func buildChildren(deptMap map[string]pojo.SysDept, parentId string) []pojo.SysDeptTreeNode {
	var children []pojo.SysDeptTreeNode
	for _, dept := range deptMap {
		if dept.ParentId == parentId {
			child := pojo.SysDeptTreeNode{SysDept: dept}
			child.Children = buildChildren(deptMap, dept.DeptId)
			children = append(children, child)
		}
	}
	// 对当前层级的子节点按 Sort 排序
	sort.Slice(children, func(i, j int) bool {
		return children[i].SysDept.DeptSort < children[j].SysDept.DeptSort
	})
	return children
}

// calculateDeptLevel 计算组织层级
func calculateDeptLevel(parentId, deptId string) string {
	if parentId == "0" || parentId == "" {
		return "0," + deptId
	}

	// 查询父级组织信息
	parentDept, err := sysDeptDaoImpl.GetSysDeptById(parentId)
	if err != nil {
		// 如果查询失败，返回默认层级
		return "0," + deptId
	}

	// 构建层级：父级层级 + 当前父级ID
	return parentDept.DeptLevel + "," + deptId
}
