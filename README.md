# YamaPouch
simple app container like docker and pouch(java(done), python(wait), go(wait) image)

## new network
##### command : pouch network create --driver bridge --subnet {network} {networkname}
##### sample : pouch network create --driver bridge --subnet 192.168.10.1/24 basebridge

## new container
##### command : pouch run -d --name {containername} -net {networkname} -p {port mapping} {imagename} {init process command}
##### output : local ip of container
##### sample : pouch run -d --name spring-prod01 -net basebridge -p 8080:8080 JavaContainer top -b

## deploy app in container
##### command : pouch deploy --name {containername} {-kill} {the command to start app}
##### sample : pouch deploy --name spring-prod01 /jdk1.8.0_281/bin/java -jar /root/demo.jar        (first deployment)
##### sample : pouch deploy --name spring-prod01 -kill /jdk1.8.0_281/bin/java -jar /root/demo.jar  (later deployment overrides previous)

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


