package event_channel

import "fmt"

type Subscribers map[string]Subscriber

type Channel struct {
	subscribers Subscribers
}

//NewChannel - сложность O(1)
func NewChannel() *Channel {
	return &Channel{
		subscribers: Subscribers{},
	}
}

//Send - сложность O(n)
func (ch *Channel) Send(msg string) { // O(n)
	for _, sub := range ch.subscribers {
		sub.OnReceive(msg)
	}
}

//Subscribe - сложность O(1)
func (ch *Channel) Subscribe(sub Subscriber) { // O(1)
	ch.subscribers[sub.GetID()] = sub
}

//UnSubscribe - сложность O(1)
func (ch *Channel) UnSubscribe(sub Subscriber) error { // O(1)
	id := sub.GetID()
	if _, ok := ch.subscribers[id]; ok {
		sub.OnUnsubscribe("Unsubscribe success")
		delete(ch.subscribers, id)
		return nil
	}
	return fmt.Errorf("can't find user %s", id)
}

//UnSubscribeAll - сложность O(n)
func (ch *Channel) UnSubscribeAll() error {
	for _, sub := range ch.subscribers {
		if err := ch.UnSubscribe(sub); err != nil {
			return fmt.Errorf("unsubscribe error: %s", err)
		}
	}
	return nil
}
