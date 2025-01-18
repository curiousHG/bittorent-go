package client

import (
	"github.com/curiousHG/bittorrent-go/message"
	"github.com/curiousHG/bittorrent-go/bitfield"
)

type Client struct {
	Choked bool
	Bitfield bitfield.Bitfield
}

func (*Client) Read() (*message.Message, error) {
	return nil, nil
}
