package p2p

import (

	"log"

	// "github.com/curiousHG/bittorrent-go/bitfield"
	"github.com/curiousHG/bittorrent-go/message"
	"github.com/curiousHG/bittorrent-go/peers"
	"github.com/veggiedefender/torrent-client/client"
)

type pieceWork struct {
	index int
	hash [20]byte
	length int
}

type pieceResult struct {

}

type pieceProgress struct {
	index int
	client *client.Client
	buf []byte
	downloaded int
	requested int
	backlog int
}

func (state *pieceProgress) readMessage() error {
	msg, err := state.client.Read()
	switch msg.ID {
		case message.MsgUnchoke:
			state.client.Choked = false
		case message.MsgChoke:
			state.client.Choked = true
		case message.MsgHave:
			index, err := message.ParseHave(msg)
			state.client.Bitfield.SetPiece(index)
		case message.MsgPiece:
			n, err := message.ParsePiece(state.index, state.buf, msg)
			state.downloaded += n
			state.backlog--
	}
	return nil
}

type Torrent struct {
	PieceLength int
	PiecesHashes [][20]byte
	PeerID [20]byte
	InfoHash [20]byte
	pieceWork []*pieceWork

}

func(t* Torrent) startDownloadWorker(peer peers.Peer, workQueue chan *pieceWork, results chan *pieceResult){
	c, err := client.New(peer, t.PeerID, t.InfoHash)
	if err != nil {
		log.Printf("Could not handshake with %s. Disconnecting\n", peer.IP)
		return
	}
	defer c.Conn.Close()
	log.Printf("Completed handshake with %s\n", peer.IP)

	c.SendUnchoke()
	c.SendInterested()

	for pw := range workQueue {
		if !c.Bitfield.HasPiece(pw.index){
			workQueue <- pw
			continue
		}

		buf, err := attemptDownloadPiece(c, pw)
		if err != nil {
			log.Printf("Exiting", err)
			workQueue <- pw
			return
		}
		err = checkIntegrity(pw, buf)
		if err != nil {
			log.Printf("Piece #%d failed integrity check\n", pw.index)
			workQueue <- pw
			continue
		}
		c.SendHave(pw.index)
		results <- &pieceResult{pw.index, buf}
	}
}
const MaxBlockSize = 16384
const MaxBacklog = 5

func attemptDownloadPiece(c *client.Client, pw *pieceWork) ([]byte, error) {
	state := pieceProgress{
		index: pw.index,
		client: c,
		buf: make([]byte, pw.length),
	}

	c.conn.SetDea
}

func checkIntegrity(pw *pieceWork, buf [] byte) (error){

}

func (t* Torrent) Download([]byte, error) {
	workQueue := make(chan *pieceWork, len(t.pieceWork))
	results := make(chan *pieceResult)
	for index,hash := range t.PiecesHashes {
		length := t.PieceLength
		workQueue <- &pieceWork{index, hash, length}
	}

	for _, peer := range t.Peers {
		go t.startDownloadWorker(peer, workQueue, results)
	}
	buf := make([]byte, t.PieceLength)
	donePieces := 0
	for donePieces < len(t.PiecesHashes) {
		res := <-results
		begin, end := t.calculateBoundsForPiece(res.index)
		copy(buf[begin:end], res.buf)
		donePieces++
	}
	close(workQueue)
}