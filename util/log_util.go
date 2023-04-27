package util

import (
	"github.com/prometheus/common/log"
	"runtime/debug"
)

/**
 * @Author: fengxiaoxiao /13156050650@163.com
 * @Desc:
 * @Version: 1.0.0
 * @Date: 2022/5/19 11:08 下午
 */

// Info logs a message using INFO as log level.
func Info(args ...interface{}) {
	debug.PrintStack()
	log.Info(args...)
}

// Infof logs a message using INFO as log level.
func Infof(format string, args ...interface{}) {
	debug.PrintStack()
	log.Infof(format, args...)
}

// Debug logs a message using DEBUG as log level.
func Debug(args ...interface{}) {
	debug.PrintStack()
	log.Debug(args...)
}

// Debugf logs a message using DEBUG as log level.
func Debugf(format string, args ...interface{}) {
	debug.PrintStack()
	log.Debugf(format, args...)
}

func Warning(args ...interface{}) {
	debug.PrintStack()
	log.Warn(args...)
}

func Warningf(format string, args ...interface{}) {
	debug.PrintStack()
	log.Warnf(format, args...)
}

func Error(args ...interface{}) {
	debug.PrintStack()
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	debug.PrintStack()
	log.Errorf(format, args...)
}
