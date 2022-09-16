# xm-golang-exercise

XM Golang Exercise - v21.0.0

## Prerequisities

- pgAdmin 14 version

## How to run

Create table inm pgAdmin using script in query tools :
```
CREATE TABLE Company
(
	id serial NOT NULL PRIMARY KEY,
	name varchar(255) NOT NULL,
	code varchar(255) NOT NULL,
	website varchar(255),
	phone varchar(255) NOT NULL,
	country varchar(255) NOT NULL
)
```
Use values of database in config file instead of Envs below

| Enviroment Variable | Description | Default  |
| ------------- |:-------------| :-----|
| `LOCALHOST:PORT` | PostgreSQL database connections string with port | "" |
| `DB_USER` | PostgreSQL database username |   "" |
| `DB_PASSWORD` | PostgreSQL database password | "" |
| `DB_NAME` | PostgreSQL database name | "" |


```bash
go run main.go
```
