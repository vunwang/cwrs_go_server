package service

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_task/pojo"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"strings"
	"sync"
)

// 任务函数类型
type TaskFunc func(params map[string]interface{}) error

// 全局注册表
var taskRegistry = make(map[string]TaskFunc)

// 🌟 全局变量（必须定义！）
var (
	scheduler = cron.New(cron.WithSeconds())  // 调度器
	taskMap   = make(map[string]cron.EntryID) // 任务ID -> Cron EntryID 映射
	mu        sync.RWMutex                    // 读写锁，保证并发安全
)

// 注册函数
func registerTask(name string, fn TaskFunc) {
	taskRegistry[name] = fn
}

// 获取函数
func getTaskFunc(name string) (TaskFunc, bool) {
	fn, ok := taskRegistry[name]
	return fn, ok
}

// 启动调度器 main函数中调用
func (*TaskRegistry) Start() {
	scheduler.Start()
}

// 将开启的定时任务添加到调度器中 main函数中调用
// 这样就不用每次重启项目再去手动开启已启用的任务了
func (*TaskRegistry) AddTaskToSchedulerByStatus() {
	// 查询所有任务
	taskList, err := sysTaskDaoImpl.GetSysTaskListByStatus("1")
	if err != nil {
		cwrs_zap_logger.Error("AddTaskToScheduler Error", zap.Error(err))
		return
	}
	//添加任务到调度器
	count := addTaskToScheduler(taskList)
	cwrs_zap_logger.Info(fmt.Sprintf("AddTaskToScheduler 共添加了 %d 个任务到调度器", count))
}

// 启动任务接口
func (*TaskRegistry) StartTask(c *gin.Context, req *pojo.StartAndStopTaskReq) {
	taskIds := strings.Split(req.TaskIds, ",")
	// 根据任务id查询任务数据
	taskList, err := sysTaskDaoImpl.GetSysTaskListByIds(taskIds)
	if err != nil {
		cwrs_res.Waring(c, err, "查询任务数据失败")
		return
	}
	//添加任务到调度器
	count := addTaskToScheduler(taskList)
	if err = sysTaskDaoImpl.EditSysTaskStatus("1", cwrs_utils.GetNowDateTime(), taskIds); err != nil {
		cwrs_res.Waring(c, err, "修改任务状态失败")
		return
	}
	cwrs_res.Success(c, fmt.Sprintf("%d个任务已启动，%d个任务启动失败", count, len(taskList)-count))
}

// 添加任务到调度器
func addTaskToScheduler(taskList []pojo.SysTaskResp) int {
	count := 0
	for _, task := range taskList {
		// 并发安全：检查是否已存在
		mu.RLock()
		if _, exists := taskMap[task.TaskId]; exists {
			mu.RUnlock()
			cwrs_zap_logger.Warn("任务已存在，跳过添加", zap.String("taskId", task.TaskId), zap.String("taskName", task.TaskName))
			continue
		}
		mu.RUnlock()

		// 检查函数是否存在
		fn, ok := getTaskFunc(task.FuncName)
		if !ok {
			cwrs_zap_logger.Error("无效的函数名", zap.String("funcName", task.FuncName), zap.String("taskId", task.TaskId))
			continue
		}

		// 创建包装函数
		wrapper := func() {
			if errCall := fn(task.TaskParams); errCall != nil { // 注意：这里改名为 errCall，避免闭包覆盖外层 err
				cwrs_zap_logger.Error("任务执行失败",
					zap.String("taskId", task.TaskId),
					zap.String("taskName", task.TaskName),
					zap.Error(errCall))
			}
		}

		// 添加到调度器
		entryID, errRes := scheduler.AddFunc(task.CronExpr, wrapper)
		if errRes != nil {
			cwrs_zap_logger.Error("任务启动失败",
				zap.String("taskId", task.TaskId),
				zap.String("taskName", task.TaskName),
				zap.Error(errRes))
			continue
		}

		// 并发安全：保存到全局映射
		mu.Lock()
		// 再次检查（防止极端并发下重复添加）
		if _, exists := taskMap[task.TaskId]; exists {
			mu.Unlock()
			scheduler.Remove(entryID) // 安全起见，移除刚添加的任务
			cwrs_zap_logger.Warn("任务在保存时被并发添加，已回滚", zap.String("taskId", task.TaskId))
			continue
		}
		taskMap[task.TaskId] = entryID
		mu.Unlock()

		count++
	}
	return count
}

// 停止任务接口
func (*TaskRegistry) StopTask(c *gin.Context, req *pojo.StartAndStopTaskReq) {
	mu.Lock()
	defer mu.Unlock()

	count := 0
	taskIds := strings.Split(req.TaskIds, ",")
	if err := sysTaskDaoImpl.EditSysTaskStatus("2", cwrs_utils.GetNowDateTime(), taskIds); err != nil {
		cwrs_res.Waring(c, err, "修改任务状态失败")
		return
	}
	for _, taskId := range taskIds {
		if entryID, exists := taskMap[taskId]; exists {
			scheduler.Remove(entryID)
			delete(taskMap, taskId)
			count++
		}
	}
	cwrs_res.Success(c, fmt.Sprintf("%d个任务已停止，%d个任务未找到", count, len(taskIds)-count))
}
