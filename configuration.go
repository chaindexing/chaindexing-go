package chaindexing

type Configuration struct {
	Opts
}

type Opts struct {
	Node     NodeConfiguration
	Database DatabaseConfiguration
}

type OptsFunc func(configOpts *Opts)

type NodeConfiguration struct {
	Endpoint string
}

type DatabaseConfiguration struct {
	Driver       SupportedDatabaseDrivers
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         uint16
	SSlMode      string
	MaxOpenConns int
	MaxIdleConns int
}

var AppConfig Configuration

func DefaultConfig() Opts {
	db := DatabaseConfiguration{
		Driver:       POSTGRES,
		Dbname:       "postgres",
		Username:     "postgres",
		Password:     "postgres",
		Host:         "localhost",
		Port:         5432,
		SSlMode:      "disable",
		MaxOpenConns: 10,
		MaxIdleConns: 2,
	}

	node := NodeConfiguration{
		Endpoint: "ENDPOINT",
	}

	return Opts{node, db}
}

// NewConfiguration constructor
func NewConfiguration(f ...OptsFunc) Configuration {
	c := DefaultConfig()

	for _, fn := range f {
		if fn != nil {
			fn(&c)
		}
	}

	config := Configuration{c}

	AppConfig = config

	return config
}

func GetConfig() Configuration {
	return AppConfig
}
