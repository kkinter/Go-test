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
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 는 소수 입니다"},
		{"not prime", 8, false, "8 는 소수가 아닙니다 (2 로 나눠짐)"},
		{"zero", 0, false, "0 는 소수가 아닙니다."},
		{"one", 1, false, "1 는 소수가 아닙니다."},
		{"nagative", -11, false, "음수는 소수가 아닙니다."},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: true 로 예상했지만, false 가 반환되었습니다", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: false 로 예상했지만, true 가 반환되었습니다", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: %s 를 예상했지만, %s 가 반환되었습니다", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	// os.Stdout의 복사본을 저장합니다.
	oldOut := os.Stdout

	// 읽기 및 쓰기 파이프를 생성합니다.
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close our writer
	_ = w.Close()

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
	// os.Stdout의 복사본을 저장합니다.
	oldOut := os.Stdout

	// 읽기 및 쓰기 파이프를 생성합니다.
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if !strings.Contains(string(out), "Enter a whole Num") {
		t.Errorf("intro text not correct expected ; but got %s", string(out))
	}

}

func Test_checkNumbers(t *testing.T) {
	// 에러 발생
	// res, _ := checkNumbers(bufio.NewScanner(os.Stdin))

	// if !strings.EqualFold(res, "7 는 소수 입니다") {
	// 	t.Error("wrong return value")
	// }

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole num"},
		{name: "zero", input: "0", expected: "0 는 소수가 아닙니다."},
		{name: "one", input: "1", expected: "1 는 소수가 아닙니다."},
		{name: "two", input: "2", expected: "2 는 소수 입니다"},
		{name: "three", input: "3", expected: "3 는 소수 입니다"},
		{name: "nagative", input: "-1", expected: "음수는 소수가 아닙니다."},
		{name: "typed", input: "three", expected: "Please enter a whole num"},
		{name: "decimal", input: "1.1", expected: "Please enter a whole num"},
		{name: "quit", input: "q", expected: ""},
		{name: "Quit", input: "Q", expected: ""},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s wrong value returned; got", e.name, e.expected, res)
		}
	}

}

func Test_readUserInput(t *testing.T) {
	// to test this function, we need a channel, and an instance of an io.Reader
	doneChan := make(chan bool)

	// create a reference to a bytes.Buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
