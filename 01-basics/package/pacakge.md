# Go Package

A Go package is both a physical and a logical unit of code organization used to encapsulate related concepts that can be reused. By convention, a group of source files stored in the same directory are considered to be part of the same package. 

While not a requirement, it is a recommended convention to set a package's name, in each source file, to match the name of the directory where the file is located. 


## Multi File Packages
The logical content of a package (source code elements such as types, functions, variables, and constants) can physically scale across multiple Go source files. A package directory can contain one or more Go source files.


## Building Packages

go build ./ch06/volt

go build .

Build all packages and sub packages
go build ./.. 

// use flag o to control the name of the binary
$ go build -o ohms

// go build vs go install 
If the package is package main, go build will place the resulting executable in the current directory. go install will put the executable in $GOPATH/bin (using the first element of $GOPATH, if you have more than one).

If the package is not package main, go install and go build will compile the package, displaying any errors encountered





## Package Member Visibility

The usefulness of a package is its ability to expose its source elements to other packages. Controlling the visibility of elements of a package is simple and follows this rule: capitalized identifiers are exported automatically.


## Package Identifiers
Following the format described earlier, the name identifier is placed before the import path as shown in the preceding snippet. A named package can be used as a way to shorten or customize the name of a package. For instance, in a large source file with numerous usage of a certain package, this can be a welcome feature to reduce keystrokes.

Assigning a name to a package is also a way to avoid package identifier collisions in a given source file. It is conceivable to import two or more packages, with different import paths, that resolve to the same package names. As an example, you may need to log information with two different logging systems from different libraries, as illustrated in the following code snippet


## Blank Identifiers
A common idiom for the blank identifier is to load packages for their side effects. This relies on the initialization sequence of packages when they are imported. Using the blank identifier will cause an imported package to be initialized even when none of its members can referenced. This is used in contexts where the code is needed to silently run certain initialization sequences.


## Accessing program arguments
When a program is executed, the Go runtime makes all command-line arguments available
as a slice via package variable os.Args.

package main
import (
    "fmt"
    "os"
)
func main() {
    for _, arg := range os.Args {
        fmt.Println(arg)
    }
}




# Go workspace
bin: auto-generated directory that stores compiled go executable artifacts

pkg: store package artifacts

src: user created directory where go source code files are stored 


