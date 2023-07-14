package logger

import (
	"vue3-bashItem/pkg/settings"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

var Logger *zap.Logger
var FileLogger *zap.SugaredLogger

func Setup() {
	InitConselLogger()
	InitFileLogger()
}

func InitConselLogger() {
	encoder := getEncoder("console")
	var l zapcore.Level
	if settings.ServerSetting.RunMode == "debug" {
		l = zap.DebugLevel
	} else if settings.ServerSetting.RunMode == "release" {
		l = zap.ErrorLevel
	} else {
		l = zap.InfoLevel
	}
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), l)
	Logger = zap.New(core, zap.AddCaller())
}

func InitFileLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder("json")
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(settings.LogConfSetting.Level))
	if err != nil {
		log.Fatalf("create logger failed!: %v", err)
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	fl := zap.New(core, zap.AddCaller())
	FileLogger = fl.Sugar()
}

func getEncoder(outType string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	if outType == "console" {
		return zapcore.NewConsoleEncoder(encoderConfig)
	} else if outType == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	} else {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
}
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   settings.LogConfSetting.Filename,   // 日志文件名
		MaxSize:    settings.LogConfSetting.MaxSize,    // 日志文件最大大小(MB)
		MaxBackups: settings.LogConfSetting.MaxBackups, // 保留旧文件最大数量 从零位算 例写2 则保留3个备份文件 算上debug.log是4个
		MaxAge:     settings.LogConfSetting.MaxAge,     // 保留旧文件最长天数
	}
	return zapcore.AddSync(lumberJackLogger)
}

/* 调用命令

listenInfo := "aaaaa"
logger.Logger.Info(listenInfo)      # 屏幕输出
logger.FileLogger.Info(listenInfo)  # 写入文件

bb := fmt.Sprintf("%v用户已存在",createUser.UserName)
logger.FileLogger.Error(bb)
*/
