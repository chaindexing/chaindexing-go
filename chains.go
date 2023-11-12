package chaindexing

type Chain int

const (
	Mainnet = Chain(1)
	Goerli  = Chain(5)
	Kovan   = Chain(42)
	Sepolia = Chain(11155111)
)

type JsonRpcUrl string

type Chains = map[Chain]JsonRpcUrl
