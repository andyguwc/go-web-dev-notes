# REST Services

## Characteristics with REST

Client-server based architecture: This architecture is most essential for the
modern web to communicate over HTTP. A single client-server may look naive
initially, but many hybrid architectures are evolving. 

Stateless: This is the most important characteristic of a REST service. A REST
HTTP request consists of all the data needed by the server to understand and
give back the response. Once a request is served, the server doesn't remember if
the request has arrived after a while. So the operation will be a stateless one.

Cacheable: Many developers think a technology stack is blocking their web
application or API. But in reality, their architecture is the reason. The database
can be a potential tuning piece in a web application.

Scripts on demand: Have you ever designed a REST service which serves the
JavaScript files and you execute them on the fly? This code on demand is also the
main characteristic REST can provide. It is more common to request scripts and
data from the server.

Multiple layered system: The REST API can be served from multiple servers.
One server can request the other, and so forth. So when a request comes from the
client, request and response can be passed between many servers to finally
supply a response back to the client.

Representation of resources: The REST API provides the uniform interface to
talk to. It uses a Uniform Resource Identifier (URI) to map the resources (data).
It also has the advantage of requesting a specific data format as the response.

## REST Verbs

REST Verb Action Success Failure
GET Fetches a record or set of resources from the server 200 404
OPTIONS Fetches all available REST operations 200 -
POST Creates a new set of resources or a resource 201 404, 409
PUT Updates or replaces the given record 200, 204 404
PATCH Modifies the given record 200, 204 404
DELETE Deletes the given resource 200 404

**GET**
A GET method fetches the given resource from the server. 
To specify a resource, GET uses a few types of URI queries:
- Query parameters: add detailed info to identify a resource
/v1/books/?category=fiction&publish_date=2017

- Path-based parameters
/v1/payments/billing-agreements/agreement_id

Query parameters are used to fetch multiple resources based on the query parameters. 
If a client needs a single resource with exact URI information, it can use Path parameters to specify the resource.


**POST vs. PUT vs. PATCH**
POST is used to create a resource on the server

The POST request can update multiple resources: /v1/books.
The POST request has a body like this:
{"name" : "Lord of the rings", "year": 1954, "author" : "J. R. R. Tolkien"}

PUT updates a single resource that already exists
/v1/books/1256
with body that is JSON like this:
{"name" : "Lord of the rings", "year": 1955, "author" : "J. R. R. Tolkien"}

The PATCH method is similar to PUT, except it won't replace the whole record. PATCH patches the column that is being modified. Let us update the book 1256 with a new column called ISBN:
/v1/books/1256
with the JSON body like this:
{"isbn" : "0618640150"}


## Status Codes
- 2xx (successful)
200 and 201 fall under the success family. They indicate that an operation was successful.
200 (Successful Operation) is the most common type of response status code in REST
201 (Successfully Created) is returned when a POST operation successfully creates a resource on the server
204 (No content) is issued when a client needs a status but not any data back

- 3xx (redirection)
301 is issued when a resource is moved permanently to a new URL endpoint. It is
essential when an old API is deprecated
The 304 status code indicates that content is cached and no modification
happened for the resource on the server

- 4xx (client error)
400 (Bad Request) is returned when the server cannot understand the client
request.
401 (Unauthorized) is returned when the client is not sending the authorization
information in the header.
403 (Forbidden) is returned when the client has no access to a certain type of
resources.
404 (Not Found) is returned when the client request is on a resource that is
nonexisting.
405 (Method Not Allowed) is returned if the server bans a few methods on
resources. GET and HEAD are exceptions.

- 5xx (server error)
500 (Internal Server Error) status code gives the development error which is
caused by some buggy code or some unexpected condition
501 (Not Implemented) is returned when the server is no longer supporting the
method on a resource
502 (Bad Gateway) is returned when the server itself got an error response from
another service vendor
503 (Service Unavailable) is returned when the server is down due to multiple
reasons, like a heavy load or for maintenance
504 (Gateway Timeout) is returned when the server is waiting a long time for a
response from another vendor and is taking too much time to serve the client

## Single Page Applications 
Consider front end talking to the backend only using the REST API


In the SPA, the flow is quite different:
1. Request the HTML template/s to the browser in one single go.
2. Then, query the JSON REST API to fill a model (data object).
3. Adjust the UI according to the data in the model (JSON).
4. When users modify the UI, the model (data object) should change automatically.
For example, in AngularJS, it is possible with two-way data binding. Finally,
make REST API calls to notify the server about changes whenever you want.
In this way, communication happens only in the form of the REST API. The client takes care
of logically representing the data. This causes systems to move from Response Oriented
Architecture (ROA) to Service Oriented Architecture (SOA).



# Using Frameworks 

## Steps

1. Design a REST API document.

HTTP verb Path Action Resource
POST /v1/train (details as JSON body) Create Train
POST /v1/station (details as JSON body) Create Station
GET /v1/train/id Read Train
GET /v1/station/id Read Station
POST /v1/schedule (source and destination) Create Route


2. Create models for a database.

folder structure 
- dbutils
    - init-tables.go
    - models.go
- railAPI
    - main.go
- railapi.db

```
type TrainResource struct {
    ID int
    DriverName string
    OperatingStatus bool
}
// StationResource holds information about locations
type StationResource struct {
    ID int
    Name string
    OpeningTime time.Time
    ClosingTime time.Time
}
// ScheduleResource links both trains and stations
type ScheduleResource struct {
    ID int
    TrainID int
    StationID int
    ArrivalTime time.Time
}
```

3. Implement the API logic.


