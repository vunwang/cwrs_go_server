package cwrs_zap_logger

import (
	"cwrs_go_server/src/cwrs_core/cwrs_viper"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
	"time"
)

var ZapLogger *zap.Logger
var SugaredLogger *zap.SugaredLogger

func init() {
	InitLogger()
}

func InitLogger() {
	//获取编码器
	encoder := getEncoder()

	//日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	//info文件WriteSyncer
	infoFileWriteSyncer := getInfoWriterSyncer()
	//error文件WriteSyncer
	errorFileWriteSyncer := getErrorWriterSyncer()

	//只记录,不输出到控制台
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer), lowPriority)
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority)

	//将infocore 和 errcore 加入core切片
	var coreArr []zapcore.Core
	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)

	//生成Logger
	ZapLogger = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()) //zap.AddCaller() 显示文件名 和 行号
	SugaredLogger = ZapLogger.Sugar()
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func levelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var level string
	switch l {
	case zapcore.DebugLevel:
		level = "[DEBUG]"
	case zapcore.InfoLevel:
		level = "[INFO]"
	case zapcore.WarnLevel:
		level = "[WARN]"
	case zapcore.ErrorLevel:
		level = "[ERROR]"
	case zapcore.DPanicLevel:
		level = "[DPANIC]"
	case zapcore.PanicLevel:
		level = "[PANIC]"
	case zapcore.FatalLevel:
		level = "[FATAL]"
	default:
		level = fmt.Sprintf("[LEVEL(%d)]", l)
	}
	enc.AppendString(level)
}

func shortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%s]", caller.TrimmedPath()))
}

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder, //zapcore.CapitalColorLevelEncoder(这个控制台输出带颜色)  zapcore.CapitalLevelEncoder(后面这俩日志文件可以解析出颜色),   或者  levelEncoder（自定义）,
		EncodeTime:     timeEncoder,                 //指定时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   shortCallerEncoder, //zapcore.ShortCallerEncoder,
	}
}

// cwrs_core 三个参数之  Encoder 获取编码器
func getEncoder() zapcore.Encoder {
	//自定义编码配置,下方NewJSONEncoder输出如下的日志格式
	//{"L":"[INFO]","T":"2022-09-16 14:24:59.552","C":"[prototest/main.go:113]","M":"name = xiaoli, age = 18"}
	//return zapcore.NewJSONEncoder(NewEncoderConfig())

	//下方NewConsoleEncoder输出如下的日志格式
	//2022-09-16 14:26:02.933 [INFO]  [prototest/main.go:113] name = xiaoli, age = 18
	return zapcore.NewConsoleEncoder(NewEncoderConfig())
}

// cwrs_core 三个参数之  日志输出路径
func getInfoWriterSyncer() zapcore.WriteSyncer {
	sysType := runtime.GOOS
	var path = "./logs/info.log"
	if sysType == "linux" {
		// LINUX系统
		fmt.Println("Linux system")
		path = "./logs/info.log"
	}

	//引入第三方库 Lumberjack 加入日志切割功能
	infoLumberIO := &lumberjack.Logger{
		Filename:   path,                                             //日志文件位置
		MaxSize:    cwrs_viper.GlobalViper.GetInt("log.max-size"),    //每个日志文件保存的大小 单位:M
		MaxBackups: cwrs_viper.GlobalViper.GetInt("log.max-backups"), // 保留旧文件最大个数
		MaxAge:     cwrs_viper.GlobalViper.GetInt("log.max-age"),     // 保留旧文件最大天数
		Compress:   cwrs_viper.GlobalViper.GetBool("log.compress"),   //是否压缩/归档旧文件
	}
	return zapcore.AddSync(infoLumberIO)
}

func getErrorWriterSyncer() zapcore.WriteSyncer {
	sysType := runtime.GOOS
	var path = "./logs/error.log"
	if sysType == "linux" {
		// LINUX系统
		fmt.Println("Linux system")
		path = "./logs/error.log"
	}

	//引入第三方库 Lumberjack 加入日志切割功能
	lumberWriteSyncer := &lumberjack.Logger{
		Filename:   path,                                             // 日志文件位置
		MaxSize:    cwrs_viper.GlobalViper.GetInt("log.max-size"),    // 每个日志文件保存的大小 单位:M
		MaxBackups: cwrs_viper.GlobalViper.GetInt("log.max-backups"), // 保留旧文件最大个数
		MaxAge:     cwrs_viper.GlobalViper.GetInt("log.max-age"),     // 保留旧文件最大天数
		Compress:   cwrs_viper.GlobalViper.GetBool("log.compress"),   // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberWriteSyncer)
}

func Debugf(format string, v ...interface{}) {
	ZapLogger.Sugar().Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	ZapLogger.Sugar().Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	ZapLogger.Sugar().Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	ZapLogger.Sugar().Errorf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	ZapLogger.Sugar().Panicf(format, v...)
}

// logs.Debug(...) 再封装
func Debug(format string, fileds ...zapcore.Field) {
	ZapLogger.Debug(format, fileds...)
}

func Info(format string, fileds ...zapcore.Field) {
	ZapLogger.Info(format, fileds...)
}

func Warn(format string, fileds ...zapcore.Field) {
	ZapLogger.Warn(format, fileds...)
}

func Error(format string, fileds ...zapcore.Field) {
	ZapLogger.Error(format, fileds...)
}

func Panic(format string, fileds ...zapcore.Field) {
	ZapLogger.Panic(format, fileds...)
}
