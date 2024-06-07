package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	bencode "github.com/jackpal/bencode-go"
)

type Torrent struct {
	Announce string
	Info     string
}

func decodeTorrent(data []byte) (map[string]interface{}, error) {
	var torrent map[string]interface{}

	reader := bytes.NewReader(data)
	err := bencode.Unmarshal(reader, &torrent)
	if err != nil {
		return nil, err
	}
	return torrent, nil
}

func main() {
	file := "test.torrent"

	data, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	torrent, err := decodeTorrent(data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(torrent)
}
