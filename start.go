package main

import (
	"fmt"
	"os"

	"github.com/opencontainers/runc/libcontainer"
	"github.com/urfave/cli"
)

var startCommand = cli.Command{
	Name:  "start",
	Usage: "executes the user defined process in a created container",
	ArgsUsage: `<container-id>

Where "<container-id>" is your name for the instance of the container that you
are starting. The name you provide for the container instance must be unique on
your host.`,
	Description: `The start command executes the user defined process in a created container.`,
	Action: func(context *cli.Context) error {
		f, _ := os.OpenFile("/tmp/check.txt", os.O_WRONLY|os.O_APPEND, 0666)
		defer f.Close()
		f.WriteString("check start operation!!!")
		container, err := getContainer(context)
		if err != nil {
			return err
		}
		status, err := container.Status()
		if err != nil {
			return err
		}
		switch status {
		case libcontainer.Created:
			return container.Exec()
		case libcontainer.Stopped:
			return fmt.Errorf("cannot start a container that has run and stopped")
		case libcontainer.Running:
			return fmt.Errorf("cannot start an already running container")
		default:
			return fmt.Errorf("cannot start a container in the %s state", status)
		}
	},
}
