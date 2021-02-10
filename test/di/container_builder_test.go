package di

import (
	"github.com/wwbweibo/EasyRoute/src/di"
	"testing"
)

type Person struct {
	Name     string
	password string
}

func TestContainerBuilder_Test_Register_Instance(t *testing.T) {
	builder := di.NewDefaultContainerBuilder()
	builder.AddInstance(&Person{Name: "123"})
	builder.AddInstance(Person{Name: "456"})
}
