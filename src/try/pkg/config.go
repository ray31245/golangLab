package pkg

type Config struct {
	A string
}

var cfg Config

func SetA(s string) Config {
	cfg.A = s
	return cfg
}

func GetConfig() Config {
	return cfg
}
