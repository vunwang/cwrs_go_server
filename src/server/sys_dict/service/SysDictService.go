package service

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_dict/dao"
	"cwrs_go_server/src/server/sys_dict/pojo"
	"strings"

	"github.com/gin-gonic/gin"
)

var sysDictDaoImpl = dao.SysDictDao{}
var sysDictItemDaoImpl = dao.SysDictItemDao{}

type SysDictService struct{}
type SysDictItemService struct{}

// AddSysDict 新增数据字典
func (*SysDictService) AddSysDict(c *gin.Context, req *pojo.AddSysDictReq) {
	var entity pojo.SysDict
	cwrs_utils.CopyStruct(req, &entity, "json")
	entity.DictId = cwrs_utils.CreateUuid()
	entity.CreatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.CreatedTime = cwrs_utils.GetNowDateTime()
	if err := sysDictDaoImpl.AddSysDict(&entity); err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "sys_dict.only_dict_name") {
			cwrs_res.Forbidden(c, err, "字典名称已存在")
		} else if strings.Contains(errStr, "sys_dict.only_dict_code") {
			cwrs_res.Forbidden(c, err, "字典编码已存在")
		} else {
			cwrs_res.Waring(c, err, "操作失败")
		}
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// EditSysDict 修改数据字典
func (*SysDictService) EditSysDict(c *gin.Context, req *pojo.EditSysDictReq) {
	var entity pojo.SysDict
	cwrs_utils.CopyStruct(req, &entity, "json")
	entity.UpdatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.UpdatedTime = cwrs_utils.GetNowDateTime()
	if err := sysDictDaoImpl.EditSysDict(&entity); err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "sys_dict.only_dict_name") {
			cwrs_res.Forbidden(c, err, "字典名称已存在")
		} else if strings.Contains(errStr, "sys_dict.only_dict_code") {
			cwrs_res.Forbidden(c, err, "字典编码已存在")
		} else {
			cwrs_res.Waring(c, err, "操作失败")
		}
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// DelSysDict 删除数据字典
func (*SysDictService) DelSysDict(c *gin.Context, req *pojo.DelSysDictReq) {
	dictIds := strings.Split(req.DictIds, ",")
	//查询改字典下是否有字典项
	if dictItemCount, err := sysDictItemDaoImpl.GetDictItemCountByDictIds(dictIds); err != nil {
		cwrs_res.Waring(c, err, "操作失败")
		return
	} else if dictItemCount > 0 {
		cwrs_res.Forbidden(c, nil, "请先删除字典项，再删除字典")
		return
	}

	if err := sysDictDaoImpl.DelSysDict(dictIds); err != nil {
		cwrs_res.Waring(c, err, "删除数据字典失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// GetSysDictDetail 查询数据字典详情
func (*SysDictService) GetSysDictDetail(c *gin.Context, req *pojo.GetSysDictDetailReq) {
	dict, err := sysDictDaoImpl.GetSysDictById(req.DictId)
	if err != nil {
		cwrs_res.Waring(c, err, "查询数据字典详情失败")
		return
	}
	cwrs_res.SuccessData(c, "操作成功", dict)
}

// GetSysDictList 分页查询数据字典（支持名称模糊查）
func (*SysDictService) GetSysDictList(c *gin.Context, req *pojo.GetSysDictListReq) {
	list, total, err := sysDictDaoImpl.GetSysDictList(req)
	if err != nil {
		cwrs_res.Waring(c, err, "查询数据字典列表失败")
		return
	}
	cwrs_res.SuccessDataList(c, "操作成功", list, total)
}

// GetAllDictMap 查询所有字典及其字典项，返回map[dictCode][]SysDictItem
func (*SysDictService) GetAllDictMap() (map[string][]pojo.SysDictItem, error) {
	codes, err := sysDictDaoImpl.GetAllDictCodes()
	if err != nil {
		return nil, err
	}
	items, err := sysDictItemDaoImpl.GetAllDictItems()
	if err != nil {
		return nil, err
	}
	dictMap := make(map[string][]pojo.SysDictItem)
	for _, code := range codes {
		dictMap[code] = []pojo.SysDictItem{}
	}
	for _, item := range items {
		if _, ok := dictMap[item.DictCode]; ok {
			dictMap[item.DictCode] = append(dictMap[item.DictCode], item)
		}
	}
	return dictMap, nil
}

// AddSysDictItem 新增数据字典项
func (*SysDictItemService) AddSysDictItem(c *gin.Context, req *pojo.AddSysDictItemReq) {
	var entity pojo.SysDictItem
	cwrs_utils.CopyStruct(req, &entity, "json")
	entity.DictItemId = cwrs_utils.CreateUuid()
	entity.CreatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.CreatedTime = cwrs_utils.GetNowDateTime()
	if err := sysDictItemDaoImpl.AddSysDictItem(&entity); err != nil {
		cwrs_res.Waring(c, err, "新增数据字典项失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// EditSysDictItem 修改数据字典项
func (*SysDictItemService) EditSysDictItem(c *gin.Context, req *pojo.EditSysDictItemReq) {
	var entity pojo.SysDictItem
	cwrs_utils.CopyStruct(req, &entity, "json")
	entity.UpdatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.UpdatedTime = cwrs_utils.GetNowDateTime()
	if err := sysDictItemDaoImpl.EditSysDictItem(&entity); err != nil {
		cwrs_res.Waring(c, err, "修改数据字典项失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// DelSysDictItem 删除数据字典项
func (*SysDictItemService) DelSysDictItem(c *gin.Context, req *pojo.DelSysDictItemReq) {
	dictItemIds := strings.Split(req.DictItemIds, ",")
	if err := sysDictItemDaoImpl.DelSysDictItem(dictItemIds); err != nil {
		cwrs_res.Waring(c, err, "删除数据字典项失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// GetSysDictItemList 分页查询数据字典项（字典编码必填+名称模糊查）
func (*SysDictItemService) GetSysDictItemList(c *gin.Context, req *pojo.GetSysDictItemListReq) {
	list, total, err := sysDictItemDaoImpl.GetSysDictItemList(req)
	if err != nil {
		cwrs_res.Waring(c, err, "查询数据字典项列表失败")
		return
	}
	cwrs_res.SuccessDataList(c, "操作成功", list, total)
}

// GetSysDictItemDetail 查询数据字典项详情
func (*SysDictItemService) GetSysDictItemDetail(c *gin.Context, req *pojo.GetSysDictItemDetailReq) {
	item, err := sysDictItemDaoImpl.GetSysDictItemById(req.DictItemId)
	if err != nil {
		cwrs_res.Waring(c, err, "查询数据字典项详情失败")
		return
	}
	cwrs_res.SuccessData(c, "操作成功", item)
}
