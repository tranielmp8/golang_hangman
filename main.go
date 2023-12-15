package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

// go run main.go words.go

func main() {
	fmt.Println("***HangmanGame***")
	// You can copy and paste words from words.go file or create your own words. If you don't want to do that you can run the program by typing: go run main.go words.go
	// words := []string{"hello", "world", "good", "night", "golang"}
	blanks := []string{}
	lives := 5
	letters := []string{}

	chosen_word := words[rand.Intn(len(words))]

	// seed := rand.NewSource(time.Now().Unix())
	// r := rand.New(seed)
	// chosen_word := words[r.Intn(len(words))]

	// fmt.Println(chosen_word) -> see random word

	for range chosen_word {
		blanks = append(blanks, "_")
	}

	for {
		fmt.Printf("💓 %d : Guess a letter ", lives)
		var guess string
		fmt.Scanln(&guess)
		guess = strings.ToLower(guess)
		containsGuess := slices.Contains(letters, guess)

		if len(guess) == 1 {
			if containsGuess {
				fmt.Println("🚩Already guessed that letter, please try again")
			}
			for _, guessLetter := range guess { //iterates by index
				correctGuess := false
				for i, wordLetter := range chosen_word { // selects the letter at that index
					letters = append(letters, blanks[i])

					if guessLetter == wordLetter {
						blanks[i] = string(guessLetter)
						correctGuess = true
					}
				}
				letters = append(letters, guess)

				if !correctGuess { // if guess not correct stays false
					lives--
				}

			}
			fmt.Println(blanks)
			fmt.Printf("%d lives:", lives)

			if lives <= 0 {
				fmt.Printf(" Sorry you lost! %d lives left - 💔", lives)
				break
			}
			if chosen_word == strings.Join(blanks, "") {
				fmt.Printf(" CONGRATS! YOU WIN! - 🙌")
				break
			}
		} else {
			fmt.Println("Only 1 letter can be guessed at a time")

		}
	}
}
