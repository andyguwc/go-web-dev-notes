# Testing 

The testing package is used with the go test command which is used on any Go source files that end with _test.go.


## Test conventions

Test file names

Within a specified package, the test tool will compile all
files with the *_test.go name pattern.

Traditionally, test files are kept in the same package (directory) as the code being tested.

If you have a server.go file, you should also have a server_test.go file that contains all the tests you want to run on the server.go file. The server_test.go file must be in the same package as the server.go file.

In the test file you create functions with the following form:
func TestXxx(*testing.T) { … }


$ go test ./...

// specifying a subpacakge to run tests
$ go test ./vector 

// execute functions whose name match the specified expression 
$ go test -run=VectorAdd -v

// test coverage 
go test -cover

## testing.T struct 

func Test<Name>(*testing.T)

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

It is sometimes necessary to skip test functions due to a number of factors such as
environment constraints, resource availability, or inappropriate environment settings. The testing API makes it possible to skip a test function using the SkipNow method from type testing.T

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


Another useful aspect of benchmarking code is to compare the performance of different
algorithms that implement similar functionalities. Exercising the algorithms using
performance benchmarks will indicate which of the implementations may be more
compute- and memory-efficient.

## Table Driven Tests

This is where a set of input and expected output is stored in a data structure, which is then used to cycle through different test scenarios.

https://github.com/vladimirvivien/learning-go/blob/master/ch12/vector/vec_test.go

func TestVectorMag(t *testing.T) {
	cases := []struct{
		vec SimpleVector
		expected float64

	}{
    	{New(1.2, 3.4), math.Sqrt(1.2*1.2 + 3.4*3.4)},
		{New(-0.21, 7.47), math.Sqrt(-0.21*-0.21 + 7.47*7.47)},
		{New(1.43, -5.40), math.Sqrt(1.43*1.43 + -5.40*-5.40)},
		{New(-2.07, -9.0), math.Sqrt(-2.07*-2.07 + -9.0*-9.0)},
	}
	for _, c := range cases {
		mag := c.vec.Mag()
		if mag != c.expected {
			t.Errorf("Magnitude failed, execpted %d, got %d", c.expected, mag)
		}
	}
}


## HTTP testing 

Sequence for doing HTTP testing using httptest package
- Create multiplexer
- Attach tested handler to multiplexer
- Create recorder
- Create request
- Send request to tested handler and write to recorder
- Check recorder for results 

Testing server code 
https://github.com/vladimirvivien/learning-go/blob/master/ch12/service/serv_test.go

```
func TestVectorAdd(t *testing.T) {
	reqBody := "[[1,2],[3,4]]"
	req, err := http.NewRequest("POST", "http://0.0.0.0/", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	actual := vector.New(1, 2).Add(vector.New(3, 4))
	w := httptest.NewRecorder()
	add(w, req)
	if actual.String() != strings.TrimSpace(w.Body.String()) {
		t.Fatalf("Expecting actual %s, got %s", actual.String(), w.Body.String())
	}
}
```


Testing client
The test function first sets up the server along with its handler function. Inside the function of http.HandlerFunc, the code first ensures that the client requests the proper path of "/vec/add". Next, the code inspects the request body from the client, ensuring proper JSON format and valid parameters for the add operation. Finally, the handler function encodes the expected result as JSON and sends it as a response to the client.



```
func TestClientAdd(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(resp http.ResponseWriter, req *http.Request) {
			// test incoming request path
			if req.URL.Path != "/vec/add" {
				t.Errorf("unexpected request path %s", req.URL.Path)
				return
			}
			// test incoming params
			body, _ := ioutil.ReadAll(req.Body)
			params := strings.TrimSpace(string(body))
			if params != "[[1,2],[3,4]]" {
				t.Errorf("unexpected params '%v'", params)
				return
			}
			// send result
			result := vector.New(1, 2).Add(vector.New(3, 4))
			err := json.NewEncoder(resp).Encode(&result)
			if err != nil {
				t.Fatal(err)
				return
			}
		},
	))
	defer server.Close()
	client := newVecClient(server.URL)
	expected := vector.New(1, 2).Add(vector.New(3, 4))
	result, err := client.add(vector.New(1, 2), vector.New(3, 4))
	if err != nil {
		t.Fatal(err)
	}
	if !result.Eq(expected) {
		t.Errorf("Expecting %s, got %s", expected, result)
	}
}

```

## Test Doubles and Dependency Injection
Simulations of objects, structures and functions used during testing 
dependency injection - decouple dependencies 

Implement an interface with all the methods and use create a fake test double 



