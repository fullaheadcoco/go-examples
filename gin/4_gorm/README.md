# gin example

## run mysql & create database, table

```shell
docker-compose up -d
```

- create databse, table

```shell
mysql> CREATE DATABASE test DEFAULT CHARACTER SET UTF8;
```

## run gin server

```shell
go run main.go
```

## test
- using `test/1.http` in goland.


