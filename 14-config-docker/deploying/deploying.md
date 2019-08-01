# Deploy to Heroku

Heroku’s premise is simple and requires only a couple of things:
- A configuration file or mechanism that defines the dependencies. For example, in Ruby this would be a Gemfile file, in Node.js a package.json file, and in Java a pom.xml file.
- A Procfile that defines what to be run. More than one executable can be run at the same time.


How to deploy services
- Change code - get port from environment variable
- Use godep for dependencies 
- Create Heroku application
- Push code 

Once godep is installed, you need to use it to bring in your dependencies. In the root
directory of your web service, run this command:
godep save

This command will create a directory named Godeps, retrieve all the dependencies in
your code, and copy their source code into the directory Godeps/_workspace. It will
also create a file named Godeps.json that lists all your dependencies

go get github.com/tools/godep

In the Procfile
web: ws-h
That’s it! What the listing says is that the web process is associated with the ws-h
executable binary, so that’s what’s going to be executed when the Heroku build
completes.

Push the code to Heroku
heroku create ws-h
git push heroku master