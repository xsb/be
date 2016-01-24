package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	api "github.com/gophergala2016/be/insightapi"
)

func cliLatestBlocks() {
	var miner string

	latestBlocks, err := api.GetLatestBlocks()
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range latestBlocks.Blocks[0:20] {
		blockDatetimeUnix, _ := strconv.Atoi(b.Time)
		if err != nil {
			log.Fatal(err)
		}
		blockDatetime := time.Unix(int64(blockDatetimeUnix), 0)
		blockDatetimeHuman := blockDatetime.Format("2006-01-02 15:04:05")
		if b.PoolInfo.PoolName != "" {
			miner = "[" + b.PoolInfo.PoolName + "]"
		} else {
			miner = ""
		}
		size := b.Size / 1024
		fmt.Printf("#%d (%s) %dtxs %dKb %s\n", b.Height, blockDatetimeHuman, b.Txlength, size, miner)
	}
}

func cliBlock(block string) {
	b, err := api.GetBlock(block)
	if err != nil {
		log.Fatal(err)
	}
	blockDatetime := time.Unix(b.Time, 0)
	blockDatetimeHuman := blockDatetime.Format("2006-01-02 15:04:05")
	size := b.Size / 1024

	fmt.Printf("Block #%d\n-------------\n", b.Height)
	fmt.Printf("Number Of Transactions: %d\n", len(b.Tx))
	fmt.Printf("Height: %d\n", b.Height)
	fmt.Printf("Block Reward: %f\n", b.Reward)
	fmt.Printf("Timestamp: %s\n", blockDatetimeHuman)
	if b.PoolInfo.PoolName != "" {
		fmt.Printf("Mined by: %s\n", b.PoolInfo.PoolName)
	}
	fmt.Printf("Merkle Root: %s\n", b.Merkleroot)
	fmt.Printf("Difficulty: %f\n", b.Difficulty)
	fmt.Printf("Size (Kilobytes): %d\n", size)
	fmt.Printf("Nonce: %d\n", b.Nonce)

	fmt.Println("\nTransactions:\n-------------")
	for _, tx := range b.Tx {
		fmt.Println(tx)
	}
}

func cliTransaction(transaction string) {
	tx, err := api.GetTx(transaction)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction ID: %s\n", tx.Txid)
	fmt.Printf("Value Out: %f\n", tx.ValueOut)
	fmt.Printf("Value In: %f\n", tx.ValueIn)
	fmt.Printf("Size: %d\n", tx.Size)
	fmt.Printf("Fees: %f\n", tx.Fees)

	fmt.Println("\nInputs:\n-------")
	for _, i := range tx.Vin {
		fmt.Println(i.Addr, formatBtc(i.Value))
	}

	fmt.Println("\nOutputs:\n--------")
	for _, o := range tx.Vout {
		fmt.Println(o.ScriptPubKey.Addresses[0], o.Value)
	}
}

func cliAddress(address string) {
	a, err := api.GetAddr(address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Address: %s\n", a.AddrStr)
	fmt.Printf("Total Received: %s\n", formatBtc(a.TotalReceived))
	fmt.Printf("Total Sent: %s\n", formatBtc(a.TotalSent))
	fmt.Printf("Final Balance: %s\n", formatBtc(a.Balance))
	fmt.Printf("Unconfirmed Balance: %s\n", formatBtc(a.UnconfirmedBalance))
	fmt.Printf("Total Transactions: %d\n", len(a.Transactions))

	fmt.Println("\nTransactions:\n-------------")
	for _, tx := range a.Transactions {
		fmt.Println(tx)
	}
}
