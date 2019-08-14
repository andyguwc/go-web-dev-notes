
# Using Redis for Caching 
Redis is an in-memory database that can store key/value pairs. It best suits the caching use
cases where we need to store information temporarily but for huge traffic.

Redis provides a way to expire the keys:values stored in it. We can run a scheduler that
updates the Redis whenever the expiration time has passed.

Whenever the same request is given within the next time (before key expiration), just
pull the response out of Redis instead of hitting the servers


## Start Redis Server
$ redis-server // start the redis server 
$ redis-cli // using Redis CLI
127.0.0.1:6379> SET FOO 1
OK
127.0.0.1:6379> GET FOO
"1"
127.0.0.1:6379> 


