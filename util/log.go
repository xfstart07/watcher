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
	zlog, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	return zlog
}
