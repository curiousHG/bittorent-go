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
	buf := make([]byte, len(h.Pstr)+49)
	buf[0] = byte(len(h.Pstr))
	curr := 1
	curr += copy(buf[curr:], h.Pstr)
	curr += copy(buf[curr:], make([]byte, 8))
	curr += copy(buf[curr:], h.InfoHash[:])
	curr += copy(buf[curr:], h.PeerID[:])
	return buf
}

// Parses a handshake from stream
func Read(r io.Reader) (*Handshake, error) {
	buf := make([]byte, 68)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}
	pstrLen := int(buf[0])
	if pstrLen != 19 {
		return nil, fmt.Errorf("pstrLen is %d, expected 19", pstrLen)
	}
	h := Handshake{}
	h.Pstr = string(buf[1:20])
	copy(h.InfoHash[:], buf[28:48])
	copy(h.PeerID[:], buf[48:68])
	return &h, nil
}