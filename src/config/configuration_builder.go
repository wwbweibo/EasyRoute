package config

type IConfigurationBuilder interface {
	AddFile(file string) IConfigurationBuilder
	Build() IConfiguration
}
