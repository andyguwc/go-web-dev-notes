
# Middleware 

Middleware should perform these functions in order:
- process the request befor hitting the handler
- process the handler function
- process the response before giving it to the client 

But in applications with middleware, it passes through a set of stages, like
logging, authentication, session validation, and so on, and then proceeds to the business
logic. This is to filter the wrong requests from interacting with the business logic. 

The most common use cases are:
- Use a logger to log each and every request hitting the REST API
- Validate the session of the user and keep the communication alive
- Authenticate the user, if not identified
- Write custom logic to scrap the request data
- Attach properties to responses while serving the client


In Go, the function signature of the outer function should exactly match the anonymous
function's signature.


# Chaining Handlers
Let us think about a scenario where an API developer only allows the JSON media type
from clients and also needs to send the server time in UTC back to the client for every
request. Using middleware, we can do that. The functions of two middleware are:
- In the first middleware, check whether the content type is JSON. If not, don't allow the request to proceed
- In the second middleware, add a timestamp called Server-Time (UTC) to the response cookie


# Logging Using Gorilla's Handlers
LoggingHandler: For logging in Apache Common Log Format
CompressionHandler: For zipping the responses
RecoveryHandler: For recovering from unexpected panics

```
func mainLogic(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing request!")
	w.Write([]byte("OK"))
	log.Println("Finished processing request")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainLogic)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8000", loggedRouter)
}
```




