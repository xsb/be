package insightapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	ApiURL    = "https://blockexplorer.com/api"
	UserAgent = "insight-go"
)

func GetResponse(url string) (bytes []byte, err error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("Error: " + resp.Status)
		return
	}

	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

func GetLatestBlocks() (blocklist BlockList, err error) {
	url := ApiURL + "/blocks"
	bytes, err := GetResponse(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &blocklist)

	// sometimes the api sends duplicated blocks
	blocks := blocklist.Blocks
	blocksUnique := []BlockInfo{}
	var lastHash string
	for _, b := range blocks {
		if b.Hash != lastHash {
			blocksUnique = append(blocksUnique, b)
			lastHash = b.Hash
		}
	}
	blocklist.Blocks = blocksUnique

	return
}

func GetBlockByHash(blockHash string) (block Block, err error) {
	url := ApiURL + "/block/" + blockHash
	bytes, err := GetResponse(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &block)
	return
}

func GetBlockByHeight(blockHeight string) (block Block, err error) {
	url := ApiURL + "/block-index/" + blockHeight
	bytes, err := GetResponse(url)
	if err != nil {
		return
	}

	var blockIndex BlockIndex
	err = json.Unmarshal(bytes, &blockIndex)
	block, err = GetBlockByHash(blockIndex.BlockHash)
	return
}

func GetTx(txId string) (tx Tx, err error) {
	url := ApiURL + "/tx/" + txId
	bytes, err := GetResponse(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &tx)
	return
}

func GetAddr(addrStr string) (addr Addr, err error) {
	url := ApiURL + "/addr/" + addrStr
	bytes, err := GetResponse(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &addr)
	return
}
