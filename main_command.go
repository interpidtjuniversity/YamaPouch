package main

import (
	"./cgroups/subsystems"
	"./container"
	"./network"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: `Create a container with namespace and cgroups limit ie: mydocker run -ti [image] [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		cli.BoolFlag{
			Name:  "d",
			Usage: "detach container",
		},
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "container name",
		},
		cli.StringFlag{
			Name:  "v",
			Usage: "volume",
		},
		cli.StringSliceFlag{
			Name:  "e",
			Usage: "set environment",
		},
		cli.StringFlag{
			Name:  "net",
			Usage: "container network",
		},
		cli.StringSliceFlag{
			Name: "p",
			Usage: "port mapping",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		var cmdArray []string
		for _, arg := range context.Args() {
			cmdArray = append(cmdArray, arg)
		}

		//get image name
		imageName := cmdArray[0]
		cmdArray = cmdArray[1:]

		createTty := context.Bool("ti")
		detach := context.Bool("d")

		if createTty && detach {
			return fmt.Errorf("ti and d paramter can not both provided")
		}
		resConf := &subsystems.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuSet:      context.String("cpuset"),
			CpuShare:    context.String("cpushare"),
		}
		log.Infof("createTty %v", createTty)
		containerName := context.String("name")
		volume := context.String("v")
		network := context.String("net")

		envSlice := context.StringSlice("e")
		portmapping := context.StringSlice("p")

		Run(createTty, cmdArray, resConf, containerName, volume, imageName, envSlice, network, portmapping)
		return nil
	},
}

var deployCommand = cli.Command{
	Name: "deploy",
	Usage: "Run a application command in a container",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "name",
			Usage: "container name",
		},
		cli.BoolFlag{
			Name: "kill",
			Usage: "kill",
		},
	},
	Action: func(context *cli.Context) error{
		//This is for callback
		if os.Getenv(ENV_EXEC_PID) != "" {
			log.Infof("pid callback pid %s", os.Getgid())
			return nil
		}
		if len(context.Args()) < 2 {
			return fmt.Errorf("Missing command or arguments")
		}

		containerName := context.String("name")
		kill := context.Bool("kill")
		if containerName == "" {
			return fmt.Errorf("no container specitied")
		}
		// TODO check if container is exist
		//
		var command []string
		command = append(command, context.Args().Get(0))
		command = append(command, context.Args().Tail()...)
		DeployAppInContainer(containerName, command, kill)
		return nil
	},
}

var executeCommand = cli.Command{
	Name: "execute",
	Usage: "execute command in container(will return!!!)",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "name",
			Usage: "container name",
		},
	},
	Action: func(context *cli.Context) error{
		//This is for callback
		if os.Getenv(ENV_EXEC_PID) != "" {
			log.Infof("pid callback pid %s", os.Getgid())
			return nil
		}
		if len(context.Args()) < 2 {
			return fmt.Errorf("Missing command or arguments")
		}
		containerName := context.String("name")
		if containerName == "" {
			return fmt.Errorf("no container specitied")
		}
		var command []string
		command = append(command, context.Args().Get(0))
		command = append(command, context.Args().Tail()...)
		ExecuteCommandInContainer(containerName, command)
		return nil
	},
}

var executeOnceCommand = cli.Command{
	Name: "executeonce",
	Usage: "execute batch of command in one process, to avoid process concurrence problems",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "name",
			Usage: "container name",
		},
	},
	Action: func(context *cli.Context) error{
		//This is for callback
		if os.Getenv(ENV_EXEC_PID) != "" {
			log.Infof("pid callback pid %s", os.Getgid())
			return nil
		}
		containerName := context.String("name")
		if containerName == "" {
			return fmt.Errorf("no container specitied")
		}
		var commands []string
		commands = append(commands, context.Args().Get(0))
		commands = append(commands, context.Args().Tail()...)
		ExecuteBatchCommandOnceInContainer(containerName, commands)
		return nil
	},
}

var enhanceCommand = cli.Command{
	Name: "enhance",
	Usage: "Send app to container if conflict then remove old",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "name",
			Usage: "name of the container to enhance",
		},
		cli.StringSliceFlag{
			Name: "executable",
			Usage: "enhancer",
		},
		cli.StringFlag{
			Name: "storepath",
			Usage: "path to store executable",
		},
	},
	Action: func(context *cli.Context) error{
		containerName := context.String("name")
		executable := context.StringSlice("executable")
		storepath := context.String("storepath")
		if containerName == "" {
			return fmt.Errorf("container is not specified")
		}
		return EnhanceContainer(containerName, executable, storepath)
	},
}

var startCommand = cli.Command {
	Name: "start",
	Usage: "Start a container which is stopped and reuse it's network",
	Flags: []cli.Flag {
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		cli.BoolFlag{
			Name:  "d",
			Usage: "detach container",
		},
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "container name",
		},
		cli.StringSliceFlag{
			Name:  "e",
			Usage: "set environment",
		},
	},
	Action: func(context *cli.Context) error{
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		var cmdArray []string
		for _, arg := range context.Args() {
			cmdArray = append(cmdArray, arg)
		}

		createTty := context.Bool("ti")
		detach := context.Bool("d")

		if createTty && detach {
			return fmt.Errorf("ti and d paramter can not both provided")
		}
		resConf := &subsystems.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuSet:      context.String("cpuset"),
			CpuShare:    context.String("cpushare"),
		}
		log.Infof("createTty %v", createTty)
		containerName := context.String("name")

		envSlice := context.StringSlice("e")

		Start(createTty, cmdArray, resConf, containerName, envSlice)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		err := container.RunContainerInitProcess()
		return err
	},
}

var listCommand = cli.Command{
	Name:  "ps",
	Usage: "list all the containers",
	Action: func(context *cli.Context) error {
		ListContainers()
		return nil
	},
}

var logCommand = cli.Command{
	Name:  "logs",
	Usage: "print logs of a container",
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Please input your container name")
		}
		containerName := context.Args().Get(0)
		logContainer(containerName)
		return nil
	},
}

var appLogCommand = cli.Command{
	Name: "applogs",
	Usage: "print latest app running log",
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Please input your container name")
		}
		containerName := context.Args().Get(0)
		logpath := context.Args().Get(1)
		appLogContainer(containerName, logpath)
		return nil
	},
}

var execCommand = cli.Command{
	Name:  "exec",
	Usage: "exec a command into container",
	Action: func(context *cli.Context) error {
		//This is for callback
		if os.Getenv(ENV_EXEC_PID) != "" {
			log.Infof("pid callback pid %s", os.Getgid())
			return nil
		}

		if len(context.Args()) < 2 {
			return fmt.Errorf("Missing container name or command")
		}
		containerName := context.Args().Get(0)
		var commandArray []string
		for _, arg := range context.Args().Tail() {
			commandArray = append(commandArray, arg)
		}
		ExecContainer(containerName, commandArray)
		return nil
	},
}

var stopCommand = cli.Command{
	Name:  "stop",
	Usage: "stop a container",
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container name")
		}
		containerName := context.Args().Get(0)
		stopContainer(containerName)
		return nil
	},
}

var removeCommand = cli.Command{
	Name:  "rm",
	Usage: "remove unused containers",
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container name")
		}
		network.Init()
		containerName := context.Args().Get(0)
		removeContainer(containerName)
		return nil
	},
}

var commitCommand = cli.Command{
	Name:  "commit",
	Usage: "commit a container into image",
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 2 {
			return fmt.Errorf("Missing container name and image name")
		}
		containerName := context.Args().Get(0)
		imageName := context.Args().Get(1)
		commitContainer(containerName, imageName)
		return nil
	},
}

// pouch build --base xxx --executable xxx
var buildImageCommand = cli.Command{
	Name: "build",
	Usage: "tar executable file into image",
	Flags: []cli.Flag {
		cli.StringSliceFlag{
			Name: "executable",
			Usage: "executable file in image",
		},
		cli.StringFlag{
			Name: "input",
			Usage: "base image path",
		},
		cli.StringFlag{
			Name: "output",
			Usage: "build image path",
		},
	},
	Action: func(context *cli.Context) error{

		input := context.String("input")
		output := context.String("output")

		executables := context.StringSlice("executable")
		if input == "" || output == ""{
			return fmt.Errorf("input and output must be specified")
		}
		return BuildImage(input, output, executables)
	},
}

var networkCommand = cli.Command{
	Name:  "network",
	Usage: "container network commands",
	Subcommands: []cli.Command {
		{
			Name: "create",
			Usage: "create a container network",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "driver",
					Usage: "network driver",
				},
				cli.StringFlag{
					Name:  "subnet",
					Usage: "subnet cidr",
				},
			},
			Action:func(context *cli.Context) error {
				if len(context.Args()) < 1 {
					return fmt.Errorf("Missing network name")
				}
				network.Init()
				err := network.CreateNetwork(context.String("driver"), context.String("subnet"), context.Args()[0])
				if err != nil {
					return fmt.Errorf("create network error: %+v", err)
				}
				return nil
			},
		},
		{
			Name: "list",
			Usage: "list container network",
			Action:func(context *cli.Context) error {
				network.Init()
				network.ListNetwork()
				return nil
			},
		},
		{
			Name: "remove",
			Usage: "remove container network",
			Action:func(context *cli.Context) error {
				if len(context.Args()) < 1 {
					return fmt.Errorf("Missing network name")
				}
				network.Init()
				err := network.DeleteNetwork(context.Args()[0])
				if err != nil {
					return fmt.Errorf("remove network error: %+v", err)
				}
				return nil
			},
		},
	},
}
