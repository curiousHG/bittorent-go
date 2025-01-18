package handshake

// test read function from handshake.go

import (
	"bytes"
	"testing"
	// "fmt"
	// "os"
)

func TestRead(t *testing.T) {
	h := Handshake{
		Pstr:     "BitTorrent protocol",
		InfoHash: [20]byte{0x01, 0x02},
		PeerID:   [20]byte{0x03, 0x04},
	}
	buf := bytes.NewBuffer(h.Serialize())
	// t.Log("hello")
	t.Log(buf)
	h2, err := Read(buf)
	t.Log(h2)
	if err != nil {
		t.Fatalf("Read() failed: %s", err)
	}
	if h.Pstr != h2.Pstr {
		t.Fatalf("Pstr mismatch: %s != %s", h.Pstr, h2.Pstr)
	}
	if h.InfoHash != h2.InfoHash {
		t.Fatalf("InfoHash mismatch: %s != %s", h.InfoHash, h2.InfoHash)
	}
	if h.PeerID != h2.PeerID {
		t.Fatalf("PeerID mismatch: %s != %s", h.PeerID, h2.PeerID)
	}
}