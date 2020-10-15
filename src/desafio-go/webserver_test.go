package main

import "testing"

func TestGreeting(t *testing.T) {
	tables := []struct {
		x string
		n string
	}{
		{"a", "<b>a</b>"},
		{"code", "<b>code</b>"},
		{"hello", "<b>hello</b>"},
	}

	for _, table := range tables {
		result := greeting(table.x)
		if result != table.n {
			t.Errorf("Fail: Result of (%s) incorrect, return: %s, expected: %s.", table.x, result, table.n)
		}
	}
}
