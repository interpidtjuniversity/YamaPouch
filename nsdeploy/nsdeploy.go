package nsdeploy

/*
#include <errno.h>
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
#include<unistd.h>

__attribute__((constructor)) void enteranddeploy_namespace(void) {
	char *mydocker_pid;
	mydocker_pid = getenv("mydocker_pid");
	if (mydocker_pid) {
	} else {
		//fprintf(stdout, "missing mydocker_pid env skip nsenter");
		return;
	}
	char *mydocker_deploy_cmd;
	mydocker_deploy_cmd = getenv("mydocker_deploy_cmd");
	if (mydocker_deploy_cmd) {
		//fprintf(stdout, "got mydocker_deploy_cmd=%s\n", mydocker_deploy_cmd);
	} else {
		//fprintf(stdout, "missing mydocker_deploy_cmd env skip nsenter");
		return;
	}
	int i;
	char nspath[1024];
	char *namespaces[] = { "ipc", "uts", "net", "pid", "mnt" };

	for (i=0; i<5; i++) {
		sprintf(nspath, "/proc/%s/ns/%s", mydocker_pid, namespaces[i]);
		int fd = open(nspath, O_RDONLY);

		if (setns(fd, 0) == -1) {
			//fprintf(stderr, "setns on %s namespace failed: %s\n", namespaces[i], strerror(errno));
		} else {
			//fprintf(stdout, "setns on %s namespace succeeded\n", namespaces[i]);
		}
		close(fd);
	}

	// kill time before procecss and reexecute it
	char *needKill;
	needKill = getenv("mydocker_deploy_kill_pretreatment");

	pid_t pid;
	pid = fork();
	if (pid == 0) {
		if (needKill) {
			pid_t pid;
			FILE *fp = NULL;
			fp = fopen("/root/deploy/pid.txt", "r");
			if(fp==NULL){
        		return;
    		}
			char buff[255];
			fscanf(fp, "%s", buff);
			fclose(fp);
			execl("/bin/bash", "/bin/bash","/root/deploy/kill-deploy.sh", buff, mydocker_deploy_cmd, NULL);
		} else {
			execl("/bin/bash", "/bin/bash","/root/deploy/deploy.sh", mydocker_deploy_cmd, NULL);
		}
		exit(0);
        return;
	} else if (pid > 0){
	} else {
		perror("container deploy process fork error");
        exit(1);
        return;
	}
    exit(0);
	return;
}
*/
import "C"