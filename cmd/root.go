package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	uppercase bool
	lowercase bool
	camelcase bool
	wordcount bool
	charcount bool
	titlecase bool
	reverse   bool
)

var RootCmd = &cobra.Command{
	Use:     "secretctl",
	Aliases: []string{"sectl"},
	Short:   "A tool for secret management for platforms such as Vault, GitLab etc.",

	Args: cobra.ArbitraryArgs,
	Run:  runTextFormatter,
}

func init() {
	RootCmd.Flags().BoolVarP(&uppercase, "upper", "u", false, "Convert text to uppercase")
	RootCmd.Flags().BoolVarP(&lowercase, "lower", "l", false, "Convert text to lowercase")
	RootCmd.Flags().BoolVarP(&titlecase, "title", "t", false, "Convert text to title case")
	RootCmd.Flags().BoolVar(&wordcount, "words", false, "Count words in text")
	RootCmd.Flags().BoolVar(&charcount, "chars", false, "Count characters in text")
	RootCmd.Flags().BoolVar(&reverse, "reverse", false, "Reverse the input")
}

func runTextFormatter(cmd *cobra.Command, args []string) {
	var text string

	// Get input text from args or stdin
	if len(args) > 0 {
		text = strings.Join(args, " ")
	} else {
		// Read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		text = strings.Join(lines, "\n")
	}

	if text == "" {
		fmt.Println("No input text provided")
		return
	}

	result := text

	if uppercase && lowercase {
		fmt.Printf("You can't uppercase and lowercase at same time!")
		return
	}

	// Apply formatting transformations
	if uppercase {
		result = strings.ToUpper(result)
	}
	if lowercase {
		result = strings.ToLower(result)
	}
	if titlecase {
		result = strings.Title(result)
	}

	// Output the formatted text
	fmt.Println(result)

	// Show analysis if requested
	if wordcount {
		words := len(strings.Fields(text))
		fmt.Printf("Words: %d\n", words)
	}
	if charcount {
		chars := len(text)
		fmt.Printf("Characters: %d\n", chars)
	}

	if reverse {
		reversive := make([]byte, 0)

		for i := len(text); i > 0; i-- {
			reversive = append(reversive, result[i])
		}
		fmt.Printf("Reverse: %s\n", reversive)
	}

}
