/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	tele "gopkg.in/telebot.v3"
)

var (
	TeleToken = os.Getenv("TELE_TOKEN")
)

// mybotCmd represents the mybot command
var mybotCmd = &cobra.Command{
	Use:     "mybot",
	Aliases: []string{"start"},
	Short:   "My first bot",
	Long: `There is no long description for such simple thing`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("mybot started")

		mybot, err := tele.NewBot(
			tele.Settings{
				URL:    "",
				Token:  TeleToken,
				Poller: &tele.LongPoller{Timeout: 10 * time.Second},
			})
		if err != nil {
			log.Fatalf("Something went wrong", err)
			return
		}

		mybot.Handle(tele.OnText, func(m tele.Context) error {
			return m.Send("Let's start from command 'help' to see what I'm capable of")
		})

		mybot.Handle("/help", func(c tele.Context) error {
			return c.Send("I am very simple bot. I only support these commands: 'greeting', 'age', 'goodbye'")
		})

		mybot.Handle("/greeting", func(c tele.Context) error {
			return c.Send("Nice to see you here! You are awesome!")
		})

		mybot.Handle("/age", func(c tele.Context) error {
			return c.Send("I am very young bot. How old are you?")
		})

		mybot.Handle("/goodbye", func(c tele.Context) error {
			return c.Send("Bye. Have a good day! Hope to see you soon")
		})

		mybot.Start()
	},
}

func init() {
	rootCmd.AddCommand(mybotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mybotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mybotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
