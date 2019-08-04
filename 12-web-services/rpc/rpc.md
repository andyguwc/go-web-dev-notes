
# RPC
Remote Procedure Call (RPC) is an interprocess communication that exchanges information between various distributed systems. Without implementing the functionality locally, we can request things from a network that lies in another place or geographical region.

The entire process can be broken down into the following steps:
- Clients prepare function name and arguments to send
- Clients send them to an RPC server by dialing the connection
- The server receives the function name and arguments
- The server executes the remote process
- The message will be sent back to the client
- The client collects the data from the request and uses it appropriately


# Creating RPC
The RPC server and RPC client should agree upon two things:
1. Arguments passed
2. Value returned

