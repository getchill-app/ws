package api

import (
	"github.com/keys-pub/keys"
)

// EventPubSub is the pub/sub name for events.
const EventPubSub = "e"

// Event to client.
// JSON is used for websocket clients.
type Event struct {
	// Type describes which field is set
	// - vault: Vault
	// - accountCreated: AccountCreated
	Type string `json:"type"`

	Vault          *VaultEvent     `json:"vault,omitempty" msgpack:"v,omitempty"`
	AccountCreated *AccountCreated `json:"accountCreated,omitempty" msgpack:"ac,omitempty"`
}

type VaultEvent struct {
	KID   keys.ID `json:"kid" msgpack:"k"`
	Index int64   `json:"idx" msgpack:"i"`
	Token string  `json:"token" msgpack:"t"`
}

type AccountCreated struct {
	KID keys.ID `json:"kid"`
}
