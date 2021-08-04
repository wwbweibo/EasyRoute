package rpc

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Configs is the config for rpc module
type Configs struct {
	configs map[string]Config `yaml:"configs"`
}

type Config struct {
	BaseUrl string `yaml:"base-url"`
}

func HttpGet(config Config, methodName string, params map[string]string, result interface{}) error {
	url := config.BaseUrl + "/" + methodName
	if params != nil || len(params) > 0 {
		url += "?"
		for k, v := range params {
			url += fmt.Sprintf("%s=%s&", k, v)
		}
	}
	response, err := http.Get(url[:len(url)-1])
	if err != nil {
		return err
	}
	ResponseReader(response.Body, result)
	return nil
}

func ResponseReader(reader io.Reader, result interface{}) {
	r := bufio.NewReader(reader)
	data := make([]byte, r.Size())
	r.Read(data)
	json.Unmarshal(data, result)
}
