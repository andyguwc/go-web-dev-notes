// get any referenced external packages 
go get ./...   

// test current package and nested subdirectories
go test ./...

// test coverage report 
go test -cover 

// packages and files 
https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc

//build files 
go build inigo.go
./inigo


// web servers graceful shutdown
// systemd examples
systemctl start myapp.service
systemctl stop myapp.service
// recommend using an initialization daemon 


// errors
// returning an http status http.Error(w, "An Error Occurred", http.StatusForbidden)


// cross compile
// GOARCH specifies the hardware architecture such as amd64, 386, or arm, whereas GOOS specifies the operating system such as windows, linux, darwin, or freebsd.
$ GOOS=windows GOARCH=386 go build

// Use gox package 
$ gox \
-os="linux darwin windows " \
-arch="amd64 386" \
-output="dist/{{.OS}}-{{.Arch}}/{{.Dir}}" .

// performance runtime monitoring 


// pipeline output to a file 

go run main.go > index.html


// format all code

go fmt ./...


