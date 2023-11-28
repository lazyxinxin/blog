package blog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

const (
	OutPut_Console = "console"
	OutPut_File    = "file"
)

type Blog struct {
	IsDebug      bool
	Log          logrus.Logger
	OutPutMethod string
}

func NewBlog(IsDebug bool, OutPutMethod string) *Blog {
	blog := Blog{
		IsDebug:      IsDebug,
		Log:          *logrus.New(),
		OutPutMethod: OutPutMethod,
	}
	// 判断是否是本地conosle输出
	if OutPutMethod == "console" {
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
	return &blog
}

func callerPrettyfier(f *runtime.Frame) (string, string) {
	filename := filepath.Base(f.File) // 只获取文件名，而不是完整路径
	line := fmt.Sprintf("%d", f.Line)
	return filename, line
}

func (b *Blog) Info(args ...interface{}) {
	log.Info(args...)
}

func (b *Blog) Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func (b *Blog) Warn(args ...interface{}) {
	log.Warn(args...)
}

func (b *Blog) Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func (b *Blog) Error(args ...interface{}) {
	log.Error(args...)
}

func (b *Blog) Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func (b *Blog) Debug(args ...interface{}) {
	if b.IsDebug {
		log.Debug(args...)
	}

}

func (b *Blog) Debugf(format string, args ...interface{}) {
	if b.IsDebug {
		log.Debugf(format, args...)
	}
}
