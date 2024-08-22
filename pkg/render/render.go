package render

import (
	"os"

	"github.com/charmbracelet/glamour"
)

func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func RenderMarkdown(content string) (string, error) {
	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
	)
	out, err := r.Render(content)
	return out, err
}
