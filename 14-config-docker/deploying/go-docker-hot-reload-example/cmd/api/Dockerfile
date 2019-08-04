# https://www.zachjohnsondev.com/posts/go-docker-hot-reload-example/

FROM golang:1.10

# WORKDIR sets the current working directory within the container to the path for this example repository.
WORKDIR /go/src/github.com/Zach-Johnson/go-docker-hot-reload-example

# Then the COPY command copies everything in the current context directory to the container. 
# We’re going to use a docker-compose file for this, so the context directory will be whatever we define it to be in the docker-compose file or the directory we run the docker-compose command from by default.

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

# hot reloader
# The ENTRYPOINT defines the command that will be ran when the container starts up. Here we run the CompileDaemon command

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./cmd/api/" -command="./api"
