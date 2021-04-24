package main

import (
	_ "./nsexecuteonce"
	log "github.com/Sirupsen/logrus"
	"os"
	"os/exec"
	"strings"
)

const ENV_EXEC_EXECUTEONCE_CMD = "mydocker_executeonce_cmd"

func ExecuteBatchCommandOnceInContainer(containerName string, comArray []string) {
	pid, err := GetContainerPidByName(containerName)
	if err != nil {
		log.Errorf("Exec container getContainerPidByName %s error %v", containerName, err)
		return
	}

	// check
	commands := strings.Join(comArray," ")
	for _,command := range  strings.Split(commands, "#") {
		strings.Split(command, " ")
	}

	log.Infof("container pid %s", pid)
	log.Infof("command %s", comArray)

	cmd := exec.Command("/proc/self/exe", "executeonce")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	os.Setenv(ENV_EXEC_PID, pid)
	os.Setenv(ENV_EXEC_EXECUTEONCE_CMD, commands)

	containerEnvs := getEnvsByPid(pid)
	cmd.Env = append(os.Environ(), containerEnvs...)

	if err := cmd.Run(); err != nil {
		log.Errorf("execute in container %s error %v", containerName, err)
	}
}


