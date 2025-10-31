package cwrs_viper

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var GlobalViper *viper.Viper

func init() {
	InitViper()
}

func InitViper() {
	GlobalViper = viper.New()

	// 设置主配置文件为 application.yaml
	GlobalViper.SetConfigName("application")
	GlobalViper.SetConfigType("yaml")

	// 添加配置路径（当前目录）
	configDir := "."
	sysType := runtime.GOOS
	if sysType == "linux" {
		configDir = "./"
		fmt.Println("Linux system")
	} else if sysType == "windows" {
		configDir = "./"
		fmt.Println("Windows system")
	}
	GlobalViper.AddConfigPath(configDir)

	// 1. 读取主配置 application.yaml
	if err := GlobalViper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("❌ 读取 application.yaml 失败: %w", err))
	}

	// 2. 获取 active-profile
	profile := GlobalViper.GetString("app.active-profile")
	if profile == "" {
		// fallback: 环境变量 > 默认 dev
		if env := os.Getenv("APP_PROFILE"); env != "" {
			profile = env
		} else {
			profile = "dev"
		}
	}

	fmt.Printf("📌 当前激活环境: %s\n", profile)

	// 3. 加载并合并 application-{profile}.yaml
	if profile != "" && profile != "default" {
		profileFile := fmt.Sprintf("application-%s.yaml", profile)
		profilePath := filepath.Join(configDir, profileFile)

		if _, err := os.Stat(profilePath); err == nil {
			// 读取 profile 配置
			profileViper := viper.New()
			profileViper.SetConfigFile(profilePath)
			if err := profileViper.ReadInConfig(); err != nil {
				log.Fatalf("❌ 读取 %s 失败: %v", profileFile, err)
			}

			// 合并：profile 配置优先级更高
			GlobalViper.MergeConfigMap(profileViper.AllSettings())
			fmt.Printf("✅ 已合并环境配置: %s\n", profileFile)
		} else {
			fmt.Printf("⚠️  环境配置文件 %s 不存在，仅使用 application.yaml\n", profileFile)
		}
	}

	fmt.Println("✅ Viper 初始化成功！")
}
