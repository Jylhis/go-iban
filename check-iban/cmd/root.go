package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/jylhis/iban"
)

var rootCmd = &cobra.Command{
	Use:   "check-iban IBAN",
	Short: "Use this to check iban",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0])
		_, err := iban.Validate(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("IBAN is valid")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
