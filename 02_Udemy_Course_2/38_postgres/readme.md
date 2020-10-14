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
