package service

import (
	"fmt"
	"time"
)

type TaskRegistry struct{}

// 注册任务函数 main函数中调用
// 所有定时任务在这里统一 注册
func (*TaskRegistry) RegisterTaskFunctions() {
	registerTask("ceshi1", func(params map[string]interface{}) error {
		fmt.Printf("[%s] 📧 发送邮件给开始: %s\n", time.Now().Format("15:04:05"), "123")
		time.Sleep(5 * time.Second)
		fmt.Printf("[%s] 📧 发送邮件结束: %s\n", time.Now().Format("15:04:05"), "456")
		return nil
	})
}
