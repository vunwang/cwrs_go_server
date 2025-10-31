package cwrs_gorm

import (
	"cwrs_go_server/src/cwrs_core/cwrs_viper"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var GormDb *gorm.DB
var err error

func init() {
	InitGORM()
}

func InitGORM() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	// MySQL 配置信息
	username := cwrs_viper.GlobalViper.GetString("mysql.user-name") // 账号
	password := cwrs_viper.GlobalViper.GetString("mysql.password")  // 密码
	host := cwrs_viper.GlobalViper.GetString("mysql.path")          // 地址
	port := cwrs_viper.GlobalViper.GetString("mysql.port")          // 端口
	DBname := cwrs_viper.GlobalViper.GetString("mysql.db-name")     // 数据库名称
	timeout := "10s"                                                // 连接超时，10秒
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, DBname, timeout)
	// Open 连接
	GormDb, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %w", err))
	}

	// ====== ⭐ 关键：配置数据库连接池 ======
	sqlDB, err := GormDb.DB()
	if err != nil {
		panic(fmt.Errorf("failed to get sql.DB: %w", err))
	}

	// 设置连接池参数（根据你的服务器和 MySQL 配置调整）
	sqlDB.SetMaxOpenConns(50)                  // 最大打开连接数（建议 20~100）
	sqlDB.SetMaxIdleConns(10)                  // 最大空闲连接数（建议 5~20）
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // 连接最大存活时间

	fmt.Println("Gorm Initialize OK !")
}
