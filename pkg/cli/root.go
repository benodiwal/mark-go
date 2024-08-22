package cli

import (
	"fmt"
	"os"

	"github.com/benodiwal/mark-go/pkg/render"
	"github.com/spf13/cobra"
)

var version = "1.0.0"
var inputFile string
var markdownText string

func init() {
	rootCmd.Version = version
	rootCmd.SetVersionTemplate("mark-go version {{.Version}}\n")
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.Flags().StringVarP(&inputFile, "file", "f", "", "Input markdown file")
	rootCmd.Flags().StringVarP(&markdownText, "text", "t", "", "Direct markdown text")
}

var rootCmd = &cobra.Command{
	Use: "mark-go",
	Short: "mark-go is a cli tool to render markdown directly on the terminal",
	Run: func(cmd *cobra.Command, args []string) {
		if inputFile == "" && markdownText == "" {
			cmd.Help()
			return
		}

		var content string
		var err error

		if inputFile != "" {
			content, err = render.ReadFile(inputFile)
		} else {
			content = markdownText
		}

		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		render.RenderMarkdown(content)
	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
