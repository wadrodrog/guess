package game

import "fmt"

type Difficulty struct {
	Name     string
	Attempts uint8
}

func (d Difficulty) String() string {
	return fmt.Sprintf("%s (%d attempts)", d.Name, d.Attempts)
}

var difficulties = [3]Difficulty{{"Easy", 10}, {"Medium", 5}, {"Hard", 3}}

type difficultyError struct {
	code  uint8
	name  string
	input string
}

func (e *difficultyError) Error() string {
	switch e.code {
	case 0:
		e.name = "difficulty is not set"
	case 1:
		e.name = fmt.Sprintf("unknown difficulty: %s", e.input)
	}
	return fmt.Sprintf("%s.\n", e.name)
}

func ParseDifficulty(input *string) (Difficulty, error) {
	switch *input {
	case "easy", "e":
		return difficulties[0], nil
	case "medium", "m":
		return difficulties[1], nil
	case "hard", "h":
		return difficulties[2], nil
	case "":
		return difficulties[1], &difficultyError{code: 0}
	default:
		return difficulties[1], &difficultyError{code: 1}
	}
}
