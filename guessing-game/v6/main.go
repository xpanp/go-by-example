package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())  //使用种子，让每次数字可以变化
	secretNumber := rand.Intn(maxNum) //制造随机数
	fmt.Println("Please input your guess")
	var guess int
	for {
		_, err := fmt.Scanf("%d\n", &guess)
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber { //判断大小关系
			fmt.Println("Your guess is bigger than the secret number.Please try again.")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number.Please try again.")
		} else {
			fmt.Println("Correct,you Legend!")
			break
		}
	}
}
