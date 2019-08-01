# Testing 

The testing package is used with the go test command which is used on any Go source files that end with _test.go.

If you have a server.go file, you should also have a server_test.go file that contains
all the tests you want to run on the server.go file. The server_test.go file must be in the
same package as the server.go file.

In the test file you create functions with the following form:
func TestXxx(*testing.T) { … }


## testing.T struct 

The testing.T struct has a number of useful functions:
- Log — Similar to fmt.Println; records the text in the error log.
- Logf — Similar to fmt.Printf. It formats its arguments according to the given format and records the text in the error log.
- Fail — Marks the test function as having failed but allows the execution to continue.
- FailNow — Marks the test function as having failed and stops its execution.

If the results aren’t what you expect, you can call any of the Fail, FailNow,
Error, Errorf, Fatal, or Fatalf functions accordingly. The Fail function, as
you’ve probably guessed, tells you that the test case has failed but allows you to continue
the execution of the rest of the test case. The FailNow function is stricter and
exits the test case once it’s encountered.


## skipping test cases

func TestLongRunningTest(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping long-running test in short mode")
    }
    time.Sleep(10 * time.Second)
}

## run tests in parallel

To run the tests in parallel, you
need to call the Parallel function on testing.T as the first statement in the test case.

Now run this in the console:
go test –v –short –parallel 3

The parallel (-parallel) flag indicates that you want to run a maximum of three test
cases in parallel. You’re still using the -short flag because you don’t want to run the
long-running test case


## Benchmarking

Determine performance of a unit of work 
Benchmark test cases are functions of the format

func BenchmarkXxx(*testing.B) {...}

go test -v -cover -short –bench .


## HTTP testing 

Sequence for doing HTTP testing using httptest package
- Create multiplexer
- Attach tested handler to multiplexer
- Create recorder
- Create request
- Send request to tested handler and write to recorder
- Check recorder for results 



## Test Doubles and Dependency Injection
Simulations of objects, structures and functions used during testing 
dependency injection - decouple dependencies 

Implement an interface with all the methods and use create a fake test double 

