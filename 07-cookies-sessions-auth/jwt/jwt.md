# JWT 


## JSON Web Tokens Flow
1. The client passes the username/password in a POST request to the login API.
2. The server authenticates the details and if successful, it generates a JWT and returns it back instead of creating a cookie. It is the client's responsibility to store this token.
3. Now, the client has the JWT. It needs to add this in subsequent REST API calls such as GET, POST, PUT, and DELETE in the request headers.
4. Once again, the server checks the JWT and if it is successfully decoded, the server sends the data back by looking at the username supplied as part of the token.



## How is JWT generated
1. Create a JWT header by doing Base64Url encoding on the header JSON.
2. Create a JWT payload by doing Base64Url encoding on the payload JSON.
3. Create a signature by encrypting the appended header and payload using a secret key.
4. JWT string can be obtained by appending the header, payload, and signature.


## JWT in Go
Creating JWT in Go

For example, it looks like the following code snippet:

```
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": "admin",
    "iat":time.Now().Unix(),
})
```

jwt.SigningMethodHS256 is an encryption algorithm that is available within the package. The second argument is a map with claims such as private (here username) and reserved (issued at). Now we can generate a tokenString using the SignedString function on a token:
tokenString, err := token.SignedString("my_secret_key")

Reading a JWT in Go

```
token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{},error) {
    // key function
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
    return nil, fmt.Errorf("Unexpected signing method: %v",
    token.Header["alg"])
    }
    return "my_secret_key", nil
})

if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    // Use claims for authorization if token is valid
    fmt.Println(claims["username"], claims["iat"])
} else {
    fmt.Println(err)
}
```

## OAuth 2
OAuth 2 is an authentication framework that is used to create authentication pattern between different systems. In this, the client, instead of making a request to the resource server, makes an initial request for some entity called resource owner. This resource owner gives back the authentication grant for the client (if credentials are successful). The client now sends this authentication grant to another entity called an authentication server.


## Authentication versus authorization
Authentication is the process of identifying whether a client is genuine or not. When a server authenticates a client, it checks the username/password pair and creates session cookie/JWT.
Authorization is the process of differentiating one client from another after a successful authentication. In cloud services, the resources requested by a client need to be served by checking that the resources belong to that client but not the other client. The permissions and access to resources vary for different clients. For example, the admin has the highest privileges of resources. A normal user's access is limited.




