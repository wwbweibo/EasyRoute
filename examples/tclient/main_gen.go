// code generate by codegenerator, DO NOT EDIT;

package main

import (
	"github.com/wwbweibo/EasyRoute/rpc"
)

func NewHomeController(config rpc.Config) *HomeController {
	return &HomeController{
		Index: func() (string, error) {
			result := ""
			err := rpc.HttpGet(config, "home", "Index", nil, &result)
			if err != nil {
				return result, err
			}
			return result, nil
		},
		IndexA: func(a string) (string, error) {
			result := ""
			params := make(map[string]string)
			params["a"] = rpc.JsonSerialize(a)
			err := rpc.HttpGet(config, "home", "IndexA", params, &result)
			if err != nil {
				return result, err
			}
			return result, nil
		},
		IndexPerson: func(person Person) (Person, error) {
			result := Person{}
			params := make(map[string]string)
			params["person"] = rpc.JsonSerialize(person)
			err := rpc.HttpGet(config, "home", "IndexPerson", params, &result)
			if err != nil {
				return result, err
			}
			return result, nil
		},
		PostIndex: func() (string, error) {
			result := ""
			err := rpc.HttpPost(config, "home", "PostIndex", nil, &result)
			if err != nil {
				return result, err
			}
			return result, nil
		},
	}
}
