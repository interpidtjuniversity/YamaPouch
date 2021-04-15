package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/google/uuid"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

// unzip input to tmp path and copy executable to /root then zip it to output
func BuildImage(input, output string, executable []string) error {
	if !Exist(input) {
		return fmt.Errorf("input image not exist %s", input)
	}
	if path.Ext(input) != ".tar" {
		return fmt.Errorf("type of input %s is not .tar", input)
	}


	tmpPath := generatePath()
	if err := os.MkdirAll(tmpPath, 0622); err != nil {
		log.Errorf("Mkdir %s error %v", tmpPath, err)
		return err
	}
	if _, err := exec.Command("tar", "-xvf", input, "-C", tmpPath).CombinedOutput(); err != nil {
		log.Errorf("Untar dir %s error %v", tmpPath, err)
		return err
	}
	// copy executable to img/root
	for _, path := range executable {
		if Exist(path) {
			exec.Command("cp","-r", path, tmpPath+"root").CombinedOutput()
		}
	}
	// tar it to output
	if _, err := exec.Command("tar", "-czf", output, "-C", tmpPath, ".").CombinedOutput(); err != nil {
		log.Errorf("Tar folder %s error %v", tmpPath, err)
	}
	// delete tmpPath
	return os.RemoveAll(tmpPath)
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func generatePath() string {
	uid, _ := uuid.NewUUID()
	y,m,d := time.Now().Date()
	id := fmt.Sprintf("%s_%s", strings.ReplaceAll(uid.String(), "-",""), fmt.Sprintf("%d_%d_%d", y,m,d))
	return fmt.Sprintf("/tmp/pouch/%s/",id)
}
