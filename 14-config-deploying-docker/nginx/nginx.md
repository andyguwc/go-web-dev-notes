# Nginx 

Nginx is a high performant web server and load balancer, and is well suited to deploying
high traffic websites. Even though this decision is opinionated, Python and Node
developers usually use this.

Nginx can also act as an upstream proxy server that allows us to redirect the HTTP requests
to multiple application servers running on the same server. The main contender of Nginx is
Apache's httpd. Nginx is an excellent static file server that can be used by the web clients.
Since we are dealing with APIs, we will look into aspects of dealing with HTTP requests.

## Config
$ brew install nginx

change the config port to 80
$ nano /usr/local/etc/nginx/nginx.conf

$ sudo nginx 
start the server

$ sudo nginx -s stop

$ service nginx restart

// missing sites-available directory https://stackoverflow.com/questions/17413526/nginx-missing-sites-available-directory


## Proxy Server
A proxy server is a server that holds the information of original servers in it. It acts as the
front block for the client request. Whenever a client makes an HTTP request, it can directly go the application server. But, if the application server is written in a programming language, you need a translator that can turn the application response into a clientunderstandable response.


The benefits of having a proxy server (Nginx):
It can act as a load balancer
It can sit in front of cluster of applications and redirect HTTP requests
It can serve a filesystem with a good performance
It streams media very well


## Nginx Paths 

Type Path Description
Configuration /etc/nginx/nginx.con This is the base Nginx configuration file. It can be used as the default file.

Configuration /etc/nginx/sites-available/ If we have multiple sites running within Nginx, we can have multiple configuration files.

Configuration /etc/nginx/sites-enabled/ These are the sites activated currently on Nginx.

Log /var/log/nginx/access.log This log file records the server activity, such as timestamp and API endpoints.

Log /var/log/nginx/error.log This log file logs all proxy server-related errors, such as disk space, file system permissions, and so on.


## Load Balancing

Load balancing
Round Robin Requests are distributed evenly across servers and server weights are taken into consideration.

Least Connection 
Requests are sent to the server that is currently serving the least number of clients.

IP Hash
This is used to send the requests from a given client's IP to the given server. Only when that server is not available is it given to another server.

Least Time 
A request from the client is sent to the machine with the lowest average latency (time to serve client) and least number of active connections.


## Load Balancing Nginx

We now see how load balancing is practically achieved in Nginx for our Go API servers.
The first step in this process is to create an upstream in the http section of the Nginx
configuration file:
http {
upstream cluster {
server site1.mysite.com weight=5;
server site2.mysite.com weight=2;
server backup.mysite.com backup;
}
}
Here, servers are the IP addresses or domain names of the servers running the same code.
We are defining an upstream called backend here. It is a server group that we can refer to
in our location directive.


## Securing the Nginx proxy server

This will be very important for our
REST API servers because, let us suppose we have servers X, Y, and Z that talk to each
other. X can serve clients directly, but X talks to Y and Z for some information by calling an
internal API. Since we know that clients should not access Y or Z, we can make it so that
only X is allowed to access the resources. We can allow or deny the IP addresses using
the nginx access module. It looks like this:
location /api {
    ...
    deny 192.168.1.2;
    allow 192.168.1.1/24;
    allow 127.0.0.1;
    deny all;
}

## Monitoring Go API Server with Supervisord
It is fine that Nginx is sitting in front of our Go API server, it just proxies a port. However, sometimes that web application may stop due to the operating system restarting or
crashing. Whenever your web server gets killed, it is someone's job to automatically bring it
back to life. Supervisord is such a task runner.

Supervisord is a tool that can monitor running processes (system) and
can restart them when they were terminated.

Now, we can ask our supervisorctl to re-read the configuration and restart the task
(process). For that, just say:
supervisorctl reread
supervisorctl update
Then, launch our supervisorctl with:
supervisorctl



