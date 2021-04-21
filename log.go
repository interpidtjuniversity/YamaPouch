package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"./container"
	"os"
	"io/ioutil"
)

func logContainer(containerName string) {
	dirURL := fmt.Sprintf(container.DefaultInfoLocation, containerName)
	logFileLocation := dirURL + container.ContainerLogFile
	file, err := os.Open(logFileLocation)
	defer file.Close()
	if err != nil {
		log.Errorf("Log container open file %s error %v", logFileLocation, err)
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Errorf("Log container read file %s error %v", logFileLocation, err)
		return
	}
	fmt.Fprint(os.Stdout, string(content))
}

// writeLayer
func appLogContainer(containerName, logPath string) {
	// container writeLayer
	dirUrl := fmt.Sprintf(container.WriteLayerUrl, containerName)
	// targetLogDir
	logUrl := fmt.Sprintf("%s%s", dirUrl, logPath)
	file, err := os.Open(logUrl)
	defer file.Close()
	if err != nil {
		log.Errorf("Log container open file %s error %v", logUrl, err)
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Errorf("Log container read file %s error %v", logUrl, err)
		return
	}
	fmt.Fprint(os.Stdout, string(content))
}
