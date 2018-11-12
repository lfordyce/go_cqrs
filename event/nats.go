package event

import (
	"bytes"
	"encoding/gob"
	"github.com/lfordyce/hero_cqrs/schema"
	"github.com/nats-io/go-nats"
)
type NatsEventStore struct {
	nc *nats.Conn
	heroCreatedSubscription *nats.Subscription
	heroCreatedChan chan HeroCreatedMessage
}

func NewNats(url string) (*NatsEventStore, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsEventStore{nc: nc}, nil
}

func (e *NatsEventStore) Close() {
	if e.nc != nil {
		e.nc.Close()
	}
	if e.heroCreatedSubscription != nil {
		e.heroCreatedSubscription.Unsubscribe()
	}
	close(e.heroCreatedChan)
}

func (e *NatsEventStore) PublishHeroCreated(h schema.Hero) error {
	message := HeroCreatedMessage{h.ID, h.Body, h.CreatedAt}
	data, err := e.writeMessage(&message)
	if err != nil {
		return err
	}
	return e.nc.Publish(message.Key(), data)
}

func (mq *NatsEventStore) writeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (mq *NatsEventStore) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}

func (e *NatsEventStore) SubscribeHeroCreated() (<- chan HeroCreatedMessage, error) {
	m := HeroCreatedMessage{}
	e.heroCreatedChan = make(chan HeroCreatedMessage, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	if err != nil {
		return nil, err
	}
	// Decode Message
	go func() {
		for {
			select {
			case msg := <- ch:
				e.readMessage(msg.Data, &m)
				e.heroCreatedChan <- m
			}
		}
	}()
	return (<- chan HeroCreatedMessage)(e.heroCreatedChan), nil
}

func (e *NatsEventStore) OnHeroCreated(f func(HeroCreatedMessage)) (err error) {
	m := HeroCreatedMessage{}
	e.heroCreatedSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		e.readMessage(msg.Data, &m)
		f(m)
	})
	return
}
