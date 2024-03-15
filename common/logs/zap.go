package logs

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"site/conf"
	"strings"
	"time"
)

// Logger 全局日志配置
var (
	Logger      *zap.Logger
	SugarLogger *zap.SugaredLogger
)

// InitLogger 初始化日志
func InitLogger() {
	// 获取编码器
	encoder := getEncoder()

	// 对日志级别进行判断、分类
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // error 级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // info 和 debug 级 debug 级别是最低的
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	// info 文件 WriteSyncer
	infoFileWriteSyncer := getInfoWriterSyncer()
	// error 文件 WriteSyncer
	errorFileWriteSyncer := getErrorWriterSyncer()

	env, err := conf.GetConfigParam("appEnv")
	if err != nil {
		panic(err)
	}
	var infoFileCore, errorFileCore zapcore.Core
	if env == "development" {
		// 开发环境 同时输出到控制台 和 指定的日志文件中
		infoFileCore = zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority)
		errorFileCore = zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority)
	} else if env == "production" {
		// 生产环境 仅输出到日志 不输出到控制台
		infoFileCore = zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer), lowPriority)
		errorFileCore = zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer), highPriority)

	} else {
		errStr := errors.New(fmt.Sprintf("Running environment err: %v", env))
		panic(errStr)
	}

	// 将 infoCore 和 errCore 加入core切片
	var coreArr []zapcore.Core
	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)

	// 生成 logger
	Logger = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()) // 添加输出文件与行号选项
	SugarLogger = Logger.Sugar()
}

// getEncoder 修改时间编码格式 日志级别大写
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 记录error以下日志级别的文件
func getInfoWriterSyncer() zapcore.WriteSyncer {
	path, err := conf.GetConfigParam("logFilePath")
	if err != nil {
		panic(err)
	}
	dir := conf.RunningServerName
	fileName := getDailyLoggerName()
	file := fmt.Sprintf("%v/%v/info_%v", path, dir, fileName)
	//引入第三方库 Lumberjack 加入日志切割功能
	infoWriteSyncer := &lumberjack.Logger{
		Filename:   file,  // 文件名称
		MaxSize:    10,    // 日志最大大小 单位 MB
		MaxBackups: 100,   // 保留旧日志最大数量
		MaxAge:     28,    // 保留旧日志最大天数
		Compress:   false, // Compress确定是否应该使用gzip压缩已旋转的日志文件 默认不执行压缩。
	}
	return zapcore.AddSync(infoWriteSyncer)
}

// 记录error及以上日志级别的文件
func getErrorWriterSyncer() zapcore.WriteSyncer {
	path, err := conf.GetConfigParam("logFilePath")
	if err != nil {
		panic(err)
	}
	dir := conf.RunningServerName
	fileName := getDailyLoggerName()
	file := fmt.Sprintf("%v/%v/err_%v", path, dir, fileName)
	//引入第三方库 Lumberjack 加入日志切割功能
	errorWriteSyncer := &lumberjack.Logger{
		Filename:   file,  // 文件名称
		MaxSize:    10,    // 日志最大大小 单位 MB
		MaxBackups: 100,   // 保留旧日志最大数量
		MaxAge:     28,    // 保留旧日志最大天数
		Compress:   false, // Compress确定是否应该使用gzip压缩已旋转的日志文件 默认不执行压缩。
	}
	return zapcore.AddSync(errorWriteSyncer)
}

// getDailyLoggerName 每天单独创建日志文件 返回日志文件名称 名称相同时不用重新创建直接打开
func getDailyLoggerName() string {
	now := time.Now().Format(time.DateOnly)
	suffix := strings.ReplaceAll(now, "-", "")
	return fmt.Sprintf("log_%v.log", suffix)
}
