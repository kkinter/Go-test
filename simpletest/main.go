package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()

	//사용자가 종료하기를 원하는 시점을 표시하는 채널을 만듭니다.
	doneChan := make(chan bool)

	go readUserInput(os.Stdin, doneChan)
	// doneChan이 값을 얻을 때 까지 차단.
	<-doneChan

	close(doneChan)

	fmt.Println("Goodbye.")

}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole num", false
	}

	_, msg := isPrime(numToCheck)

	return msg, false
}

func intro() {
	fmt.Println("Is it Prime ?")
	fmt.Println("-------------")
	fmt.Println("Enter a whole Num. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d 는 소수가 아닙니다.", n)
	}

	if n < 0 {
		return false, "음수는 소수가 아닙니다."
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d 는 소수가 아닙니다 (%d 로 나눠짐)", n, i)
		}
	}

	return true, fmt.Sprintf("%d 는 소수 입니다", n)
}
