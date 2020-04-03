package event_channel

import "fmt"

type Subscribers map[string]Subscriber

type Channel struct {
	subscribers Subscribers
}

func NewChannel() *Channel {
	return &Channel{
		subscribers: Subscribers{},
	}
}

func (ch *Channel) Send(msg string) {
	for _, sub := range ch.subscribers {
		sub.OnReceive(msg)
	}
}

func (ch *Channel) Subscribe(sub Subscriber) {
	ch.subscribers[sub.GetID()] = sub
}

func (ch *Channel) UnSubscribe(sub Subscriber) error {
	id := sub.GetID()
	if _, ok := ch.subscribers[id]; ok {
		sub.OnUnsubscribe("Unsubscribe success")
		delete(ch.subscribers, id)
		return nil
	}
	return fmt.Errorf("can't find user %s", id)
}

func (ch *Channel) UnSubscribeAll() error {
	for _, sub := range ch.subscribers {
		if err := ch.UnSubscribe(sub); err != nil {
			return fmt.Errorf("unsubscribe error: %s", err)
		}
	}
	return nil
}
