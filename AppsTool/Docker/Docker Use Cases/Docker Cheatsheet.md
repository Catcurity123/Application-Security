#### A. What is an image
(+) Docker image is an object that contains an OS filesystem, an application, and all application dependencies. If you work in operations, it’s like a virtual machine template.

(+) `docker image ls`: list all docker image
(+) `docker image pull ubuntu:latest`: pulling an image from docker host

(+) It’s also worth noting that each image gets its own unique ID. When referencing images, you can refer to them using either IDs or names. If you’re working with image ID’s, it’s usually enough to type the first few characters of the ID — as long as it’s unique, Docker will know which image you mean.

#### B. What is a container
(+) Container is a running instance of an image
(+) `docker container run ubuntu:latest /bin/bash`: run /bin/bash for ubuntu:latest
	(-) `-it`: for interactivity
	(-) `-d`: to detach, meaning run the container in the background
	(-) `-v`: specify a volume
	(-) `-p`: specify port mapping

(+) When inside a container, `ctrl + Z` then `ctrl + C` to exit without destroy the container.
	(-) To resume connection with a container: `docker exec -it <id> /bin/bash`

(+) `docker container ls -a`: to list all existing container on the machine

(+) `docker start -i <id>`: To start a container 
(+) `docker stop <container name>`: To stop a container
(+) `docker rm <container name>`: to remove a container

#### C. How to build an image
(+) `docker build -t <repository>/<additional_info>:<tag> <where to build (.)>`: build an image from a dockerfile
==> `docker build -t test/webapptest:v1.0 .

#### D. How to run a container from an image
(+) `docker container run -d --name <container name> --publish <port mapping> <repository>/<tag>`: run a container in detached mode with name and port mapping from an image
==> `docker container run -d --name webapptest --publish 8080:8080 test/webapptest:v1.0`
(+) `docker run --rm -it app/httpd /bin/sh`: run then discard the container


#### E. Remove dangling
(+) `docker rmi -f $(docker images -f "dangling=true" -q)`:  Remove dangling

#### F. Copy file to Docker Container
(+) `docker cp .\Ansible_Test_Key.ppk 8115384d0752:/home_folder`: Copy file to a folder in docker container.

#### G. Resume a connection to a docker container
(+) `docker exec -it <container_name_or_id> /bin/bash`: Enter bash of a contaienr
(+) `docker start -i <id>`: start an exited container

#### H. How to filter docker images
(+) `docker image ls --filter dangling =true`: filter all dangling image
	(-) `before + <image ID>`: Filter image before sth
	(-) `since + <image ID>`: Filter since
	(-) `label + <label>`: filter base on label

#### I. Pushing image
(+) `docker login`
(+) `docker image tag <image_name> <repository>/<additional_info>:<tag>`
==> `docker image tag myhello dixluwn/testk8s:v1.0`
(+) `docker image push <repository>/<additional_info>:<tag>`
==> `docker image push dixluwn/testk8s:v1.0`
![[Pasted image 20240410135044.png |center]]



#### J. Convert key in linux instance

###### 1. Install putty-tools
```
sudo apt-get update
sudo apt-get install putty-tools
```

###### 2. Convert .ppk to .pem
```
puttygen your_key.ppk -O private-openssh -o your_key.pem
chmod 600 your_key.pem
```

###### 3. SSH into Linux
```
ssh -i your_key.pem user@hostname

==> ssh -i "Ansible_Test_Key.pem" ec2-user@ec2-3-217-12-8.compute-1.amazonaws.com
```