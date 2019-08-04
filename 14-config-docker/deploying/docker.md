# Deploy to Docker
https://docker-curriculum.com/#docker-compose
https://hashnode.com/post/docker-tutorial-for-beginners-cjrj2hg5001s2ufs1nker9he2


Containers are a different take on infrastructure virtualization, providing virtualization at the OS level and allowing resources to be partitioned through multiple isolated user space instance

Containers offer a logical packaging mechanism in which applications can be abstracted from the environment in which they actually run. This decoupling allows container-based applications to be deployed easily and consistently, regardless of whether the target environment is a private data center, the public cloud, or even a developer’s personal laptop. This gives developers the ability to create predictable environments that are isolated from rest of the applications and can be run anywhere.


The Docker daemon is the process that sits on the host OS that answers requests for service and orchestrates
the management of containers. Docker containers, containers for short, are lightweight virtualization of all the programs
that are needed to run a particular application, including the OS.

Docker containers are built on Docker images, which are read-only templates that help to launch containers.

You run containers from images. Docker images can be built in different ways. One way of doing it involves
using a set of instructions contained in a single file called the Dockerfile.

## Docker Components
Images - The blueprints of our application which form the basis of containers. In the demo above, we used the docker pull command to download the busybox image.

Containers - Created from Docker images and run the actual application. We create a container using docker run which we did using the busybox image that we downloaded. A list of running containers can be seen using the docker ps command.

Docker Daemon - The background service running on the host that manages building, running and distributing Docker containers. The daemon is the process that runs in the operating system to which clients talk to.

Docker Client - The command line tool that allows the user to interact with the daemon. More generally, there can be other forms of clients too - such as Kitematic which provide a GUI to the users.

Docker Hub - A registry of Docker images. You can think of the registry as a directory of all available Docker images. If required, one can host their own Docker registries and can use them for pulling images.


## Steps for Dockerizing Applications Locally
Deploy to local
- Create Dockerfile
- Build Docker image with Dockerfile
- Create Docker container from Docker image


build image from Dockerfile
docker build –t ws-d .

run the image to create and start the container
docker run --publish 80:8080 --name simple_web_service --rm ws-d

See active running containers 
docker ps

Use curl and POST request to the server to create a record
curl -i -X POST -H "Content-Type: application/json" -d '{"content":"My first
post","author":"Sau Sheong"}' http://127.0.0.1/post/


## Dockering in Cloud

Deploy to cloud
- Create docker host in cloud provider
- Connect to remote docker host
- Build Docker image in remote host
- Start Docker container in remote host

Use Docker Machine to create a Docker host in any of the cloud providers

An example to create Docker host on Digital Ocean 
docker-machine create --driver digitalocean --digitalocean-access-token<tokenwsd

Connect to Docker host on Digital Ocean

Change environment and build docker image


# Docker Commands
Load up container
$ docker run busybox

rm tag automatically removes the container as it exits
$ docker run --rm prakhar1989/static-site

-d is detached mode 
-P will publish all exposed ports to random ports and finally 
--name corresponds to a name we want to give
$ docker run -d -P --name static-site prakhar1989/static-site

See the ports 
$ docker port static-site
If using docker toolbox you might need to use below to get the IP
$ docker-machine ip default 
You can also specify a custom port to which the client will forward connections to the container.
docker run -p 8888:80 prakhar1989/static-site


View currently running containers
$ docker ps

View all container we ran and their statuses 
$ docker ps -a

Stop all current containers
$ docker stop $(docker ps -qa)

Run many commands in the container 
Running the run command with the -it flags attaches us to an interactive tty in the container. Now we can run as many commands in the container as we want. Type exit to exit.
docker run -it busybox sh

Remove specific container
$ docker rm 305297d7a235

Remove all exited containers
$ docker rm $(docker ps -a -q -f status=exited)

Remove docker image
$ docker rmi




# Dockerfile

Docker file 
The first line tells Docker to start from the golang image, which is a Debian image
with the latest Go installed, and a workspace configured at /go. The next two lines
copy the local code (in the current directory) to the container and set the working
directory accordingly. After that, you use the RUN command to tell Docker to get the
PostgreSQL driver and build the web service code, placing the executable binaries in
/go/bin. Once you have that, use the ENTRYPOINT command to tell Docker to run
/go/bin/ws-d by default whenever the container is started. Finally, use EXPOSE to
expose the port 8080 to other containers. Note that this doesn’t open up port 8080 to
the public; it simply opens up the port to other containers in the same machine.

```
# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/sausheong/mosaicgo

WORKDIR /go/src/github.com/sausheong/mosaicgo

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/sausheong/mosaic-docker

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/mosaic-docker

# Document that the service listens on port 8080.
EXPOSE 8080
```



```
FROM alpine:latest

RUN mkdir /app
WORKDIR /app
ADD consignment-service /app/consignment-service

CMD ["./consignment-service"]
```


Installing go dependencies
Git to be installed before using go 
```
FROM golang:1.7-alpine

ADD . /home
        
WORKDIR /home

RUN \
       apk add --no-cache bash git openssh && \
       go get -u github.com/minio/minio-go 
       

CMD ["go","run","sample.go"]
```



Choosing run vs. cmd vs. entrypoint
https://goinbigdata.com/docker-run-vs-cmd-vs-entrypoint/
- Use RUN instructions to build your image by adding layers on top of initial image.
- Prefer ENTRYPOINT to CMD when building executable Docker image and you need a command always to be executed. Additionally use CMD if you need to provide extra default arguments that could be overwritten from command line when docker container runs.
- Choose CMD if you need to provide a default command and/or arguments that can be overwritten from command line when docker container runs.




# Makefile
GOOS and GOARCH allow you to cross-compile your go binary for another operating system.

Here, we run our consignment-service docker image, exposing the port 50051. if you wanted to run this service on port 8080, you would change the -p argument to 8080:50051

```
build:
    ... 
    GOOS=linux GOARCH=amd64 go build
    docker build -t shippy-service-consignment .

run: 
    docker run -p 50051:50051 shippy-service-consignment

```

# Docker Compose
https://ewanvalentine.io/microservices-in-golang-part-3/
Docker-compose allows you do define a list of docker containers in a yaml file, and specify metadata about their run-time. Docker-compose services map more or less to the same docker commands we're already using. 

To build docker-compose stack 
$ docker-compose build

Run docker-compose
$ docker-compose run

To run in the background
$ docker-compose up -d

List of services
Each service has name. Includes a build path which is reference to a location containing a Dockerfile. You can also use image here to use a pre-built image


```
# docker-compose.yml
version: '3.1'

services:

  shippy-cli-consignment:
    build: ./shippy-cli-consignment

  shippy-service-consignment:
    build: ./shippy-service-consignment
    ports:
      - 50051:50051
    environment:
      MICRO_ADDRESS: ":50051"
      DB_HOST: "datastore:27017"

  shippy-service-vessel:
    build: ./shippy-service-vessel
    ports:
      - 50052:50051
    environment:
      MICRO_ADDRESS: ":50051"
```




## Docker Testing 


