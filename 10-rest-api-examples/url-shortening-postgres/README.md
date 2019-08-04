

1. When URL shortening request comes in, store the URL in the database and get the ID of that record inserted.
2. Pass this ID to the client as the API response.
3. Whenever a client loads the shortened URL, it hits our API server.
4. The API server then converts the short URL back to the database ID and fetches
the record from the original URL.
5. Finally, the client can use this URL to redirect to the original site.


We pass a database ID into the ToBase62 algorithm and get a shorter string out

urlshortener
├── main.go
├── models
│ └── models.go
└── utils
└── encodeutils.go


