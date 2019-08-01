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

