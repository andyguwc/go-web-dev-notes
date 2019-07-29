# ChitChat

This is the simple forum application written with Go for the book "Go Web Programming" from Manning Publication Co. The source code for this application is the basis for the Chapter 2 - Go ChitChat. 

However the code is a reference and is not a 1 - 1 match with the code in the book. There are portions of the code that is more detailed in here than in Chapter 2 (which is a simplified version of the source code here).

Some differences include:

* This version of ChitChat is configurable with a `config.json` file
* This version is able to log to a `chitchat.log` file
* There are test files in this code
* All structs are fully fleshed out (in the book chapter, they are only implied)
* Some of the functions are placed as methods for the struct types instead of being a part of the package


# Run the Code
- first make sure to create the chitchat database and create tables (in the setup.sql)

- then run the code 
    go build
    ./chitchat

- check PORT 8080 and create a user to engage with the app 

# The High Level Flow
1 The client sends a request to the server.
2 This is received by the multiplexer, which redirects it to the correct handler.
3 The handler processes the request.
4 When data is needed, it will use one or more data structs that model the data in the database.
5 The model connects with the database, triggered by functions or methods on the data struct.
6 When processing is complete, the handler triggers the template engine, sometimes sending in data from the model.
7 The template engine parses template files to create templates, which in turn are combined with data to produce HTML.
8 The generated HTML is sent back to the client as part of the response.

