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

**insert a record**

```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (1, 'Mark', 7, '1212 E. Lane, Someville, AK, 57483', 43000.00 ,'1992-01-13');
```

**list records in a table**

```
SELECT * FROM <table name>;
```

**insert a record - variations**

* omitted values will have the default value:

```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,BDAY) VALUES (2, 'Marian', 8, '7214 Wonderlust Ave, Lost Lake, KS, 22897', '1989-11-21');
```

* we can use DEFAULT rather leaving a field blank or specifying a value:

```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (3, 'Maxwell', 6, '7215 Jasmine Place, Corinda, CA 98743', 87500.00, DEFAULT);
```

* we can insert multiple rows:

```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (4, 'Jasmine', 5, '983 Star Ave., Brooklyn, NY, 00912 ', 55700.00, '1997-12-13' ), (5, 'Orranda', 9, '745 Hammer Lane, Hammerfield, Texas, 75839', 65350.00 , '1992-12-13');
```


**auto increment key field**

```
CREATE TABLE phonenumbers(
	ID  SERIAL PRIMARY KEY,
	PHONE           TEXT      NOT NULL
);
```
```
INSERT INTO phonenumbers (PHONE) VALUES ( '234-432-5234'), ('543-534-6543'), ('312-123-5432');
```
```
\d phonenumbers
```
```
SELECT * FROM phonenumbers;
```

**join queries**

* Join queries allow us to select records from two or more tables.
* A join query combines columns from one or more tables - it joins a bunch of columns from different tables together.

**cross join**

* A cross join returns the Cartesian product of rows from tables in the join. 
* In other words, it will produce rows which combine each row from the first table with each row from the second table.

```
CREATE TABLE person (
   ID  SERIAL PRIMARY KEY NOT NULL,
   NAME           CHAR(50) NOT NULL
);
```

```
INSERT INTO person (NAME) VALUES ('Shen'), ('Daniel'), ('Juan'), ('Arin'), ('McLeod');
```

```
CREATE TABLE sport (
   ID  SERIAL PRIMARY KEY NOT NULL,
   NAME           CHAR(50) NOT NULL,
   P_ID         INT      references person(ID)
);
```

```
INSERT INTO sport (NAME, P_ID) VALUES ('Surf',1),('Soccer',3),('Ski',3),('Sail',3),('Bike',3);
```

```
SELECT person.NAME, sport.NAME FROM person CROSS JOIN sport;
```

**outer join**

**left outer join**

* A left outer join gives you everything in one table, and also the matching records in another table.
* For tables A and B a left outer join would give you all rows of the "left" table (A), even if the join-condition does not find any matching row in the "right" table (B).
* This means that if the ON clause matches 0 (zero) rows in B (for a given row in A), the join will still return a row in the result (for that row)—but with NULL in each column   from B.

```
SELECT person.NAME, sport.NAME FROM person LEFT OUTER JOIN sport ON person.ID = sport.P_ID;
```

**right outer join**

* A right outer join is like a left outer join, but for the table on the right.

```
INSERT INTO sport (NAME) VALUES ('Squirrel Suit Flying');
```
```
SELECT person.NAME, sport.NAME FROM person RIGHT OUTER JOIN sport ON person.ID = sport.P_ID;
```

**full outer join**

* A full outer join is like running both a left outer join and a right outer join at the same time. 
* It gives you everything from all tables, and matches what matches.

```
SELECT person.NAME, sport.NAME FROM person FULL OUTER JOIN sport ON person.ID = sport.P_ID;
```

**where**

* Adding WHERE to a SQL query allows you to filter results.
```
SELECT * FROM employees WHERE salary > 60000;
```

**and**

```
SELECT * FROM employees WHERE salary > 60000 AND score = 26;
```

**in**

```
SELECT * FROM employees WHERE score IN (25,26);
```

**not**

```
SELECT * FROM employees WHERE score NOT IN (25,26);
```

**between**

```
SELECT * FROM employees WHERE score BETWEEN 23 AND 26;
```

**is not null**

```
SELECT * FROM employees WHERE score IS NOT NULL;
```

**like**

```
SELECT * FROM employees WHERE name LIKE '%an%';
```

**or**

```
SELECT * FROM employees WHERE score <= 24 OR salary < 50000;
```

**limit**

* Limit the number of records returned

```
SELECT * FROM employees LIMIT 4;
```
```
SELECT * FROM employees ORDER BY id LIMIT 4;
```

**update**

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

**order by**

```
SELECT * FROM employees ORDER BY id;
```

**delete**

```
DELETE FROM table
WHERE <condition>;
```
```
SELECT * FROM sport;
```
```
DELETE FROM sport WHERE id = 6;
```

**WARNING: this deletes all records:**
```
DELETE FROM sport;
```

**users & privileges**

* see current user
```
SELECT current_user;
```
* details of users
```
\du
```
* create user
```
CREATE USER james WITH PASSWORD 'password';
```
* grant privileges
* privileges: SELECT, INSERT, UPDATE, DELETE, RULE, ALL
```
GRANT ALL PRIVILEGES ON DATABASE company to james;
```
* revoke privileges

```
REVOKE ALL PRIVILEGES ON DATABASE company from james;
```

* alter

```
ALTER USER james WITH SUPERUSER;
```
```
ALTER USER james WITH NOSUPERUSER;
```

* remove
```
DROP USER james;
```


* Arguments to the SQL function are referenced in the function body using the syntax `$n`:
* `$1` refers to the first argument, `$2` to the second, and so on.
* If an argument is of a composite type, then the dot notation, e.g., `$1.name`, can be used to access attributes of the argument. 
