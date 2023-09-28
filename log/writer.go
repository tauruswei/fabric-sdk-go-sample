package log

import (
	"github.com/op/go-logging"
	"io"
)

type Writer struct {
	io.Writer
}

var customLogger = logging.MustGetLogger("custom")

func (w Writer) Write(p []byte) (n int, err error) {
	customLogger.Info(string(p))
	return len(p), nil
}
