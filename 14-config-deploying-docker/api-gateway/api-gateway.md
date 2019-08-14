# API Gateway

Suppose a company named XYZ developed the API for its internal purpose. There are two
ways in which it exposes that API for external use:
- Exposes it using authentication from known clients
- Exposes it as an API as a service

In the first case, this API is consumed by the other services inside the company. Since they
are internal, we don't restrict the access. But in the second case, since API details are given
to the outside world, we need a broker in between to check and validate the requests. This
broker is the API gateway. An API gateway is a broker that sits in between the client and
the server and forwards the request to the server on passing specific conditions.

Things apply to an API
- Authentication
- Loggin or requests and responses

Basically, an API getaway does these things:
- Logging
- Security
- Traffic control
- Transformations


# Implementing Kong (Using Docker)

Create three containers 
- Kong database
- Go container
- Kong application

// create a container along with the user and databse name passed as variables to the container
docker run -d --name kong-database \
-p 5432:5432 \
-e "POSTGRES_USER=kong" \
-e "POSTGRES_DB=kong" \
postgres:latest

// apply migrations required by Kong
docker run --rm \
--link kong-database:kong-database \
-e "KONG_DATABASE=postgres" \
-e "KONG_PG_HOST=kong-database" \
kong:1.0.3 kong migrations bootstrap && kong migrations up \ 
