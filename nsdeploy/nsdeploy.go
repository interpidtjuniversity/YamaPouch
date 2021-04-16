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
	char *mydocker_deploy_process_pid_path;
	mydocker_deploy_process_pid_path = getenv("mydocker_deploy_process_pid_path");
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

	char *mydocker_deploy_log_path;
	mydocker_deploy_log_path = getenv("mydocker_deploy_log_path");
	// create deploy script
	FILE *fpWrite=fopen("/root/deploy.sh","w");
	if(fpWrite==NULL){
        return;
    }
    fprintf(fpWrite,"%s","#!/bin/bash\n");
	if (needKill) {
		pid_t pid;
		FILE *fp = NULL;
		fp = fopen(mydocker_deploy_process_pid_path, "r");
		if(fp==NULL){
        	return;
    	}
		char buff[255];
		fscanf(fp, "%s", buff);
		int num = atoi(buff);
		fprintf(fpWrite,"children=$(ps --ppid %d | awk '{if($1~/[0-9]+/) print $1}')\n", num);
    	fprintf(fpWrite,"kill -15 %d\n", num);
    	fprintf(fpWrite,"%s","for ((i=0;i<${#children[@]};i++))\n");
    	fprintf(fpWrite,"%s","do\n");
    	fprintf(fpWrite," %s","kill -15 ${children[i]}\n");
    	fprintf(fpWrite,"%s","done\n");
	}
    fprintf(fpWrite,"echo $$ > \"%s\"\n",mydocker_deploy_process_pid_path);
    fprintf(fpWrite,"%s > %s\n",mydocker_deploy_cmd, mydocker_deploy_log_path);
    fclose(fpWrite);

	pid_t pid;
	pid = fork();
	if (pid == 0) {
		execl("/bin/bash", "/bin/bash","/root/deploy.sh",NULL);
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