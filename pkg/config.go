package pkg

type Config struct {
	HttpConfig HttpConfig
}

type HttpConfig struct {
	Host   string `json:"host" yaml:"host"`
	Port   string `json:"port" yaml:"port"`
	Prefix string `json:"prefix" yaml:"prefix"`
}
