Overview
This is what we will do:

install Docker
run a software image in a container
locate an interesting image on Docker Hub
run the image on your own machine
modify an image to create your own and run it
create a Docker Hub account and repository
push your image to Docker Hub for others to share
Install
Install docker

Choose Docker engine / install

Verify installation
Check the docker version

docker version
Run a sample "hello world" application

docker run hello-world 
Images, Repository, Containers
You use an image to make a container.

You can make an unlimited number of containers from one image.

An image is stored in a docker image repository.

Docker hub is one docker image repository you can use.

See all containers (processes) on your machine

docker ps -a
See all running containers on your machine

docker ps