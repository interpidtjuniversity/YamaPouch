package nsexecuteonce

/*
#include <errno.h>
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>

void batchsplit(char *src,const char *separator,char **dest,int *num);

__attribute__((constructor)) void enterandexecuteonce_namespace(void) {

	char *mydocker_pid;
	mydocker_pid = getenv("mydocker_pid");
	if (mydocker_pid) {
	} else {
		return;
	}

	char *mydocker_executeonce_cmd;
	mydocker_executeonce_cmd = getenv("mydocker_executeonce_cmd");
	if (mydocker_executeonce_cmd) {
	} else {
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

    	char array[2048] = {0};
    	strncpy(array, mydocker_executeonce_cmd, strlen(mydocker_executeonce_cmd) + 1);
		// 先限制长度为10吧...
		char *revbuf[10] = {0};
		int num = 0;
		batchsplit(array,"#",revbuf,&num);

		for(int i = 0;i < 10; i++) {
			if(revbuf[i]) {
				system(revbuf[i]);
			}
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

void batchsplit(char *src,const char *separator,char **dest,int *num) {
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