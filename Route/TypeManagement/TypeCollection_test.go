package TypeManagement

import "testing"

type testType struct {
}

func Test_new_instance(t *testing.T) {
	instance := NewTypeCollect()
	if instance == nil || len(instance.types) <= 0 {
		t.Error("the instance is nil or type collect is empty")
	}
}

func Test_register_type_by_instance(t *testing.T) {
	inst := NewTypeCollect()
	inst.Register(testType{})
	if _, ok := inst.types["TypeManagement.testType"]; !ok {
		t.Error("Error to get Registered type")
	}
}

func Test_register_type_by_instancePointer(t *testing.T) {
	inst := NewTypeCollect()
	inst.Register(&testType{})
	if _, ok := inst.types["TypeManagement.testType"]; !ok {
		t.Error("Error to get Registered type")
	}
}

func Test_register_type_already_exist(t *testing.T) {
	inst := NewTypeCollect()
	registeredTypes := len(inst.types)

	inst.Register(1)

	afterRegister := len(inst.types)

	if registeredTypes != afterRegister {
		t.Errorf("test failed: the expected type count is %d but got %d", registeredTypes, afterRegister)
	}
}

func Test_create_instance(t *testing.T) {
	inst := NewTypeCollect()
	inst.Register(testType{})

	c, err := inst.InstanceOf("TypeManagement.testType")
	if err != nil {
		t.Error("try to create instance error", err)
	}
	if _, ok := c.(testType); !ok {
		t.Error("Error to cast type")
	}
}

func Test_create_instance_type_not_register(t *testing.T) {
	inst := NewTypeCollect()
	_, err := inst.InstanceOf("TypeManagement.testType")
	if err == nil {
		t.Error("test failed, error should not be empty")
	}
}
