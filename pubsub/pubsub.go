package pubsub

import (
	"github.com/gorilla/websocket"
	"fmt"
)

type PubSub struct {
	Clients       [] Client
	Subscriptions [] Subscription
}

type Client struct {
	Id   string
	Conn *websocket.Conn
}

type Message struct {
	Topic   string
	Action  string
	Message string
}

type Subscription struct {
	Topic  string
	Client *Client
}

func (p *PubSub) AddClient(client Client) (*PubSub) {

	clients := append(p.Clients, client)

	p.Clients = clients

	return p
}

func (p *PubSub) HandleReceivedMessage(client Client, messageType int, message []byte) (*PubSub) {

	fmt.Printf("Received message %d %s %s", messageType, message, client.Id)

	for _, c := range p.Clients {

		c.Conn.WriteMessage(messageType, message)
	}

	return p
}

func (p *PubSub) GetSubscriptions(topic string, client *Client) ([]Subscription) {

	var s []Subscription

	for _, sub := range p.Subscriptions {

		if client != nil {
			if sub.Client.Id == client.Id && sub.Topic == topic {
				s = append(s, sub)
			}
		} else {
			if sub.Topic == topic {
				s = append(s, sub)
			}
		}

	}

	return s
}

func (p *PubSub) Subscribe(topic string, client *Client) (*PubSub) {

	subs := p.GetSubscriptions(topic, client)

	if len(subs) > 0 {
		return p
	}

	newSubscription := Subscription{
		Client: client,
		Topic:  topic,
	}

	p.Subscriptions = append(p.Subscriptions, newSubscription)

	return p

}

func (p *PubSub) publish(topic string, message []byte, excludeClient *Client) (*PubSub) {

	subscriptions := p.GetSubscriptions(topic, nil)

	for _, sub := range subscriptions {

		if excludeClient != nil && sub.Client.Id != excludeClient.Id {

			// send message
			sub.Client.send(1, message)

		} else {
			sub.Client.send(1, message)
		}
	}

	return p
}

func (c *Client) send(messageType int, message [] byte) (error) {

	return c.Conn.WriteMessage(messageType, message)
}
