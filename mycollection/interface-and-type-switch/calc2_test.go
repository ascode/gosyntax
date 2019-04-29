package calc

import "testing"

func TestAdd2(t *testing.T) {
	if Add("1", "2") != "3" {
		t.Error("string类型相加错误")
	}
}