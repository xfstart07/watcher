package main

import (
	"log"
	"syscall"

	"github.com/judwhite/go-svc/svc"
	"github.com/xfstart07/watcher/service"
	"github.com/xfstart07/watcher/config"
)

type program struct{}

func main() {
	prg := &program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
		log.Fatal(err)
	}
}

func (p *program) Init(env svc.Environment) error {
	log.Printf("is win service? %v\n", env.IsWindowsService())
	return nil
}

func (p *program) Start() error {
	log.Println("app start...")

	initApp()

	return nil
}

func (p *program) Stop() error {
	log.Println("app stop...")
	return nil
}

func initApp() {
	config.Load()
	service.ConnRedis()

	go service.StorePaths()
	service.WatchFile()
}
