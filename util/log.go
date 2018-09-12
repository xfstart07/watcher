// Author: Xu Fei
// Date: 2018/9/4
package util

import (
	"log"

	"go.uber.org/zap"
)

var zlog = NewLog()

// FIXME: use development config
func NewLog() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.DisableStacktrace = false
	config.DisableCaller = false
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}

	zlog, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}

	return zlog
}
