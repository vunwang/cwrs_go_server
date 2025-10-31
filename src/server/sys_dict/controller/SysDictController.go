package controller

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/server/sys_dict/pojo"
	"cwrs_go_server/src/server/sys_dict/service"

	"github.com/gin-gonic/gin"
)

var sysDictServiceImpl = service.SysDictService{}
var sysDictItemServiceImpl = service.SysDictItemService{}

// @Tags 数据字典【平台】
// @Summary 新增数据字典
// @Description 新增数据字典
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.AddSysDictReq true "新增数据字典参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDict/add [post]
func AddSysDict(c *gin.Context) {
	var req pojo.AddSysDictReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysDictServiceImpl.AddSysDict(c, &req)
}

// @Tags 数据字典【平台】
// @Summary 修改数据字典
// @Description 修改数据字典
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysDictReq true "修改数据字典参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDict/edit [put]
func EditSysDict(c *gin.Context) {
	var req pojo.EditSysDictReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysDictServiceImpl.EditSysDict(c, &req)
}

// @Tags 数据字典【平台】
// @Summary 删除数据字典(支持批量)
// @Description 删除数据字典
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.DelSysDictReq true "删除数据字典参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDict/del [delete]
func DelSysDict(c *gin.Context) {
	var req pojo.DelSysDictReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysDictServiceImpl.DelSysDict(c, &req)
}

// @Tags 数据字典【平台】
// @Summary 查询数据字典详情
// @Description 查询数据字典详情
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysDictDetailReq true "查询详情参数"
// @Success 200 {object} cwrs_res.ResSuccessData{data=pojo.SysDict} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDict/detail [get]
func GetSysDictDetail(c *gin.Context) {
	var req pojo.GetSysDictDetailReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysDictServiceImpl.GetSysDictDetail(c, &req)
}

// @Tags 数据字典【平台】
// @Summary 分页查询数据字典
// @Description 分页查询数据字典（支持名称模糊查）
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysDictListReq true "分页查询参数"
// @Success 200 {object} cwrs_res.ResSuccessDataList{data=[]pojo.SysDict} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDict/list [get]
func GetSysDictList(c *gin.Context) {
	var req pojo.GetSysDictListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysDictServiceImpl.GetSysDictList(c, &req)
}

// @Tags 数据字典【平台】
// @Summary 查询所有字典及其字典项
// @Description 查询所有字典及其字典项，返回map[dictCode][]SysDictItem
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Success 200 {object} map[string][]pojo.SysDictItem "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDict/allMap [get]
func GetAllDictMap(c *gin.Context) {
	dictMap, err := sysDictServiceImpl.GetAllDictMap()
	if err != nil {
		cwrs_res.Waring(c, err, "查询所有字典及字典项失败")
		return
	}
	cwrs_res.SuccessData(c, "操作成功", dictMap)
}

// @Tags 数据字典项【平台】
// @Summary 新增数据字典项
// @Description 新增数据字典项
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.AddSysDictItemReq true "新增数据字典项参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDictItem/add [post]
func AddSysDictItem(c *gin.Context) {
	var req pojo.AddSysDictItemReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysDictItemServiceImpl.AddSysDictItem(c, &req)
}

// @Tags 数据字典项【平台】
// @Summary 修改数据字典项
// @Description 修改数据字典项
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysDictItemReq true "修改数据字典项参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDictItem/edit [put]
func EditSysDictItem(c *gin.Context) {
	var req pojo.EditSysDictItemReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysDictItemServiceImpl.EditSysDictItem(c, &req)
}

// @Tags 数据字典项【平台】
// @Summary 删除数据字典项(支持批量)
// @Description 删除数据字典项
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.DelSysDictItemReq true "删除数据字典项参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDictItem/del [delete]
func DelSysDictItem(c *gin.Context) {
	var req pojo.DelSysDictItemReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysDictItemServiceImpl.DelSysDictItem(c, &req)
}

// @Tags 数据字典项【平台】
// @Summary 分页查询数据字典项
// @Description 分页查询数据字典项（字典编码必填+名称模糊查）
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysDictItemListReq true "分页查询参数"
// @Success 200 {object} cwrs_res.ResSuccessDataList{data=[]pojo.SysDictItem} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDictItem/list [get]
func GetSysDictItemList(c *gin.Context) {
	var req pojo.GetSysDictItemListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysDictItemServiceImpl.GetSysDictItemList(c, &req)
}

// @Tags 数据字典项【平台】
// @Summary 查询数据字典项详情
// @Description 查询数据字典项详情
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysDictItemDetailReq true "查询详情参数"
// @Success 200 {object} cwrs_res.ResSuccessData{data=pojo.SysDictItem} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDictItem/detail [get]
func GetSysDictItemDetail(c *gin.Context) {
	var req pojo.GetSysDictItemDetailReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysDictItemServiceImpl.GetSysDictItemDetail(c, &req)
}
