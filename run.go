package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/danpi/mydocker_self/container"
	"os"
)

func Run(tty bool,command string){
	parent:=container.NewParentProcess(tty,command)
	if err:=parent.Start();err!=nil{
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
