package api

import (
	"github.com/keys-pub/keys"
)

// EventPubSub is the pub/sub name for events.
const EventPubSub = "e"

type EventType string

const (
	ChannelType  EventType = "channel"
	VaultType    EventType = "vault"
	ChannelsType EventType = "channels"
)

// Event to client.
// JSON is used for websocket clients.
type Event struct {
	Type  EventType `json:"type" msgpack:"type"`
	Token string    `json:"token,omitempty" msgpack:"token,omitempty"`

	Channel *Channel `json:"channel,omitempty" msgpack:"channel,omitempty"`
	Vault   *Vault   `json:"vault,omitempty" msgpack:"vault,omitempty"`
}

type Channel struct {
	ID    keys.ID `json:"id" msgpack:"id"`
	Index int64   `json:"idx" msgpack:"idx"`
}

type Vault struct {
	ID    keys.ID `json:"id" msgpack:"id"`
	Index int64   `json:"idx" msgpack:"idx"`
}
