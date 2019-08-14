# Microservices

Microservices bring the following benefits to the plate:
If the team is big, people can work on chunks of applications
Adaptability is easy for the new developers
Adopting best practices, such as Continuous Integration (CI) and Continuous
Delivery (CD)
Easily replaceable software with loosely coupled architecture

In a monolith application (traditional application), a single huge server serves the incoming
requests by multiplexing the computing power. It is good because we have everything, such
as an application server, database, and other things, in a single place. It also has
disadvantages. When a piece of software breaks, everything breaks. Also, developers need
to set up an entire application to develop a small piece

- Tightly coupled architecture
- Single point of failure
- Velocity of adding new features and components
- Fragmentation of work is limited to teams
- Continuous deployment is very tough because an entire application needs to be pushed

Microservices also create a platform that allows us to use containers (docker, and so on). In
microservices, orchestration and service discovery are very important to track the loosely
coupled elements. A tool such as Kubernetes is used to manage the docker containers.
Generally, it is a good practice to have a docker container for a microservice. Service
discovery is the automatic detection of an IP address and other details on the fly. This
removes the potential threat of hardcoding the stuff that is needed for microservices to
consult each other.

## Go Kit - a package for building microservices

Transport layer: This takes care of transferring data from one service to another
Endpoint layer: This takes care of building endpoints for the given services
Service layer: This is the actual business logic for the API handlers

To create a service, we need to design a few things upfront. They are:
Service implementation
Endpoints
Request/response models
Transport


import (
    "context"
)

// EncryptService is a blueprint for our service
    type EncryptService interface {
    Encrypt(context.Context, string, string) (string, error)
    Decrypt(context.Context, string, string) (string, error)
}

// The service needs to implement these functions to satisfy the interface. Next, create models for your services. Models specify what data a service can receive and produce back.


We are creating a struct called EncyptionServiceInstance that has two methods, Encrypt
and Decrypt. So it satisfies the preceding interface



