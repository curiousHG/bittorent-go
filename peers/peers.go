package peers

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"time"
)

type Peer struct {
	IP net.IP
	Port uint16
}

func Unmarshal(peersBin []byte) ([]Peer, error){
	const peerSize = 6 
	numPeers := len(peersBin)/peerSize
	if (len(peersBin)%peerSize) != 0 {
		err := fmt.Errorf("recieved malformed peers")
		return nil, err
	}

	peers := make([]Peer, numPeers)
	for i := 0; i < numPeers; i++ {
		offset := i*peerSize
		peers[i].IP = net.IP(peersBin[offset:offset + 4])
		peers[i].Port = binary.BigEndian.Uint16(peersBin[offset + 4:offset + 6])
	}
	return peers, nil
}

func (p Peer) String() string {
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}
func connectToPeer(peer Peer) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", peer.String(), 3*time.Second)
	if err != nil {
		return nil, err
	}
	return conn, nil
}