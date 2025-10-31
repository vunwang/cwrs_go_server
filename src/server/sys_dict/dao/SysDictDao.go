package dao

import (
	"cwrs_go_server/src/cwrs_core/cwrs_gorm"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_dict/pojo"

	"go.uber.org/zap"
)

var dictLog = cwrs_zap_logger.ZapLogger

const (
	tableSysDict     = "sys_dict"
	tableSysDictItem = "sys_dict_item"
)

type SysDictDao struct{}
type SysDictItemDao struct{}

// AddSysDict 添加数据字典
func (*SysDictDao) AddSysDict(dict *pojo.SysDict) error {
	fields := make([]string, 0)
	if dict.DictName != "" {
		fields = append(fields, "dict_name")
	}
	if dict.DictCode != "" {
		fields = append(fields, "dict_code")
	}
	if dict.DictStatus != "" {
		fields = append(fields, "dict_status")
	}
	if dict.DictSort != 0 {
		fields = append(fields, "dict_sort")
	}
	if dict.Desc != "" {
		fields = append(fields, "desc")
	}
	fields = append(fields, "dict_id", "created_user_id", "created_time")
	if err := cwrs_gorm.GormDb.Table(tableSysDict).Select(fields).
		Create(&dict).Error; err != nil {
		dictLog.Error("AddSysDict Error", zap.Error(err))
		return err
	}
	return nil
}

// EditSysDict 修改数据字典
func (*SysDictDao) EditSysDict(dict *pojo.SysDict) error {
	fields := make([]string, 0)
	if dict.DictName != "" {
		fields = append(fields, "dict_name")
	}
	if dict.DictCode != "" {
		fields = append(fields, "dict_code")
	}
	if dict.DictStatus != "" {
		fields = append(fields, "dict_status")
	}
	if dict.DictSort != 0 {
		fields = append(fields, "dict_sort")
	}
	fields = append(fields, "desc", "updated_user_id", "updated_time")
	if err := cwrs_gorm.GormDb.Table(tableSysDict).Select(fields).
		Where("dict_id = ?", dict.DictId).Updates(&dict).Error; err != nil {
		dictLog.Error("EditSysDict Error", zap.Error(err))
		return err
	}
	return nil
}

// DelSysDict 删除数据字典
func (*SysDictDao) DelSysDict(dictIds []string) error {
	if err := cwrs_gorm.GormDb.Table(tableSysDict).
		Where("dict_id IN (?)", dictIds).Delete(&pojo.SysDict{}).Error; err != nil {
		dictLog.Error("DelSysDict Error", zap.Error(err))
		return err
	}
	return nil
}

// GetDictItemCountByDictIds 查询字典下是否有字典项
func (*SysDictItemDao) GetDictItemCountByDictIds(dictIds []string) (int64, error) {
	var count int64
	if err := cwrs_gorm.GormDb.Table(tableSysDictItem).
		Joins("INNER JOIN sys_dict ON sys_dict.dict_code = sys_dict_item.dict_code").
		Select("count(*)").
		Where("sys_dict.dict_id IN (?)", dictIds).
		Count(&count).Error; err != nil {
		dictLog.Error("GetDictItemCountByDictIds Error", zap.Error(err))
		return 0, err
	}
	return count, nil
}

// GetSysDictById 查询数据字典详情
func (*SysDictDao) GetSysDictById(dictId string) (*pojo.SysDict, error) {
	var dict pojo.SysDict
	if err := cwrs_gorm.GormDb.Table(tableSysDict).
		Select("dict_id", "dict_name", "dict_code", "dict_status", "dict_sort", "desc", "created_user_id", "created_time", "updated_user_id", "updated_time").
		Where("dict_id = ?", dictId).First(&dict).Error; err != nil {
		dictLog.Error("GetSysDictById Error", zap.Error(err))
		return nil, err
	}
	return &dict, nil
}

// GetSysDictList 分页+名称模糊查
func (*SysDictDao) GetSysDictList(req *pojo.GetSysDictListReq) ([]pojo.SysDict, int64, error) {
	var list []pojo.SysDict
	var total int64
	db := cwrs_gorm.GormDb.Table(tableSysDict).
		Select("dict_id", "dict_name", "dict_code", "dict_status", "dict_sort", "desc", "created_user_id", "created_time", "updated_user_id", "updated_time")
	if req.DictName != "" {
		db = db.Where("dict_name LIKE ?", "%"+req.DictName+"%").Or("dict_code LIKE ?", "%"+req.DictName+"%")
	}
	if req.DictStatus != "" {
		db = db.Where("dict_status = ?", req.DictStatus)
	}
	db.Count(&total)
	if req.PageNum > 0 && req.PageSize > 0 {
		offset := cwrs_utils.CalcOffset(req.PageNum, req.PageSize)
		db = db.Offset(offset).Limit(req.PageSize)
	}
	if err := db.Order("dict_sort").Find(&list).Error; err != nil {
		dictLog.Error("GetSysDictList Error", zap.Error(err))
		return nil, 0, err
	}
	return list, total, nil
}

