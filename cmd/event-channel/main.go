package main

import (
	"log"

	eventChannel "github.com/anton-tars/event-channel"
)

func main() {

	user1 := eventChannel.NewUser("jkulvich")
	user2 := eventChannel.NewUser("vasya")

	ch1 := eventChannel.NewChannel()
	ch1.Subscribe(user1)
	ch1.Subscribe(user2)

	ch2 := eventChannel.NewChannel()
	ch2.Subscribe(user2)

	pub := eventChannel.NewPublisher()
	pub.AddChannel("test", ch1)
	pub.AddChannel("test2", ch2)

	if err := pub.Send("HELLO!", "test"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("HELLO FROM CH2", "test2"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("HI ALL1"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.DeleteChannel("test"); err != nil {
		log.Fatalf("can't delete: %s", err)
	}

	if err := pub.SendAll("HI ALL11"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

}
