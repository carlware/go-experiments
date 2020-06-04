package fibo

import "testing"

func TestChanneled(t *testing.T) {
	for _, ft := range FibTests {
		if v := FibChanneled(ft.a); v != ft.expected {
			t.Errorf("FibChanneled(%d) returned %d, expected %d", ft.a, v, ft.expected)
		}
	}
}
