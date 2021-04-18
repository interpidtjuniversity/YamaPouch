package main

import (
	_ "./nsdeploy"
	log "github.com/Sirupsen/logrus"
	"os"
	"os/exec"
	"strings"
)

const ENV_EXEC_DEPLOY_CMD = "mydocker_deploy_cmd"
const ENV_EXEC_DEPLOY_APP_KILL_PRETREATMENT = "mydocker_deploy_kill_pretreatment"
//                                                                                                   0 false   1 true
func DeployAppInContainer(containerName string, comArray []string, kill bool) {
	pid, err := GetContainerPidByName(containerName)
	if err != nil {
		log.Errorf("Exec container getContainerPidByName %s error %v", containerName, err)
		return
	}

	cmdStr := strings.Join(comArray, " ")
	log.Infof("container pid %s", pid)
	log.Infof("command %s", cmdStr)

	cmd := exec.Command("/proc/self/exe", "deploy")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	os.Setenv(ENV_EXEC_PID, pid)
	os.Setenv(ENV_EXEC_DEPLOY_CMD, cmdStr)
	if kill {
		os.Setenv(ENV_EXEC_DEPLOY_APP_KILL_PRETREATMENT, "true")
	}

	containerEnvs := getEnvsByPid(pid)
	cmd.Env = append(os.Environ(), containerEnvs...)

	if err := cmd.Run(); err != nil {
		log.Errorf("Deploy in container %s error %v", containerName, err)
	}
}


