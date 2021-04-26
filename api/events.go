package api

import (
	"github.com/keys-pub/keys"
)

// EventPubSub is the pub/sub name for events.
const EventPubSub = "e"

// Event to client.
// JSON is used for websocket clients.
type Event struct {
	Type  string `json:"type" msgpack:"type"`
	Token string `json:"token,omitempty" msgpack:"token,omitempty"`

	Vault *Vault `json:"vault,omitempty" msgpack:"vault,omitempty"`
}

type Vault struct {
	KID   keys.ID `json:"kid" msgpack:"kid"`
	Index int64   `json:"idx" msgpack:"idx"`
}
