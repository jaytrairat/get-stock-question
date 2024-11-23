package main

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

func main() {
	var titleLength int

	var rootCmd = &cobra.Command{
		Use:   "get-stock-question",
		Short: "Formats text and saves it to the clipboard",
		Long:  `This CLI formats the given text and copies it to the clipboard.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Join the arguments into a single string
			inputText := strings.Join(args, " ")

			// Format the text (e.g., insert the title length and text)
			formattedText := fmt.Sprintf("Can you create descriptive names without quotation marks, concise usage descriptions (under %d characters), and 50 one-word keywords for my stock photo of %s", titleLength, inputText)

			// Save to clipboard
			err := clipboard.WriteAll(formattedText)
			if err != nil {
				fmt.Println("Error saving to clipboard:", err)
				return
			}

			// Output formatted text
			fmt.Println("Formatted text copied to clipboard:", formattedText)
		},
	}

	// Add the -l flag to specify title length
	rootCmd.Flags().IntVarP(&titleLength, "length", "l", 200, "Maximum length for the title description")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
	}
}
