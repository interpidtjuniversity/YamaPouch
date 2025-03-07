package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"./container"
	"io/ioutil"
	"os"
	"text/tabwriter"
)

func ListContainers(print bool) []*container.ContainerInfo{
	dirURL := fmt.Sprintf(container.DefaultInfoLocation, "")
	dirURL = dirURL[:len(dirURL)-1]
	files, err := ioutil.ReadDir(dirURL)
	if err != nil {
		log.Errorf("Read dir %s error %v", dirURL, err)
		return nil
	}

	var containers []*container.ContainerInfo
	for _, file := range files {
		if file.Name() == "network" {
			continue
		}
		tmpContainer, err := getContainerInfo(file)
		if err != nil {
			log.Errorf("Get container info error %v", err)
			continue
		}
		containers = append(containers, tmpContainer)
	}

	if print {
		w := tabwriter.NewWriter(os.Stdout, 12, 1, 3, ' ', 0)
		fmt.Fprint(w, "ID\tNAME\tPID\tSTATUS\tCOMMAND\tCREATED\tPORTMAPPING\tIP\tNETWORKNAME\n")
		for _, item := range containers {
			portMapping, _ := json.Marshal(item.PortMapping)
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
				item.Id,
				item.Name,
				item.Pid,
				item.Status,
				item.Command,
				item.CreatedTime, string(portMapping), item.Ip, item.NetWorkName)
		}
		if err := w.Flush(); err != nil {
			log.Errorf("Flush error %v", err)
			return nil
		}
	}
	return containers
}

func getContainerInfo(file os.FileInfo) (*container.ContainerInfo, error) {
	containerName := file.Name()
	configFileDir := fmt.Sprintf(container.DefaultInfoLocation, containerName)
	configFileDir = configFileDir + container.ConfigName
	content, err := ioutil.ReadFile(configFileDir)
	if err != nil {
		log.Errorf("Read file %s error %v", configFileDir, err)
		return nil, err
	}
	var containerInfo container.ContainerInfo
	if err := json.Unmarshal(content, &containerInfo); err != nil {
		log.Errorf("Json unmarshal error %v", err)
		return nil, err
	}

	return &containerInfo, nil
}
