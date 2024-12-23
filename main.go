package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/url"
	"strconv"


)

func (bto bencodeTorrent) toTorrentFile() (TorrentFile, error) {
	file := TorrentFile{}
	return file,nil
}

func (t *TorrentFile) buildTrackerurl (peerID [20]byte, port uint16) (string, error){
	base, err := url.Parse(t.Announce)

	if err != nil {
		return "", err
	}

	params := url.Values {
		"info_hash" : []string{string(t.InfoHash[:])},
		"peer_id":    []string{string(peerID[:])},
        "port":       []string{strconv.Itoa(int(port))},
        "uploaded":   []string{"0"},
        "downloaded": []string{"0"},
        "compact":    []string{"1"},
        "left":       []string{strconv.Itoa(t.Length)},
	}

	base.RawQuery = params.Encode()
	return base.String(), nil
}



conn, err := net.Dial("tcp", peer.String, 3*time.Second)
if err != nil {
	return nil, err
}