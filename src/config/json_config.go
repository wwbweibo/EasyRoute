package config

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type JsonConfigurationBuilder struct {
	fileData []byte
}

type JsonConfiguration struct {
	entry map[string]interface{}
}

func (builder *JsonConfigurationBuilder) AddFile(file string) IConfigurationBuilder {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	builder.fileData = data
	return builder
}

func (builder *JsonConfigurationBuilder) Build() IConfiguration {
	config := &JsonConfiguration{
		entry: make(map[string]interface{}),
	}
	err := json.Unmarshal(builder.fileData, &config.entry)
	if err != nil {
		return nil
	}
	return config
}

func (config *JsonConfiguration) GetValue(key string) interface{} {
	if strings.Contains(key, ":") {
		sections := strings.Split(key, ":")
		entry := config.entry
		for idx, section := range sections {
			if value, ok := entry[section]; ok && idx == len(sections)-1 {
				return value
			} else if ok {
				temp := entry[section]
				entry = temp.(map[string]interface{})
			} else {
				return nil
			}
		}
	} else {
		if value, ok := config.entry[key]; ok {
			return value
		} else {
			return nil
		}
	}
	return nil
}

func (config *JsonConfiguration) Bind(entity interface{}) {
	data, _ := json.Marshal(config.entry)
	json.Unmarshal(data, entity)
}

func (config *JsonConfiguration) GetSection(key string) IConfigurationSection {
	c := &JsonConfiguration{}
	if strings.Contains(key, ":") {
		sections := strings.Split(key, ":")
		entry := config.entry
		for idx, section := range sections {
			if value, ok := entry[section]; ok && idx == len(sections)-1 {
				c.entry = value.(map[string]interface{})
			} else if ok {
				temp := entry[section]
				entry = temp.(map[string]interface{})
			}
		}
	} else {
		if value, ok := config.entry[key]; ok {
			c.entry = value.(map[string]interface{})
		} else {
			return nil
		}
	}
	return c
}
