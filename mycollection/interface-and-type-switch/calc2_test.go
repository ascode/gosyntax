package interface_and_type_switch

import "testing"

func TestAdd2(t *testing.T) {
	if Add("1", "2") != "3" {
		t.Error("string类型相加错误")
	}
}