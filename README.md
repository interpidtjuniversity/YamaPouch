# YamaPouch
simple app container like docker and pouch

## new network
##### command : pouch network create --driver bridge --subnet {network} {networkname}
##### sample : pouch network create --driver bridge --subnet 192.168.10.1/24 basebridge

## new container
##### command : pouch run -d --name {containername} -net {networkname} -p {port mapping} {imagename} {init process command}
##### output : local ip of container
##### sample : pouch run -d --name spring-prod01 -net basebridge -p 8080:8080 JavaContainer top -b

## deploy app in container
##### command : pouch deploy --name {containername} --app-log-path {app log path} --deploy-path {app process pid txt} {the command to start app}
##### sample : pouch deploy --name spring-prod01 --app-log-path /root/logs/log.log --deploy-path /root/pid.txt /jdk1.8.0_281/bin/java -jar /root/demo.jar

## stop container
##### command : pouch stop {containername}
##### sample : pouch stop spring-prod01

## restart container()
##### command : pouch start -d --name {containername}  {imagename} {init process command}
##### sample : pouch start -d --name spring-prod01  JavaContainer top -b

## remove container
##### command : pouch rm {containername}
##### sample : pouch rm spring-prod01

## build image
##### command : pouch build --input {baseimagepath} --output {outputimagepath} --executable {executable path}
##### sample : pouch build --input /root/JavaContainer.tar --output /root/testapp_0410_release.tar --executable /root/demo.jar,/root/test.jar,/boot.sh


