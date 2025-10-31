package service

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_oper_log/dao"
	"cwrs_go_server/src/server/sys_oper_log/pojo"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var sysOperLogDaoImpl = dao.SysOperLogDao{}

type SysOperLogService struct{}

// AddTomorrowPartition 新增指定天数的日分区
// days 1 新增明天的日分区
// days 2 新增后天的日分区
// days 3 新增大后天的日分区
func (*SysOperLogService) AddTomorrowPartition(days int) {
	tomorrow := cwrs_utils.ChangeDateTime(time.Now(), 0, 0, days, 0, 0, 0) // 当前时间+1天 明天
	year, month, day := tomorrow.Date()

	// 分区名：pYYYYMMDD
	partitionName := fmt.Sprintf("p%04d%02d%02d", year, month, day)
	// 明天0点（作为分区下限）
	tomorrowStart := cwrs_utils.SetDateYear_Month_Day(year, month, day)
	// 后天0点（作为分区上限）
	dayAfterTomorrow := tomorrowStart.AddDate(0, 0, 1).Format("2006-01-02")

	sysOperLogDaoImpl.AddTomorrowPartition(partitionName, dayAfterTomorrow)
}

// DropOldDailyPartitions 删除N天前的旧分区（如保留最近60天）
// 参数：keepDays int 保留的天数
func (*SysOperLogService) DropOldDailyPartitions(keepDays int) {
	cutoffTime := cwrs_utils.ChangeDateTime(time.Now(), 0, 0, -keepDays, 0, 0, 0)
	sysOperLogDaoImpl.DropOldDailyPartitions(cutoffTime)
}

// DelSysOperLog 删除操作日志
func (*SysOperLogService) DelSysOperLog(c *gin.Context, req *pojo.DelSysOperLogReq) {
	operIds := strings.Split(req.OperIds, ",")
	if err := sysOperLogDaoImpl.DelSysOperLog(operIds); err != nil {
		cwrs_res.Waring(c, err, "删除操作日志失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// GetSysOperLogDetail 查询操作日志详情
func (*SysOperLogService) GetSysOperLogDetail(c *gin.Context, req *pojo.GetSysOperLogDetailReq) {
	item, err := sysOperLogDaoImpl.GetSysOperLogById(req.OperId)
	if err != nil {
		cwrs_res.Waring(c, err, "查询操作日志详情失败")
		return
	}
	cwrs_res.SuccessData(c, "操作成功", item)
}

// GetSysOperLogList 分页查询操作日志列表
func (*SysOperLogService) GetSysOperLogList(c *gin.Context, req *pojo.GetSysOperLogListReq) {
	list, total, err := sysOperLogDaoImpl.GetSysOperLogList(req)
	if err != nil {
		cwrs_res.Waring(c, err, "查询操作日志列表失败")
		return
	}
	cwrs_res.SuccessDataList(c, "操作成功", list, total)
}
