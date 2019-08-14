/* 
Adding instrumentation 

For any microservice, along with logging, instrumentation is vital. The metrics package of
Go Kit records statistics about your service’s runtime behavior: counting the number of jobs
processed, recording the duration of requests after they have finished, and so on. This is
also a middleware that tampers the HTTP requests and collects metrics. To define a
middleware, simply add one more struct, similar to the logging middleware. Metrics are
useless unless we monitor.

*/


package main

import (
	"log"
	"net/http"
	"os"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/narenaryan/encryptService/helpers"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
)

func main() {
	logger := kitlog.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "encryption",
		Subsystem: "my_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "encryption",
		Subsystem: "my_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	var svc helpers.EncryptService
	svc = helpers.EncryptServiceInstance{}
	svc = helpers.LoggingMiddleware{Logger: logger, Next: svc}
	svc = helpers.InstrumentingMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, Next: svc}
	encryptHandler := httptransport.NewServer(helpers.MakeEncryptEndpoint(svc),
		helpers.DecodeEncryptRequest,
		helpers.EncodeResponse)

	decryptHandler := httptransport.NewServer(helpers.MakeDecryptEndpoint(svc),
		helpers.DecodeDecryptRequest,
		helpers.EncodeResponse)

	http.Handle("/encrypt", encryptHandler)
	http.Handle("/decrypt", decryptHandler)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

