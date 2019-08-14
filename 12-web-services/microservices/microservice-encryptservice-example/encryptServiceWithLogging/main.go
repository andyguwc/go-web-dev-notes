/* 
Add transport level and application level logging to the Go Kit microservices

Add a logging middleware - wraps the logs for incoming requests


We need to maintain the order in which the log should print. After logging our request
details, we make sure to allow the request to go to the next middleware/handler using
this function. Next is of the type EncryptService, which is our actual implementation:
mw.Next.(Encrypt/Decrypt)


*/


package main

import (
	"log"
	"net/http"
	"os"

	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/narenaryan/encryptService/helpers"
)

func main() {
	logger := kitlog.NewLogfmtLogger(os.Stderr)
	var svc helpers.EncryptService
	svc = helpers.EncryptServiceInstance{}
	svc = helpers.LoggingMiddleware{Logger: logger, Next: svc}
	encryptHandler := httptransport.NewServer(helpers.MakeEncryptEndpoint(svc),
		helpers.DecodeEncryptRequest,
		helpers.EncodeResponse)

	decryptHandler := httptransport.NewServer(helpers.MakeDecryptEndpoint(svc),
		helpers.DecodeDecryptRequest,
		helpers.EncodeResponse)

	http.Handle("/encrypt", encryptHandler)
	http.Handle("/decrypt", decryptHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
