**list databases**

```
\l
```
**connect to a database**

```
\c <database name>
```

**switch back to postgres database**

```
\c postgres
```

**log out**

```
\q
```

**create database**

```
CREATE DATABASE employees;
```

**see current user**

```
SELECT current_user;
```

**see current database**

```
SELECT current_database();
```

**drop (remove, delete) database**

```
DROP DATABASE <database name>;
```

**create table**

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

**show tables in a database (list down)**

```
\d
```

**show details of a table**

```
\d <table name>
```

**drop a table**

```
DROP TABLE <table name>;
```
