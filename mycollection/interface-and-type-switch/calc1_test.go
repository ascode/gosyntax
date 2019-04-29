package interface_and_type_switch

import "testing"

func TestAdd(t *testing.T) {
	if Add(1,1) != "2" {
		t.Error("int类型加法错误")
	}
}