package pubsub

import (
	"github.com/gorilla/websocket"
	"encoding/json"
)

var PUBLISH = "publish"
var SUBSCRIBE = "subscribe"

type PubSub struct {
	Clients       [] Client
	Subscriptions [] Subscription
}

type Client struct {
	Id   string
	Conn *websocket.Conn
}

type Message struct {
	Topic   string          `json:"topic"`
	Action  string          `json:"action"`
	Message json.RawMessage `json:"message"`
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

func (p *PubSub) RemoveClient(client Client) (*PubSub) {

	// remove client's subscriptions
	for i := 0; i < len(p.Subscriptions); i++ {

		if client.Id == p.Subscriptions[i].Client.Id {
			p.Subscriptions = append(p.Subscriptions[:i], p.Subscriptions[i+1:]...)

		}
	}

	// remove client
	for i := 0; i < len(p.Clients); i++ {

		if client.Id == p.Clients[i].Id {
			p.Clients = append(p.Clients[:i], p.Clients[i+1:]...)

		}
	}

	return p
}

func (p *PubSub) HandleReceivedMessage(client *Client, messageType int, message []byte) (*PubSub) {

	var m Message

	if err := json.Unmarshal(message, &m); err != nil {
		// this is not type of PubSub message so we do not do anything.
		return p
	}

	switch m.Action {

	case PUBLISH:

		var payload = map[string]interface{}{
			"topic":   m.Topic,
			"message": m.Message,
		}

		sendPayload, err := json.Marshal(payload)

		if err != nil {
			break
		}

		p.publish(m.Topic, sendPayload, nil)

		break

	case SUBSCRIBE:

		p.Subscribe(m.Topic, client)

		break

	default:
		break
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

		if excludeClient != nil {

			if sub.Client.Id != excludeClient.Id {
				sub.Client.send(1, message)
			}

		} else {
			sub.Client.send(1, message)
		}
	}

	return p
}

func (c *Client) send(messageType int, message [] byte) (error) {

	return c.Conn.WriteMessage(messageType, message)
}
