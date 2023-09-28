package log

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/7/15 下午6:49
 */
import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/op/go-logging"
	"os"
	"time"
)

var logger *logging.Logger

func InitLog(format string, level string, rotationTime int, maxAge int) {
	formatter, err := logging.NewStringFormatter(format)
	if err != nil {
		formatter, err = logging.NewStringFormatter("%{color}%{time:2006-01-02 15:04:05.000 CST} [%{module}] %{shortfile:15s} %{shortfunc:-10s} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")
	}
	logBackend := logging.NewLogBackend(os.Stdout, "", 0)
	if rotationTime == 0 {
		rotationTime = 1
	}
	if maxAge == 0 {
		maxAge = 30
	}
	fileWriter, err := rotatelogs.New(
		"log/client-%Y%m%d.log",
		rotatelogs.WithMaxAge(time.Duration(maxAge)*24*time.Hour),             // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(rotationTime)*24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		fmt.Println(err)

	}
	logBackendFile := logging.NewLogBackend(fileWriter, "", 0)

	logLevel, err := logging.LogLevel(level)
	if err != nil {
		logLevel = logging.DEBUG
	}

	levelBackend := logging.AddModuleLevel(logBackend)

	levelBackendFile := logging.AddModuleLevel(logBackendFile)

	levelBackend.SetLevel(logLevel, "")

	if logLevel > logging.WARNING {
		logLevel = logging.WARNING
	}
	levelBackendFile.SetLevel(logLevel, "")
	// backend := logging.NewBackendFormatter(levelBackend, formatter)
	logging.SetBackend(levelBackend, levelBackendFile)
	logging.SetFormatter(formatter)
	logger = logging.MustGetLogger("default")
	logger.ExtraCalldepth = 1
}

// Info logs a message using INFO as log level.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof logs a message using INFO as log level.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Debug logs a message using DEBUG as log level.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf logs a message using DEBUG as log level.
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Warning(args ...interface{}) {
	logger.Warning(args...)
}

func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func MustGetLogger(module string) Logger {
	return logging.MustGetLogger(module)
}

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}
