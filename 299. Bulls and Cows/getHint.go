/*
You write down a number and ask your friend to guess what the number is.
Each time your friend makes a guess, you provide a hint that indicates
	how many digits in said guess match your secret number exactly in both digit and position (called "bulls") and
	how many digits match the secret number but locate in the wrong position (called "cows").
*/

package bullsandcows

import (
	"fmt"
)

// GetHint return a hint according to the secret number and friend's guess, use A to indicate the bulls and B to indicate the cows.
// Input: secret = "1807", guess = "7810"
// Output: "1A3B"
// Explanation: 1 bull and 3 cows. The bull is 8, the cows are 0, 1 and 7.
// Input: secret = "1123", guess = "0111"
// Output: "1A1B"
// Explanation: The 1st 1 in friend's guess is a bull, the 2nd or 3rd 1 is a cow.
func GetHint(secret string, guess string) string {
	if len(secret) != len(guess) {
		return "0A0B"
	}
	if secret == guess {
		return fmt.Sprintf("%vA0B", len(secret))
	}

	secretChan := make(chan rune, len(secret))
	guessChan := make(chan rune, len(guess))
	secretCounterTaskDone := make(chan bool)
	guessCounterTaskDone := make(chan bool)
	bullCounterChannel := make(chan int)

	secretCounter := make(map[rune]int)
	guessCounter := make(map[rune]int)

	go func() {
		for _, c := range secret {
			secretChan <- c
			secretCounter[c]++
		}
		secretCounterTaskDone <- true
		close(secretChan)
	}()

	go func() {
		for _, c := range guess {
			guessChan <- c
			guessCounter[c]++
		}
		guessCounterTaskDone <- true
		close(guessChan)
	}()

	go func() {
		bullCounter := 0
		for secretChar := range secretChan {
			guessChar, ok := <-guessChan
			if !ok {
				break
			}
			if secretChar == guessChar {
				bullCounter++
			}
		}
		bullCounterChannel <- bullCounter
	}()

	<-secretCounterTaskDone
	<-guessCounterTaskDone
	numBulls := <-bullCounterChannel
	numCows := -numBulls
	for key, numGuess := range guessCounter {
		numSecret, hasChar := secretCounter[key]
		if !hasChar {
			continue
		}
		numCows += min(numGuess, numSecret)
	}

	return fmt.Sprintf("%vA%vB", numBulls, numCows)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
