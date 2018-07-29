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

## Test PubSub From JS Client In Browser console

```javascript

var ws = new WebSocket("ws://127.0.0.1:3001/ws");

ws.onmessage = (msg) => {
	console.log("received server message:", msg.data)
}

// Subscribe to a topic

ws.send('{"action": "subscribe", "topic": "topic-xyz"}')


// Publish a message

ws.send('{"action": "public", "topic": "topic-xyz", "message": "Your message"}')


```


<img src="https://firebasestorage.googleapis.com/v0/b/tabvn-fireshot.appspot.com/o/shots%2FQrC4k82w1uVqSO8ckTnvisBko7l1%2F-LIVDnqNVwxN4hWma2MU.png?alt=media&token=f2ab391b-a23c-47f3-9d9c-1a860e11559f" />
