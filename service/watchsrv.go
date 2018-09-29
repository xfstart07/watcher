// Author: Xu Fei
// Date: 2018/9/12
package service

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"github.com/xfstart07/watcher/config"
	"go.uber.org/zap"
)

func WatchFile() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				zlog.Sugar().Info(event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					zlog.Sugar().Info("write")
				}
				Store(event.Name)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				zlog.Error("watchErr",zap.Error(err))
			}
		}
	}()

	paths := config.Config.WatchPaths
	for idx := range paths {
		path := paths[idx]

		if err := watcher.Add(path); err != nil {
			zlog.Error("监控文件夹失败", zap.Error(err))
		}
	}
}
