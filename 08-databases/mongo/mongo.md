
# MongoDB Intro
MongoDB is a popular NoSQL database that is attracting a lot of developers worldwide. It
is different from traditional relational databases such as MySQL, PostgreSQL, and SQLite3.
The main big difference of MongoDB compared to other databases is the ease of scalability
at the time of internet traffic. It also has JSON as its data model, which allows us to store
JSON directly into the database.

MongoDB stores data in a document; think of this as a row in SQL databases. All MongoDB
documents are stored in a collection, and the collection is a table (in SQL analogy). A
sample document for an IMDB movie looks like this:
{
    _id: 5,
    name: 'Star Trek',
    year: 2009,
    directors: ['J.J. Abrams'],
    writers: ['Roberto Orci', 'Alex Kurtzman'],
    boxOffice: {
    budget:150000000,
    gross:257704099
    }
}

Advantages of MongoDB
- Easy to model (schema free)
- Can leverage querying power
- Document structure suits modern-day web applications (JSON)
- More scalable than relational databases

## Installing MongoDB 
For installing MongoDB on macOS X, use the Homebrew software. We can easily install it
using the following command:
$ brew install mongodb

After that, we need to create the db directory where MongoDB stores its database:
$ mkdir -p /data/db

Then, change the permissions of that file using chown:
$ chown -R `id -un` /data/db
Now we have MongoDB ready. We can run it in a terminal window with the following
command, which starts the MongoDB daemon:
$ mongod

## Mongo Shell
// launch mongo shell 
> mongo 
// show databases
> show databases 
// create a new database
> use ab_name 

// create a collection called movies and insert a document 
// The JSON you inserted has an ID called _id. We can either provide it while inserting a document or MongoDB can insert one for you itself.

> db.movies.insertOne({ _id: 5, name: 'Star Trek', year: 2009, directors:
    ['J.J. Abrams'], writers: ['Roberto Orci', 'Alex Kurtzman'], boxOffice: {
    budget:150000000, gross:257704099 } } )
    {
    "acknowledged" : true,
    "insertedId" : 5
    }

// see all documents in the collection 
> db.movies.find()

// filter for specific records 
> db.movies.find({year: {$eq: 2008}})
> db.movies.find({$or: [{'boxOffice.budget': {$gt: 150000000}}, {year:2009}]})

// delete records 
> db.movies.deleteOne({"_id": ObjectId("59574125bf7a73d140d5ba4a")})


# mgo - MongoDB driver for Go

