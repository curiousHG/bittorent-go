package handshake

import (
	"fmt"
	"io"
)

type Handshake struct {
	Pstr string
	InfoHash [20]byte
	PeerID [20]byte
}

func (h *Handshake) Serialize() []byte {
	
}