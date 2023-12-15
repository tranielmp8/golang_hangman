package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

func main() {
	fmt.Println("***Hangman***")
	words := []string{"hello", "world", "good", "night", "golang"}
	blanks := []string{}
	lives := 5
	letters := []string{}

	chosen_word := words[rand.Intn(len(words))]

	// seed := rand.NewSource(time.Now().Unix())
	// r := rand.New(seed)
	// chosen_word := words[r.Intn(len(words))]
	fmt.Println(chosen_word)

	// for _, character := range chosen_word {
	// 	fmt.Printf("%c\n", character)
	// }

	for range chosen_word {
		blanks = append(blanks, "_")
	}

	fmt.Printf("Word: %s Blanks: %s\n", chosen_word, strings.Join(blanks, ""))

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
			for _, guessLetter := range guess {
				correctGuess := false
				for i, wordLetter := range chosen_word {
					letters = append(letters, blanks[i])

					if guessLetter == wordLetter {
						blanks[i] = string(guessLetter)
						correctGuess = true
					}
				}
				letters = append(letters, guess)

				if !correctGuess {
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
