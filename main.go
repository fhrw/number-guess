package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	random := rand.Intn(100 + 1)
	fmt.Println("Guess my number! It's between 1 and 100")
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {
		guessRemain := 9 - i
		input, err := reader.ReadString('\n')
		inputNum, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			log.Fatal(err)
		}
		if inputNum == random {
			fmt.Println("Nice! You got it...")
			break
		} else if inputNum < random {
			fmt.Println("Nice try! You're number was a little low!")
			fmt.Println(guessRemain, " guesses remain...")
		} else if inputNum > random {
			fmt.Println("Too high!")
			fmt.Println(guessRemain, "guesses remain...")
		}
	}

	fmt.Println("Sorry you didn't guess my number in time.")
	fmt.Println("The correct number was ", random)
}
