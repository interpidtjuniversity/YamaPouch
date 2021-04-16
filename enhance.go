package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/xianlubird/mydocker/container"
	"os"
	"os/exec"
)

// if container not in running, reject if
func EnhanceContainer(containerName string, executable []string, storePath string) error {
	cinf, err := GetContainerInfoByName(containerName)
	if err != nil {
		return err
	}
	if cinf.Status != "running" {
		return fmt.Errorf("container %s not in running", containerName)
	}

	physicPath := fmt.Sprintf("%s%s",fmt.Sprintf(container.MntUrl, containerName), storePath)
	if err := os.MkdirAll(physicPath, 0622); err != nil {
		log.Errorf("Mkdir %s error %v", physicPath, err)
		return err
	}

	for _, path := range executable {
		if Exist(path) {
			exec.Command("cp","-r", path, physicPath).CombinedOutput()
		}
	}

	return nil
}
