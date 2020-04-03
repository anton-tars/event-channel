package event_channel

type Subscriber interface {
	OnReceive(msg string)
	OnUnsubscribe(msg string)
	GetID() string
}

type SubscriberDefault struct{}

func (SubscriberDefault) OnReceive(string) {
	panic("not implemented")
}

func (SubscriberDefault) OnUnsubscribe(string) {
	panic("not implemented")
}

func (SubscriberDefault) GetID(string) {
	panic("not implemented")
}
