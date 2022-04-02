# gin example

## run mysql & create database, table

```shell
./start-mysql-docker.sh
docker exec -it gin-todo-mysql /bin/bash
mysl -u root -p
```

- create databse, table

```shell
mysql> CREATE DATABASE todos DEFAULT CHARACTER SET UTF8;
mysql> USE todos;
mysql> CREATE TABLE todos
(
id INT NOT NULL AUTO_INCREMENT,
title VARCHAR(32),
description VARCHAR(32),
   PRIMARY KEY(ID)
);
```

## run gin server

```shell
go run main.go
```

## test

- using `test/1.http` in goland.


