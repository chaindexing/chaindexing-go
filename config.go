package chaindexing

type ConfigErrorType int

const (
	NoContract ConfigErrorType = iota
	NoChain
)

type ConfigError struct {
	Type    ConfigErrorType
	Message string
}

type Config struct {
	Chains          map[string]string
	contracts       []Contract
	blocksPerBatch  int
	handlerRateMs   int
	ingestionRateMs int
}
