package nsexecute

/*
#include <errno.h>
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>

void split(char *src,const char *separator,char **dest,int *num);

__attribute__((constructor)) void enterandexecute_namespace(void) {
	char *mydocker_pid;
	mydocker_pid = getenv("mydocker_pid");
	if (mydocker_pid) {
	} else {
		//fprintf(stdout, "missing mydocker_pid env skip nsenter");
		return;
	}
	char *mydocker_execute_cmd;
	mydocker_execute_cmd = getenv("mydocker_execute_cmd");
	if (mydocker_execute_cmd) {
		//fprintf(stdout, "got mydocker_execute_cmd=%s\n", mydocker_execute_cmd);
	} else {
		//fprintf(stdout, "missing mydocker_execute_cmd env skip nsenter");
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

	pid_t pid;
	pid = fork();
	if (pid == 0) {
    	char array[255] = {0};
    	strncpy(array, mydocker_execute_cmd, strlen(mydocker_execute_cmd) + 1);
		// 先限制长度为10吧...
		char *revbuf[10] = {0};
		int num = 0;
		split(array," ",revbuf,&num);
    	execl(revbuf[0], revbuf[0], revbuf[1], revbuf[2], revbuf[3],revbuf[4],revbuf[5],revbuf[6],revbuf[7], revbuf[8],revbuf[9],NULL);
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

void split(char *src,const char *separator,char **dest,int *num) {
	char *pNext;
	int count = 0;
	if (src == NULL || strlen(src) == 0)
		return;
	if (separator == NULL || strlen(separator) == 0)
		return;
	pNext = (char *)strtok(src,separator);
	while(pNext != NULL) {
		*dest++ = pNext;
		++count;
		pNext = (char *)strtok(NULL,separator);
	}
	*num = count;
}

*/
import "C"