package calc

import "testing"

func TestAdd(t *testing.T) {
	if Add(1,1) != "2" {
		t.Error("int类型加法错误")
	}
}