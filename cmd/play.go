package cmd

import (
	"github.com/wadrodrog/guess/game"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var difficultyString string

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play the game",
	RunE: func(cmd *cobra.Command, args []string) error {
		difficultyVar, ok := viper.Get("difficulty").(string)
		if ok && difficultyString == "" {
			difficultyString = difficultyVar
		}

		difficulty, err := game.ParseDifficulty(&difficultyString)

		if err != nil {
			return err
		}

		game.StartGame(difficulty)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	playCmd.Flags().StringVarP(
		&difficultyString, "difficulty", "d", "",
		"Number of attempts (easy: 10, medium: 5, hard: 3)",
	)
}
