package torrentfile

import (
	"io"
	"github.com/jackpal/bencode-go"
	"strconv"
	"net/url"
	"os"
)

type TorrentFile struct {
	Announce     string
	InfoHash     [20]byte
	PiecesHashes [][20]byte
	PieceLength  int
	Length       int
	Name         string
}



type BencodeInfo struct {
	Pieces       string `bencode:"pieces"`
	PiecesLength int    `bencode:"piece length"`
	Length       int    `bencode:"length"`
	Name         string `bencode:"name"`
}

type BencodeTorrent struct {
	Announce string      `bencode:"announce"`
	Info     BencodeInfo `bencode:"info"`
}

func (bto BencodeTorrent) DownloadToFile(r io.Reader) ( error ) {
	return nil 
}

// open parses a torrent file
func Open(path string) (TorrentFile, error) {
	file, err := os.Open(path)

	if err != nil {
		return TorrentFile{}, err
	}
	defer file.Close()

	bto := BencodeTorrent{}
	// create io.Reader from file of type *os.File
	
	
	err := bencode.Unmarshal(file, &bto)

	if err != nil {
		return TorrentFile{}, err
	}

	return bto.toTorrentFile()

}

func (bto BencodeTorrent) toTorrentFile() (TorrentFile, error) {
	file := TorrentFile{}
	return file, nil
}


func (t *TorrentFile) buildTrackerurl(peerID [20]byte, port uint16) (string, error){
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
