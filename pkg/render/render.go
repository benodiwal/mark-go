package render

import (
	"fmt"
	"os"
)

func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func RenderMarkdown(content string){
	fmt.Println(content)
}
