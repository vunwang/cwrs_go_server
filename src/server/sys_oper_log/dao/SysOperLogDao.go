package dao

import (
	"cwrs_go_server/src/cwrs_core/cwrs_gorm"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_oper_log/pojo"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

var sysOperLogLog = cwrs_zap_logger.ZapLogger

const tableSysOperLog = "sys_oper_log"
const tableSelSysOperLog = "sys_oper_log sol"

type SysOperLogDao struct{}

// DelSysOperLog 删除操作日志
func (*SysOperLogDao) DelSysOperLog(operIds []string) error {
	if err := cwrs_gorm.GormDb.Table(tableSysOperLog).
		Where("oper_id IN (?)", operIds).Delete(&pojo.SysOperLog{}).Error; err != nil {
		sysOperLogLog.Error("DelSysOperLog Error", zap.Error(err))
		return err
	}
	return nil
}

// AddTomorrowPartition 新增日分区
func (*SysOperLogDao) AddTomorrowPartition(partitionName, dayAfterTomorrow string) {
	sql := fmt.Sprintf("ALTER TABLE %s REORGANIZE PARTITION p_future INTO (PARTITION %s VALUES LESS THAN (TO_DAYS('%s')),PARTITION p_future VALUES LESS THAN MAXVALUE)",
		tableSysOperLog,
		partitionName,
		dayAfterTomorrow,
	)

	sysOperLogLog.Info("➕ 准备创建明日分区", zap.String("partitionName", partitionName))
	err := cwrs_gorm.GormDb.Exec(sql).Error
	if err != nil {
		// 如果已存在，忽略错误（幂等性）
		if strings.Contains(err.Error(), "1517") {
			sysOperLogLog.Warn("ℹ️ 分区已存在，跳过创建", zap.String("partitionName", partitionName))
			return
		}
		sysOperLogLog.Error("AddTomorrowPartition 创建分区失败", zap.Error(err))
		return
	}
	sysOperLogLog.Info("✅ 分区创建成功", zap.String("partitionName", partitionName))
	return
}

// DropOldDailyPartitions 删除N天前的旧分区（如保留最近30天）
// 参数：keepDays int 保留的天数
func (*SysOperLogDao) DropOldDailyPartitions(cutoffTime time.Time) {
	cutoffYear, cutoffMonth, cutoffDay := cutoffTime.Date()
	cutoffPartitionName := fmt.Sprintf("p%04d%02d%02d", cutoffYear, cutoffMonth, cutoffDay)

	sysOperLogLog.Info("📅 保留分区从: "+cutoffPartitionName+" 开始", zap.String("cutoffPartitionName", cutoffPartitionName))

	rows, err := cwrs_gorm.GormDb.Raw(
		"SELECT PARTITION_NAME FROM information_schema.PARTITIONS WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = ? AND PARTITION_NAME LIKE 'p2%' AND PARTITION_NAME < ? ORDER BY PARTITION_NAME ASC",
		tableSysOperLog, cutoffPartitionName).Rows()
	if err != nil {
		sysOperLogLog.Error("DropOldDailyPartitions 查询分区列表失败", zap.Error(err))
		return
	}
	defer rows.Close()

	var oldPartitions []string
	for rows.Next() {
		var partitionName string
		if err = rows.Scan(&partitionName); err != nil {
			sysOperLogLog.Error("DropOldDailyPartitions 分区列表数据读取失败", zap.Error(err))
			return
		}
		if partitionName != "p_future" {
			oldPartitions = append(oldPartitions, partitionName)
		}
	}

	if len(oldPartitions) == 0 {
		sysOperLogLog.Info("✅ 没有需要删除的旧分区")
		return
	}

	dropSQL := fmt.Sprintf("ALTER TABLE %s DROP PARTITION %s", tableSysOperLog, strings.Join(oldPartitions, ", "))

	sysOperLogLog.Info("🗑️ 准备删除 "+strconv.Itoa(len(oldPartitions))+" 个旧分区", zap.Strings("oldPartitions", oldPartitions))
	err = cwrs_gorm.GormDb.Exec(dropSQL).Error
	if err != nil {
		sysOperLogLog.Error("DropOldDailyPartitions 删除分区失败", zap.Error(err))
		return
	}
	sysOperLogLog.Info("✅ 成功删除旧分区", zap.Strings("oldPartitions", oldPartitions))
	return
}

// GetSysOperLogById 查询操作日志详情
func (*SysOperLogDao) GetSysOperLogById(operId string) (*pojo.SysOperLogResp, error) {
	var item pojo.SysOperLogResp
	if err := cwrs_gorm.GormDb.Table(tableSelSysOperLog).
		Select("sol.oper_id", "sol.oper_user_id", "sol.method", "sol.path", "sol.ip", "sol.status", "sol.req_body", "sol.res_body", "sol.latency", "sol.oper_time", "su.user_name as oper_name").
		Joins("LEFT JOIN sys_user su ON sol.oper_user_id = su.user_id").
		Where("sol.oper_id = ?", operId).First(&item).Error; err != nil {
		sysOperLogLog.Error("GetSysOperLogById Error", zap.Error(err))
		return nil, err
	}
	return &item, nil
}

// GetSysOperLogList 分页查询操作日志列表
func (*SysOperLogDao) GetSysOperLogList(req *pojo.GetSysOperLogListReq) ([]pojo.SysOperLogListResp, int64, error) {
	var list []pojo.SysOperLogListResp
	var total int64
	db := cwrs_gorm.GormDb.Table(tableSelSysOperLog).
		Select("sol.oper_id", "sol.oper_user_id", "sol.method", "sol.path", "sol.ip", "sol.status", "sol.latency", "sol.oper_time", "su.user_name as oper_name").
		Joins("LEFT JOIN sys_user su ON sol.oper_user_id = su.user_id")
	if req.OperName != "" {
		db = db.Where("su.user_name LIKE ?", "%"+req.OperName+"%")
	}
	if req.StartTime != "" && req.EndTime != "" {
		db = db.Where("sol.oper_time BETWEEN ? AND ?", req.StartTime, req.EndTime)
	}
	db.Count(&total)
	if req.PageNum > 0 && req.PageSize > 0 {
		offset := cwrs_utils.CalcOffset(req.PageNum, req.PageSize)
		db = db.Offset(offset).Limit(req.PageSize)
	}
	if err := db.Order("sol.oper_time desc").Find(&list).Error; err != nil {
		sysOperLogLog.Error("GetSysOperLogList Error", zap.Error(err))
		return nil, 0, err
	}
	return list, total, nil
}
