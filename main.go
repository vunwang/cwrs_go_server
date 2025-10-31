package main

import (
	"cwrs_go_server/src/cwrs_core/cwrs_gin"
	"cwrs_go_server/src/cwrs_core/cwrs_viper"
	"cwrs_go_server/src/cwrs_utils"
	logService "cwrs_go_server/src/server/sys_oper_log/service"
	taskService "cwrs_go_server/src/server/sys_task/service"
	"fmt"
)

// @title Cwrs Manager API
// @version 1.0
// @description Cwrs管理系统API文档
// @termsOfService 本系统为开源项目，可用于学习和研究，如用于商业用途造成的后果与本组织无关。

// @contact.name vn_wang
// @contact.QQ 1820816162
// @contact.email vn_wang@163.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9090
// @BasePath
var taskRegistryImpl = taskService.TaskRegistry{}
var sysOperLogImpl = logService.SysOperLogService{}

func main() {
	//初始化Gin框架
	gin := cwrs_gin.InitGin()
	//新增 今天和明天的日志分区 确保分区存在
	sysOperLogImpl.AddTomorrowPartition(0)
	sysOperLogImpl.AddTomorrowPartition(1)
	//预注册定时任务
	taskRegistryImpl.RegisterTaskFunctions()
	//将开启的定时任务添加到调度器中
	taskRegistryImpl.AddTaskToSchedulerByStatus()
	//运行定时任务调度器
	taskRegistryImpl.Start()

	fmt.Printf("----------------------------------------------------------------------------\n")
	fmt.Printf("\tcwrs_go_server is running! Access URLs:\n")
	fmt.Printf("\tLocal-URL: http://localhost:%s\n", cwrs_viper.GlobalViper.GetString("gin.port"))
	fmt.Printf("\tAPI-DOC文档: http://localhost:%s/swagger/index.html\n", cwrs_viper.GlobalViper.GetString("gin.port"))
	fmt.Printf("\tExternal-URL: http://%s:%s\n", cwrs_utils.GetLocalIP(), cwrs_viper.GlobalViper.GetString("gin.port"))
	fmt.Printf("----------------------------------------------------------------------------\n")
	err := gin.Run(fmt.Sprintf(":%s", cwrs_viper.GlobalViper.GetString("gin.port")))
	if err != nil {
		fmt.Printf("服务启动失败\n")
		fmt.Printf("gin start error ***** Error ***** : %#v", err)
		panic(err)
	}
}
