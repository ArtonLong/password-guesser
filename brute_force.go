package main

import (
	"fmt"
	"time"
)

var password string = "zzzzzz"

func checkPassword(attempt string) bool {
	return password == attempt
}

func getPasswordFromDecimal(decimal int, characters string) string {
	password := ""
	for decimal > 0 {
		remainder := decimal % len(characters)
		password = string(characters[remainder]) + password
		decimal = decimal / len(characters)
	}
	return password
}

func bruteForce(characters string, maxLength int, ch chan string, threadIndex int, threads int) {
	attempts := threadIndex

	for {
		password := getPasswordFromDecimal(attempts, characters)
		if checkPassword(password) {
			ch <- password + " attempts: " + fmt.Sprint(attempts)
			close(ch)
			return
		}
		if attempts%100000 == 0 {
			fmt.Println("attempts: ", attempts, " password: ", password)
		}
		if len(password) > maxLength {
			break
		}
		attempts += threads
	}
	fmt.Println("password not found")
}

func getTime() int64 {
	now := time.Now()
	return now.Unix()
}

func main() {
	startTime := getTime()

	characters := "0123456789abcdefghijklmnopqrstuvwxyz"
	maxLength := 15
	threads := 8

	c := make(chan string)

	for i := range threads {
		go bruteForce(characters, maxLength, c, i, threads)
	}

	result := <-c
	fmt.Println("found password: ", result)

	fmt.Println(getTime() - startTime)

	//fmt.Println(getPasswordFromDecimal(2176782335, characters))
}
