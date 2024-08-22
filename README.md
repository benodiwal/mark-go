# mark-go

mark-go is a command-line interface (CLI) tool that renders Markdown directly in the terminal. It provides a quick and easy way to preview Markdown files or snippets without leaving your command line environment.

## Features

- Render Markdown files directly in the terminal
- Render Markdown text input on the fly
- Syntax highlighting and formatting for better readability
- Debug mode for troubleshooting

## Installation

To install mark-go, you need to have Go installed on your system. If you don't have Go installed, you can download it from [the official Go website](https://golang.org/dl/).

Once Go is installed, you can install mark-go using the following steps:

1. Clone the repository:
   ```
   git clone https://github.com/benodiwal/mark-go.git
   ```

2. Navigate to the project directory:
   ```
   cd mark-go
   ```

3. Build the project:
   ```
   go build
   ```

This will create an executable named `mark-go` in your project directory.

## Usage

mark-go supports two main ways of input: file input and direct text input.

### Rendering a Markdown file

To render a Markdown file, use the `-f` or `--file` flag followed by the path to your Markdown file:

```
./mark-go -f path/to/your/markdown_file.md
```

### Rendering Markdown text directly

To render Markdown text directly, use the `-t` or `--text` flag followed by your Markdown text in quotes:

```
./mark-go -t "# Hello, Markdown!\n\nThis is a sample text."
```

## Dependencies

mark-go uses the following external libraries:

- [github.com/charmbracelet/glamour](https://github.com/charmbracelet/glamour) for rendering Markdown in the terminal
- [github.com/spf13/cobra](https://github.com/spf13/cobra) for creating the command-line interface

Make sure these dependencies are installed before building the project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
