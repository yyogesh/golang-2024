package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

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
		{"Not prime", 8, false, "8 is not a prime number because it is divisible by 2"},
		{"negative", -2, false, "Negative numbers are not prime, by definition!"},
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

func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	// call our function
	prompt()

	// close the write end of the pipe
	w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}

}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	// call our function
	intro()

	// close the write end of the pipe
	w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro text not correct; got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "zero", input: "0", expected: "0 is not prime, by definition!"},
		{name: "one", input: "1", expected: "1 is not prime, by definition!"},
		{name: "two", input: "2", expected: "2 is a prime number!"},
		{name: "three", input: "3", expected: "3 is a prime number!"},
		{name: "negative", input: "-1", expected: "Negative numbers are not prime, by definition!"},
		{name: "typed", input: "three", expected: "Please enter a whole number!"},
		{name: "decimal", input: "1.1", expected: "Please enter a whole number!"},
		{name: "quit", input: "q", expected: ""},
		{name: "QUIT", input: "Q", expected: ""},
		{name: "greek", input: "επτά", expected: "Please enter a whole number!"},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)

		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s", e.name, e.expected, res)
		}
	}

}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)

	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)

	<-doneChan

	close(doneChan)
}
