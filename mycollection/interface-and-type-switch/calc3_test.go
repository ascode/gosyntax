package interface_and_type_switch

import (
	"fmt"
	"testing"
)

func TestAdd3(t *testing.T) {
	mock := []struct {
		a interface{}
		b interface{}
		result interface{}
	}{
		{1, 2, 3},
		{4, 5, 9},
		{"a", "b", "ab"},
	}

	for _, item := range mock {
		result := 0
		switch c := item.result.(type) {
		case int:
			result = c
		}
		if Add(item.a, item.b) != fmt.Sprint(result) {
			t.Error("add批量测试错误", "参数为：a:", item.a, ",b:", item.b, "result:", item.result)
		}
	}
}
