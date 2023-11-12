package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jurshsmith/ethabi-go"
)

type Log struct{}

type EventsIngesterJsonRpc interface {
	GetBlockNumber() (int64, error)
	GetLogs() ([]Log, error)
}

var transferEvent = "event Transfer(address indexed from, address indexed to, uint tokens)"

func main() {
	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/CAsX0AA7YOiV9C5c9WMTFDmgBT07g24o")
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := ethabi.ParseABI(&transferEvent)
	if err != nil {
		log.Fatal(err)
	}
	eventAbiName := ethabi.GetABIName(&transferEvent)

	contractAddress := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(6383820),
		ToBlock:   big.NewInt(6383840),
		Addresses: []common.Address{
			contractAddress,
		},
		Topics: [][]common.Hash{
			{contractAbi.Events[eventAbiName].ID},
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal("Line 50", err)
	}

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		fmt.Printf("Log Name: Transfer\n")
	}
}
