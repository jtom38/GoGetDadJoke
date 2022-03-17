package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	//"github.com/spf13/viper"
)


var rootCmd = &cobra.Command{
	Use:   "godadjoke",
	Short: "GoDadJoke will get you a random dad joke from the internet",
	Long: `A Fast and Flexible Static Site Generator built with love by spf13 and friends in Go.
	Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("demo")
	},

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

