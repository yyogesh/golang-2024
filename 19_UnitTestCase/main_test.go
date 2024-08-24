package main

import "testing"

func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)

	if result {
		t.Errorf("with %d as test parameter, got true, but expected false", 0)
	}

	if msg != "0 is not prime, by definition!" {
		t.Error("wrong message received: ", msg)
	}

	result, msg = isPrime(7)
	if !result {
		t.Errorf("with %d as test parameter, got false, but expected true", 7)
	}

	if msg != "7 is a prime number!" {
		t.Error("wrong message received: ", msg)
	}
}

func Test_isPrime1(t *testing.T) {
	primeTests := []struct {
		name     string
		input    int
		expected bool
		msg      string
	}{
		{"Zero", 0, false, "0 is not prime, by definition!"},
		{"One", 1, false, "1 is not prime, by definition!"},
		{"Two", 2, true, "2 is a prime number!"},
		{"Seven", 7, true, "7 is a prime number!"},
	}

	for _, tt := range primeTests {
		result, msg := isPrime(tt.input)

		if tt.expected && !result {
			t.Errorf("%s: got %t, expected %t", tt.name, result, tt.expected)
		}

		if !tt.expected && result {
			t.Errorf("%s: got %t, expected %t", tt.name, result, tt.expected)
		}

		if tt.msg != msg {
			t.Errorf("%s: got %s, expected %s", tt.name, msg, tt.msg)
		}
	}
}
