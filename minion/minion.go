package minion

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
)

type option func(m *minion) error

type minion struct {
	client  *redis.Client
	handles map[string]RedisHandler
}

func New(opts ...option) (*minion, error) {
	m := &minion{}

	for _, opt := range opts {
		if err := opt(m); err != nil {
			return nil, err
		}
	}

	if _, err := m.client.Ping().Result(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *minion) Run() error {

	channels := []string{}

	for ch := range m.handles {
		channels = append(channels, ch)
	}

	if len(channels) < 1 {
		return errors.New("no channels provided")
	}

	pubsub := m.client.Subscribe(channels...)

	_, err := pubsub.Receive()
	if err != nil {
		return err
	}

	ch := pubsub.Channel()

	fmt.Println("waiting for messages")

	for msg := range ch {
		h, ok := m.handles[msg.Channel]
		if !ok {
			// How could this happen??
			fmt.Println("ignoring message from %q without a handler", msg.Channel)
			continue
		}

		go h.Handle(msg.Payload)
	}

	return nil
}

// Minion options

// RedisURL parses a Redis URL and creats a new redis client
func RedisURL(url string) option {
	return func(m *minion) error {
		ro, err := redis.ParseURL(url)
		if err != nil {
			return err
		}

		m.client = redis.NewClient(ro)
		return nil
	}
}

// RedisHandle registers a hanlder of type Handler to operate on
// an incoming message
func RedisHandle(ch string, h RedisHandler) option {
	return func(m *minion) error {
		if m.handles == nil {
			m.handles = map[string]RedisHandler{}
		}

		m.handles[ch] = h

		return nil
	}
}
