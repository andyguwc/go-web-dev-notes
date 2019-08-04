# Common Postgres terminal commands

https://github.com/GoesToEleven/golang-web-dev/blob/master/044_postgres/README.md


# Getting data from Postgres

1 Connect to the database using the database connection pool.
2 Send an SQL query to the database, which will return one or more rows.
3 Create a struct.
4 Iterate through the rows and scan them into the struct.


# Postgres Commands
## install
[postgresql website](https://www.postgresql.org/download/)

## log in

```
psql
```

or 
```
/Applications/Postgres.app/Contents/Versions/9.6/bin/psql -p5432
```

## list databases
```
\l
```

## log out
```
\q
```

#  create database
```
CREATE DATABASE employees;
```

## list databases
```
\l
```

## connect to a database
```
\c <database name>
```

## switch back to postgres database
```
\c postgres
```

## see current user
```
SELECT current_user;
```

## see current database
```
SELECT current_database();
```

## drop (remove, delete) database
```
DROP DATABASE <database name>;
```

# create table
```
CREATE TABLE employees (
   ID INT PRIMARY KEY     NOT NULL,
   NAME           TEXT    NOT NULL,
   RANK           INT     NOT NULL,
   ADDRESS        CHAR(50),
   SALARY         REAL DEFAULT 25500.00,
   BDAY			  DATE DEFAULT '1900-01-01'
);
```

## show tables in a database (list down)
```
\d
```

## show details of a table
```
\d <table name>
```

## drop a table
```
DROP TABLE <table name>;
```

## insert a record
```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (1, 'Mark', 7, '1212 E. Lane, Someville, AK, 57483', 43000.00 ,'1992-01-13');
```

## list records in a table
```
SELECT * FROM <table name>;
```

## insert a record - variations
omitted values will have the [default value](https://www.postgresql.org/docs/9.3/static/ddl-default.html):
```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,BDAY) VALUES (2, 'Marian', 8, '7214 Wonderlust Ave, Lost Lake, KS, 22897', '1989-11-21');
```

we can use DEFAULT rather leaving a field blank or specifying a value:
```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (3, 'Maxwell', 6, '7215 Jasmine Place, Corinda, CA 98743', 87500.00, DEFAULT);
```

we can insert multiple rows:
```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (4, 'Jasmine', 5, '983 Star Ave., Brooklyn, NY, 00912 ', 55700.00, '1997-12-13' ), (5, 'Orranda', 9, '745 Hammer Lane, Hammerfield, Texas, 75839', 65350.00 , '1992-12-13');
```

## update

syntax
```
UPDATE table
SET col1 = val1, col2 = val2, ..., colN = valN
WHERE <condition>;
```

```
SELECT * FROM employees;
```

```
UPDATE employees SET score = 99 WHERE ID = 3;
```

## order by
```
SELECT * FROM employees ORDER BY id;
```

## users & privileges

## see current user
```
SELECT current_user;
```

## details of users
```
\du
```

## create user
```
CREATE USER james WITH PASSWORD 'password';
```
CREATE ROLE naren with LOGIN PASSWORD 'passme123';

## grant privileges
privileges: SELECT, INSERT, UPDATE, DELETE, RULE, ALL
```
GRANT ALL PRIVILEGES ON DATABASE company to james;
```

## revoke privileges
```
REVOKE ALL PRIVILEGES ON DATABASE company from james;
```

## alter
```
ALTER USER james WITH SUPERUSER;
```

```
ALTER USER james WITH NOSUPERUSER;
```

## remove
```
DROP USER james;
```


