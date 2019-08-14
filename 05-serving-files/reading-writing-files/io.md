# IO with readers and writers

Go models data input and output as a stream that flows from sources to targets. Data resources, such as files, networked connections, or even some in-memory objects, can be modeled as streams of bytes from which data can be read or written to

# io.Reader interface

The Read method returns the total number of bytes transferred into the provided slice and an error value (if necessary)


type Reader interface {
    Read(p []byte) (n int, err error)
}

Example of a trivial implementation of the io.Reader - alphaReader
Implementations should return an error value of io.EOF when the reader has no more data to transfer into stream p

```
type alphaReader string 

func(a alphaReader) Read(p []byte) (int, error) {
    count := 0
    for i:=0; i<len(n); i++ {
        if (a[i] >= 'A' && a[i] <='Z') || 
            (a[i] >= 'a' && a[i] <= 'z') {
                p[i] = a[i]
        }
        count++
    }
    return count, io.EOF
}

func main() {
    str := alphaReader("Hello! World")
    // since alphaReader type implements the io.Reader interface, can pass this as a parameter to io.Copy
    io.Copy(os.Stdout, &str)
    fmt.Println()
}

```

# io.Writer interface
The interface requires the implementation of a single method, Write(p []byte)(c int,
e error), that copies data from the provided stream p and writes that data to a sink
resource such as an in-memory structure, standard output, a file, a network connection, or any number of io.Writer implementations that come with the Go standard library. The Write method returns the number of bytes copied from p followed by an error value if any was encountered.

type Writer interface {
    Write(p []byte) (n int, err error)
}

Example implementation - channelWriter type, a writer that decomposes and serializes its stream that is sent over a Go channel as consecutive bytes

```
type channelWriter struct {
    Channel chan byte
}

func NewChannelWriter() *channelWriter {
    return &channelWriter{
        Channel: make(chan byte, 1024), 
    }
}

func (c *channelWriter) Write(p []byte) (int, error) {
    if len(p) == 0 {
        return 0, nil
    }

    go func() {
        defer close(c.Channel) // when done
        for _, b := range p {
            c.Channel <- b
        }
    }()

    return len(p), nil 
}

func main() {
    cw := NewChannelWriter()
    go func() {
        fmt.Fprintf(cw, "Stream data")
    }()

    for c:= range cw.Channel {
        fmt.Printf("%c\n", c)
    }
}
```

# io Package

## io.Copy


The io.Copy function (and its variants io.CopyBuffer and io.CopyN) make it easy to copy data from an arbitrary io.Reader source into an equally arbitrary io.Writer sink as shown in the following snippet:

```
data := strings.NewReader("Write me down.")
file, _ := os.Create("./iocopy.data")
io.Copy(file, data)
```


## PipeReader and PipeWriter
The io package includes the PipeReader and PipeWriter types that
model IO operations as an in-memory pipe.

Note that the pipe writer will block until the reader completely
consumes the pipe content or an error is encountered. Therefore,
both the reader and writer should be wrapped in a goroutine to
avoid deadlocks.

```
file, _ := os.Create("./iopipe.data")
pr, pw := io.Pipe()
go func() {
    fmt.Fprint(pw, "Pipe streaming")
    pw.Close()
}()
wait := make(chan struct{})
go func() {
    io.Copy(file, pr)
    pr.Close()
    close(wait)
}()
<-wait //wait for pr to finish
```

## LimitedReader
As its name suggests, the io.LimitedReader struct is a reader
that reads only N number of bytes from the specified io.Reader.

```
str := strings.NewReader("The quick brown " + "fox jumps over the lazy dog")
limited := &io.LimitedReader{R: str, N: 19}
io.Copy(os.Stdout, limited)
```


The io/ioutil sub-package implements a small number of
functions that provide utilitarian shortcuts to IO primitives such as
file read, directory listing, temp directory creation, and file write.


# Working with Files

## Create and Open Files 
The os.Create function creates a new file with the specified path. If the file already exists, os.Create will overwrite it. The os.Open function, on the other hand, opens an existing file for reading.

```
func main() {
f1, err := os.Open("./file0.go")
if err != nil {
    fmt.Println("Unable to open file:", err)
    os.Exit(1)
}
defer f1.Close()

f2, err := os.Create("./file0.bkp")
if err != nil {
    fmt.Println("Unable to create file:", err)
    os.Exit(1)
}
defer f2.Close()

n, err := io.Copy(f2, f1)
if err != nil {
    fmt.Println("Failed to copy:", err)
    os.Exit(1)
}

fmt.Printf("Copied %d bytes from %s to %s\n", n, f1.Name(), f2.Name())
}

```


The os.OpenFile function take three parameters. 
- The first one is the path of the file, 
- the second parameter is a masked bit-field value to indicate the behavior of the operation (for example, read-only, read-write, truncate, and so on) 
- and the last parameter is a posixcompliant permission value for the file.

```
f1, err := os.OpenFile("./file0.go", os.O_RDONLY, 0666)
if err != nil {
    fmt.Println("Unable to open file:", err)
    os.Exit(1)
}
defer f1.Close()
```
