package mdx

import (
	"os"
	"strings"
)

type invalidFileError struct {
	error
}

func (e *invalidFileError) Error() string {
	return "Invalid file type. File must have .md or .mdx extension"
}

// Transform .mdx or .md file into HTML string.
// On successful transformation, returns string representing HTML and nil error.
// On failure returns empty string with non nil error.
func Transform(inputFilename string) (string, error) {
	if !(strings.HasSuffix(inputFilename, ".md") || strings.HasSuffix(inputFilename, ".mdx")) {
		return "", &invalidFileError{}
	}

	data, readErr := os.ReadFile(inputFilename)
	if readErr != nil {
		return "", readErr
	}

	lexer := newLexer(string(data))
	parser := newParser(lexer)
	elements, parseErr := parser.parse(eof)

	if parseErr != nil {
		return "", parseErr
	}

	htmlString := transformMDX(elements)

	return htmlString, nil
}

// Generates HTML file based on the given configuration object.
// On successful generation, returns number of bytes written to file and nil error.
// On failure returns bytes written with non nil error.
func Generate(config *GeneratorConfig) (int, error) {
	if !(strings.HasSuffix(config.InputFilename, ".md") || strings.HasSuffix(config.InputFilename, ".mdx")) {
		return 0, &invalidFileError{}
	}

	data, readErr := os.ReadFile(config.InputFilename)
	if readErr != nil {
		return 0, readErr
	}

	lexer := newLexer(string(data))
	parser := newParser(lexer)
	elements, parseErr := parser.parse(eof)

	if parseErr != nil {
		return 0, parseErr
	}

	n, err := generateHtml(elements, config)

	return n, err
}
