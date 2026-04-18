package main

import (
	"testing"
)

func TestStringFunctions(t *testing.T) {
	t.Run("length", func(t *testing.T) {
		if length("hello") != 5 {
			t.Errorf("Expected length of 'hello' to be 5")
		}
	})

	t.Run("substr", func(t *testing.T) {
		if substr("hello", 1, 3) != "ell" {
			t.Errorf("Expected substr('hello', 1, 3) to be 'ell'")
		}
	})

	t.Run("index", func(t *testing.T) {
		if index("hello", "ll") != 2 {
			t.Errorf("Expected index('hello', 'll') to be 2")
		}
	})

	t.Run("split", func(t *testing.T) {
		result := split("a,b,c", ",")
		if len(result) != 3 || result[0] != "a" || result[1] != "b" || result[2] != "c" {
			t.Errorf("Expected split('a,b,c', ',') to be ['a', 'b', 'c']")
		}
	})

	t.Run("sub", func(t *testing.T) {
		result := sub("a", "x", "banana")
		if result != "bxnana" {
			t.Errorf("Expected sub('a', 'x', 'banana') to be 'bxnana'")
		}
	})

	t.Run("gsub", func(t *testing.T) {
		result := gsub("a", "x", "banana")
		if result != "bxnxnx" {
			t.Errorf("Expected gsub('a', 'x', 'banana') to be 'bxnxnx'")
		}
	})

	t.Run("match", func(t *testing.T) {
		if match("a.*n", "banana") != 1 {
			t.Errorf("Expected match('a.*n', 'banana') to be 1")
		}
	})

	t.Run("sprintf", func(t *testing.T) {
		result := sprintf("%d %s", 42, "hello")
		if result != "42 hello" {
			t.Errorf("Expected sprintf('%d %s', 42, 'hello') to be '42 hello'")
		}
	})

	t.Run("tolower", func(t *testing.T) {
		if tolower("HELLO") != "hello" {
			t.Errorf("Expected tolower('HELLO') to be 'hello'")
		}
	})

	t.Run("toupper", func(t *testing.T) {
		if toupper("hello") != "HELLO" {
			t.Errorf("Expected toupper('hello') to be 'HELLO'")
		}
	})
}