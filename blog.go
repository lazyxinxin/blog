package blog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	// 初始化Logrus

	log.Out = os.Stdout
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors:    false, // 启用颜色输出
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02T15:04:05Z07:00",
		CallerPrettyfier: callerPrettyfier,
	})
	log.SetLevel(logrus.InfoLevel) // 设置日志级别为Info
	log.SetReportCaller(true)
}

func callerPrettyfier(f *runtime.Frame) (string, string) {
	filename := filepath.Base(f.File) // 只获取文件名，而不是完整路径
	line := fmt.Sprintf("%d", f.Line)
	return filename, line
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}
