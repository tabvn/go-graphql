# go-graphql
Create GraphQL Api Service + Realtime PubSub use Golang

## Installation

* MYSQL (currently use MYSQL database)
* Change database connection in /config/config.go
* Import schema.sql to database
```
mysql -u root -p YOUR-DATABASE-NAME < schema.sql
```

# start Server 

```
go run main.go

```

