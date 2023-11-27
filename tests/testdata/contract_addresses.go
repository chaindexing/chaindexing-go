package testdata

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/chaindexing/chaindexing-go"
)

func GenerateContractAddresses(n int) []chaindexing.ContractAddress {
	var cAddresses = make([]chaindexing.ContractAddress, n)

	for i := 0; i < n; i++ {
		cAddresses[i] = chaindexing.ContractAddress{
			Address:                     gofakeit.BitcoinAddress() + gofakeit.BitcoinAddress(),
			ContractName:                gofakeit.Word() + gofakeit.Word(),
			ChainID:                     int32(gofakeit.Number(9, 99_999)),
			StartBlockNumber:            int64(gofakeit.Number(100_000, 999999999999999)),
			NextBlockNumberToIngestFrom: int64(gofakeit.Number(100_000, 999999999999999)),
			NextBlockNumberToHandleFrom: int64(gofakeit.Number(100_000, 999999999999999)),
		}
	}

	return cAddresses
}
