package main

import (
	"fmt"
	"math/rand"
	"time"

	"bufio"
	"os"
	"strconv"
	"strings"
)

const maxNum = 100

func main() {
	rand.Seed(time.Now().UnixNano())

	secretNumber := rand.Intn(maxNum)
	reader := bufio.NewReader(os.Stdin)

	//fmt.Println("The secret number:", secretNumber)

	for {
		fmt.Print("Please enter your guess: ")
		// 这里的参数为读取字符串的分隔符
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input.Please try again.", err)
			return
		}
		// 移除input尾部的\r\n字符串
		input = strings.TrimSuffix(input, "\r\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input.Please enter an integer value.")
			fmt.Println(err)
			return
		}

		fmt.Println("Your guess is", guess)

		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number.Please enter again.")

		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number.Please enter again.")
		} else {
			fmt.Println("Correct, you legend!")
			break
		}

		fmt.Println()
	}

}
