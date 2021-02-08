package config

import (
	"github.com/wwbweibo/EasyRoute/src/config"
	"testing"
)

type person struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type personInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJsonConfigurationBuilder_AddFile(t *testing.T) {
	var builder config.IConfigurationBuilder = new(config.JsonConfigurationBuilder)
	builder = builder.AddFile("test.json")
	if builder == nil {
		t.Error("the builder is empty")
	}
}

func TestJsonConfigurationBuilder_AddFile_Return_Nil(t *testing.T) {
	var builder config.IConfigurationBuilder = new(config.JsonConfigurationBuilder)
	builder = builder.AddFile("test1.json")
	if builder != nil {
		t.Error("the builder is not empty")
	}
}

func TestJsonConfigurationBuilder_Build(t *testing.T) {
	var builder config.IConfigurationBuilder = new(config.JsonConfigurationBuilder)
	var config config.IConfiguration = builder.AddFile("test.json").Build()
	if config == nil {
		t.Error("the config is empty")
	}
}

func TestJsonConfigurationBuilder_Build_Return_Nil(t *testing.T) {
	var builder config.IConfigurationBuilder = new(config.JsonConfigurationBuilder)
	var config config.IConfiguration = builder.AddFile("error_test.json").Build()
	if config != nil {
		t.Error("the config is not empty")
	}
}

func TestJsonConfiguration_GetValue(t *testing.T) {
	var builder config.IConfigurationBuilder = new(config.JsonConfigurationBuilder)
	var config config.IConfiguration = builder.AddFile("test.json").Build()
	username := config.GetValue("username")
	t.Logf("username: %s", username)
	if username != "123" {
		t.Error("username not equal")
	}
	age := config.GetValue("info:age")
	t.Logf("age: %f", age)
	if age != 12.0 {
		t.Error("age not equal")
	}
}

func TestJsonConfiguration_Bind(t *testing.T) {
	var builder config.IConfigurationBuilder = new(config.JsonConfigurationBuilder)
	var config config.IConfiguration = builder.AddFile("test.json").Build()
	p := new(person)
	config.Bind(p)
	t.Logf("UserName: %s", p.Username)
}

func TestJsonConfiguration_GetSection(t *testing.T) {
	var builder config.IConfigurationBuilder = new(config.JsonConfigurationBuilder)
	var config config.IConfiguration = builder.AddFile("test.json").Build().GetSection("info")
	name := config.GetValue("name")
	if name != "王二狗" {
		t.Error("name not equal")
	}

}
