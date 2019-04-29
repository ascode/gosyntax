package interface_and_type_switch

import "fmt"

func Add(a, b interface{}) string {
	var a1, b1 int
	a1, b1 = 0, 0
	switch at := a.(type) {
	case int:
		a1 = at
	}
	switch bt := b.(type) {
	case int:
		b1 = bt
	}
	return fmt.Sprint(a1 + b1)
}