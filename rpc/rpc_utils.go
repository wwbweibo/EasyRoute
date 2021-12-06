package rpc

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

// Configs is the config for rpc module
type Configs struct {
	configs map[string]Config `yaml:"configs"`
}

type Config struct {
	BaseUrl string `yaml:"base-url"`
}

func HttpGet(config Config, controllerName, methodName string, params map[string]string, result interface{}) error {
	url := fmt.Sprintf("%s/%s/%s", config.BaseUrl, controllerName, methodName)
	if params != nil || len(params) > 0 {
		url += "?"
		for k, v := range params {
			url += fmt.Sprintf("%s=%s&", k, v)
		}
		url = url[:len(url)-1]
	}
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	return ResponseReader(response.Body, result)
}

func HttpPost(config Config, controllerName, methodName string, params map[string]string, result interface{}) error {
	url := fmt.Sprintf("%s/%s/%s", config.BaseUrl, controllerName, methodName)
	param := make(map[string][]string)
	for k, v := range params {
		param[k] = []string{v}
	}
	response, err := http.PostForm(url, param)
	if err != nil {
		return err
	}
	return ResponseReader(response.Body, result)
}

func ResponseReader(reader io.Reader, result interface{}) error {
	r := bufio.NewReader(reader)
	data := make([]byte, r.Size())
	n, err := r.Read(data)
	if err != nil && err != io.EOF {
		return err
	}
	err = json.Unmarshal(data[:n], result)
	return err
}

func JsonSerialize(a interface{}) string {
	v, _ := json.Marshal(a)
	return string(v)
}

func ParseInput(input interface{}) map[string]string {
	inputType := reflect.TypeOf(input)
	result := make(map[string]string)
	for i := 0; i < inputType.NumField(); i++ {
		result[inputType.Field(i).Name] = JsonSerialize(reflect.ValueOf(input).Field(i))
	}
	return result
}
