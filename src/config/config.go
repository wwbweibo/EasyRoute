package config

type IConfiguration interface {
	IConfigurationSection
}

type IConfigurationSection interface {
	Bind(entity interface{})
	GetValue(key string) interface{}
	GetSection(key string) IConfigurationSection
}
