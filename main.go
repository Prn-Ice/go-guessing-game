package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("This is a number guessing game!")
	
	// Generata a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10) //generates a number from 0 - 9
	
	var guess int
	guessCount := 0
	for{
		fmt.Printf("What's your guess (%d guesses left)\n", 5-guessCount)
		fmt.Scan(&guess)
		if guessCount > 5 {
			break
		}else if guess > secretNumber {
			fmt.Println("Too Big")
		}else if guess < secretNumber {
			fmt.Println("Too Small")
		}else {
			fmt.Println("You Win! 🤩️")
			break
		}
		guessCount ++
	}
	
}