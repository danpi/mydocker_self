package mydocker_self

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var runCommand=cli.Command{
	Name:"run",
	Usage:`Create a container with namespace and cgroups limit ie: mydocker run -ti [image] [command]`,
	Flags:[]cli.Flag{
		cli.BoolFlag{
			Name:"ti",
			Usage:"enable tty",
		},
	},
	Action: func(context *cli.Context)error {
		if len(context.Args())<1{
			return fmt.Errorf("missing container command")
		}
		//cmd:=context.Args().Get(0)
		//tty:=context.Bool("ti")
		//Run(cmd,tty)
		return nil
	},
}

var initCommand=cli.Command{
	Name:                   "init",
	Usage:                  "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context)error {
		log.Infof("init come on")
		//err:=container.RunContainerInitProcess()
		return nil
	},

}
