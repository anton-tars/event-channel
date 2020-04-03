package event_channel

import (
	"fmt"
)

type Channels map[string]*Channel

type Publisher struct {
	channels Channels
}

//NewPublisher - create new publisher
func NewPublisher() *Publisher {
	return &Publisher{
		channels: Channels{},
	}
}

//AddChannel - add channel
func (p *Publisher) AddChannel(name string, channel *Channel) error {
	_, ok := p.channels[name]
	if ok {
		return fmt.Errorf("channel %s exist", name)
	}
	p.channels[name] = channel
	return nil
}

//DeleteChannel - delete channel
func (p *Publisher) DeleteChannel(name string) error {
	channel, ok := p.channels[name]
	if !ok {
		return fmt.Errorf("channel %s can't be found", name)
	}
	channel.UnSubscribeAll()
	delete(p.channels, name)
	return nil
}

//TODO: Удаление канала и получение списка каналов

//Send - send message to channel or channels
func (p *Publisher) Send(msg string, channels ...string) error {
	if len(channels) == 0 {
		p.SendAll(msg)
		return nil
	}
	for _, ch := range channels {
		channel, ok := p.channels[ch]
		if !ok {
			return fmt.Errorf("channel %s can't be found", ch)
		}
		channel.Send(msg)
	}
	return nil
}

//SendAll - send message to all channels
func (p *Publisher) SendAll(msg string) error {
	for _, ch := range p.channels {
		ch.Send(msg)
	}
	return nil
}
