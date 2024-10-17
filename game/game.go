package game

import (
	"fmt"

	"github.com/wadrodrog/guess/utils"
)

const minRange uint8 = 1
const maxRange uint8 = 100

type Game struct {
	Attempts   uint8
	Difficulty Difficulty
	secret     uint8
	win        bool
}

func NewGame(difficulty Difficulty) Game {
	game := Game{
		Attempts:   0,
		Difficulty: difficulty,
		secret:     utils.RandomNumber(minRange, maxRange),
		win:        false,
	}
	return game
}

func StartGame(difficulty Difficulty) {
	fmt.Println("Let's start the game!")
	fmt.Println("Difficulty:", difficulty)

	game := NewGame(difficulty)

	for !game.IsFinished() {
		fmt.Println("")
		input := utils.ReadNumber(1, 100, "Enter your guess")
		fmt.Println(game.TryGuess(input))
	}

	if !game.HasWon() {
		fmt.Printf("Game over! You didn't guess the correct number %d.\n",
			game.secret,
		)
	}
}

func (game *Game) TryGuess(number uint8) string {
	if number < minRange || number > maxRange {
		return fmt.Sprintf("Invalid number! It must be in range %d-%d.",
			minRange, maxRange)
	}

	game.Attempts += 1

	if number < game.secret {
		return fmt.Sprintf("Incorrect! The number is greater than %d.",
			number)
	}
	if number > game.secret {
		return fmt.Sprintf("Incorrect! The number is less than %d.", number)
	}

	return game.WinGame()
}

func (game *Game) WinGame() string {
	game.win = true
	return fmt.Sprintf(
		"Congratulations! You guessed the correct number in %d attempts.",
		game.Attempts,
	)
}

func (game Game) HasWon() bool {
	return game.win
}

func (game Game) HasAttempts() bool {
	return game.Attempts < game.Difficulty.Attempts
}

func (game Game) IsFinished() bool {
	return game.HasWon() || !game.HasAttempts()
}
