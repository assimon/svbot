package log

import (
	"fmt"
	"github.com/assimon/svbot/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

var Sugar *zap.SugaredLogger

func init() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	Sugar = zap.New(core, zap.AddCaller()).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	logC := config.GetLogInstance()
	logPath := fmt.Sprintf("%s/%s", config.GetGlobalInstance().AppPath, "/log")
	file := fmt.Sprintf("%s/log_%s.log",
		logPath,
		time.Now().Format("20060102"))
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,
		MaxSize:    logC.MaxSize,
		MaxBackups: logC.MaxBackups,
		MaxAge:     logC.MaxAge,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
