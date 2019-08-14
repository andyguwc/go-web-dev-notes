/* 

Encrypt micro services using Go Kit

To create a service, we need to design a few things upfront. They are:
- Service implementation
- Endpoints
- Request/response models
- Transport


*/


package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/narenaryan/encryptService/helpers"
)

func main() {
	svc := helpers.EncryptServiceInstance{}
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