// AddSysDictItem 添加数据字典项
func (*SysDictItemDao) AddSysDictItem(item *pojo.SysDictItem) error {
	fields := make([]string, 0)
	if item.DictCode != "" {
		fields = append(fields, "dict_code")
	}
	if item.ItemName != "" {
		fields = append(fields, "item_name")
	}
	if item.ItemValue != "" {
		fields = append(fields, "item_value")
	}
	if item.ItemColor != "" {
		fields = append(fields, "item_color")
	}
	if item.ItemStatus != "" {
		fields = append(fields, "item_status")
	}
	if item.ItemSelect != "" {
		fields = append(fields, "item_select")
	}
	if item.ItemSort != 0 {
		fields = append(fields, "item_sort")
	}
	fields = append(fields, "dict_item_id", "created_user_id", "created_time")
	if err := cwrs_gorm.GormDb.Table(tableSysDictItem).Select(fields).
		Create(&item).Error; err != nil {
		dictLog.Error("AddSysDictItem Error", zap.Error(err))
		return err
	}
	return nil
}

// EditSysDictItem 修改数据字典项
func (*SysDictItemDao) EditSysDictItem(item *pojo.SysDictItem) error {
	fields := make([]string, 0)
	if item.DictCode != "" {
		fields = append(fields, "dict_code")
	}
	if item.ItemName != "" {
		fields = append(fields, "item_name")
	}
	if item.ItemValue != "" {
		fields = append(fields, "item_value")
	}
	if item.ItemColor != "" {
		fields = append(fields, "item_color")
	}
	if item.ItemStatus != "" {
		fields = append(fields, "item_status")
	}
	if item.ItemSelect != "" {
		fields = append(fields, "item_select")
	}
	if item.ItemSort != 0 {
		fields = append(fields, "item_sort")
	}
	fields = append(fields, "updated_user_id", "updated_time")
	if err := cwrs_gorm.GormDb.Table(tableSysDictItem).Select(fields).
		Where("dict_item_id = ?", item.DictItemId).
		Updates(&item).Error; err != nil {
		dictLog.Error("EditSysDictItem Error", zap.Error(err))
		return err
	}
	return nil
}

// DelSysDictItem 删除数据字典项
func (*SysDictItemDao) DelSysDictItem(itemIds []string) error {
	if err := cwrs_gorm.GormDb.Table(tableSysDictItem).
		Where("dict_item_id IN (?)", itemIds).
		Delete(&pojo.SysDictItem{}).Error; err != nil {
		dictLog.Error("DelSysDictItem Error", zap.Error(err))
		return err
	}
	return nil
}

// GetSysDictItemList 分页+字典编码必填+名称模糊查
func (*SysDictItemDao) GetSysDictItemList(req *pojo.GetSysDictItemListReq) ([]pojo.SysDictItem, int64, error) {
	var list []pojo.SysDictItem
	var total int64
	db := cwrs_gorm.GormDb.Table(tableSysDictItem).
		Select("dict_item_id", "dict_code", "item_name", "item_value", "item_color", "item_status", "item_select", "item_sort", "created_user_id", "created_time", "updated_user_id", "updated_time")
	db = db.Where("dict_code = ?", req.DictCode)
	if req.ItemName != "" {
		db = db.Where("item_name LIKE ?", "%"+req.ItemName+"%")
	}
	db.Count(&total)
	if req.PageNum > 0 && req.PageSize > 0 {
		offset := cwrs_utils.CalcOffset(req.PageNum, req.PageSize)
		db = db.Offset(offset).Limit(req.PageSize)
	}
	if err := db.Order("item_sort asc").Find(&list).Error; err != nil {
		dictLog.Error("GetSysDictItemList Error", zap.Error(err))
		return nil, 0, err
	}
	return list, total, nil
}

func (*SysDictItemDao) GetSysDictItemById(itemId string) (*pojo.SysDictItem, error) {
	var item pojo.SysDictItem
	if err := cwrs_gorm.GormDb.Table(tableSysDictItem).
		Select("dict_item_id", "dict_code", "item_name", "item_value", "item_color", "item_status", "item_select", "item_sort", "created_user_id", "created_time", "updated_user_id", "updated_time").
		Where("dict_item_id = ?", itemId).
		First(&item).Error; err != nil {
		dictLog.Error("GetSysDictItemById Error", zap.Error(err))
		return nil, err
	}
	return &item, nil
}

// GetAllDictCodes 查询所有字典编码
func (*SysDictDao) GetAllDictCodes() ([]string, error) {
	var codes []string
	if err := cwrs_gorm.GormDb.Table(tableSysDict).
		Select("dict_code").
		Order("dict_sort asc").
		Find(&codes).Error; err != nil {
		dictLog.Error("GetAllDictCodes Error", zap.Error(err))
		return nil, err
	}
	return codes, nil
}

// GetAllDictItems 查询所有字典项
func (*SysDictItemDao) GetAllDictItems() ([]pojo.SysDictItem, error) {
	var items []pojo.SysDictItem
	if err := cwrs_gorm.GormDb.Table(tableSysDictItem).
		Select("dict_item_id", "dict_code", "item_name", "item_value", "item_color", "item_status", "item_select", "item_sort").
		Order("dict_code asc, item_sort asc").
		Find(&items).Error; err != nil {
		dictLog.Error("GetAllDictItems Error", zap.Error(err))
		return nil, err
	}
	return items, nil
}
