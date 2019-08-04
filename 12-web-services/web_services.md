# Web Services
Web Services communicate via HTTP
- SOAP based
- REST based
- XML RPC based 

# SOAP
SOAP-based services are robust, are explicitly described using WSDL
(Web Service Definition Language), and have built-in error handling

SOAP is known to be cumbersome and unnecessarily complex. The SOAP XML
messages can grow to be verbose and difficult to troubleshoot, and you may often
need other tools to manage them.

SOAP is highly structured and heavily defined, and the XML used for the transportation
of the data can be complex. Every operation and input or output of the service
is clearly defined in the WSDL. The WSDL is the contract between the client and the
server, defining what the service provides and how it’s provided.

As you may realize by now, all the data about the message is contained in the envelope.
For SOAP-based web services, this means that the information sent through
HTTP is almost entirely in the SOAP envelope.


# REST

REST-based web services are a lot more flexible. REST isn’t an architecture in itself
but a design philosophy. It doesn’t require XML, and very often REST-based web services
use simpler data formats like JSON, resulting in speedier web services. RESTbased
web services are often much simpler than SOAP-based ones.

Another difference between the two is that SOAP-based web services are functiondriven;
REST-based web services are data-driven

## Action on the Resource
REST doesn’t allow you to have arbitrary actions on the resources, and you’re more or less restricted to the list of available HTTP methods

- Convert action to a resource 
You can convert the activate action to a resource activation

To activate a user 
POST /user/456/activation HTTP/1.1
{ "date": "2015-05-15T13:05:05Z" }
This code will create an activation resource that represents the activation state of the
user. Doing this also gives the added advantage of giving the activation resource additional
properties.

- Make the action a property of the resource
If activation is a simple state of the customer account, you can simply make the action
a property of the resource, and then use the PATCH HTTP method to do a partial update to the resource. 

For example, you can do this:
PATCH /user/456 HTTP/1.1
{ "active" : "true" }

# Protocol Buffers
Protocol buffers are a flexible, efficient, automated mechanism for serializing structured
data – think XML, but smaller, faster, and simpler. You define how you want your data to
be structured once, then you can use special generated source code to easily write and read
your structured data to and from a variety of data streams and using a variety of
languages. You can even update your data structure without breaking deployed programs
that are compiled against the "old" format.

Protocol buffers have many advantages over JSON/XML for serializing structured data,
such as:
They are simpler
They are 3 to 10 times smaller
They are 20 to 100 times faster
They are less ambiguous
They generate data access classes that are easier to use programmatically

## Syntax
syntax 'proto3';
message NetworkInterface {
int index = 1;
int mtu = 2;
string name = 3;
string hardwareaddr = 4;
}


## Compiling Protocol Buffer with Protoc
1. Install the protoc command-line tool and the proto library.
2. Write a protobuf file with the .proto extension.
3. Compile it to target a programming language (here, it is Go).
4. Import structs from the generated target file and serialize the data.
5. On a remote machine, receive the serialized data and decode it into a struct or
class.

Use this command to compile protocol buffer files to Go files 
protoc --go_out=. *.proto


# GRPC

GRPC is a transport mechanism that sends and receives messages between two systems.
These two systems are traditionally a server and a client. As we described in the previous
chapters, RPC can be implemented in Go for transferring JSON. We called it a JSON RPC
service. Similarly, Google RPC is specially designed to transfer data in the form of protocol
buffers.

GRPC has the following benefits over traditional HTTP/REST/JSON architecture:
- GRPC uses HTTP/2, which is a binary protocol
- Header compression is possible in HTTP/2, which means less overhead
- We can multiplex many requests on one connection
- Usage of protobufs for strict typing of data
- Streaming of requests or responses is possible instead of request/response transactions

