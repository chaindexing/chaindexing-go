package chaindexing

type Contract struct {
	name      string
	addresses []UnsavedContractAddress
}

type UnsavedContractAddress struct {
	chainId                     int
	nextBlockNumberToIngestFrom int
	nextBlockNumberToHandleFrom int
	startBlockNumber            int
	address                     string
	contractName                string
}

type ContractAddress struct {
	id                          int
	chainId                     int
	nextBlockNumberToIngestFrom int
	nextBlockNumberToHandleFrom int
	startBlockNumber            int
	address                     string
	contractName                string
}
