# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/andyguwc/kongExample
WORKDIR /go/src/github.com/andyguwc/kongExample

# Install Gorilla Mux & other dependencies
RUN go get github.com/gorilla/mux

# Install our package
RUN go install github.com/andyguwc/kongExample

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/kongExample
