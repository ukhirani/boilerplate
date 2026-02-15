package styles

import (
	"bytes"
	"fmt"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

// HighlightCode applies syntax highlighting based on filename and theme.
// It returns highlighted output ready for terminal display.
func HighlightCode(code string, filename string) (string, error) {
	lexer := lexers.Match(filename)
	if lexer == nil {
		lexer = lexers.Analyse(code)
	}
	if lexer == nil {
		lexer = lexers.Fallback
	}

	styleName := "github"
	if GetTheme().IsDark {
		styleName = "github-dark"
	}
	style := styles.Get(styleName)
	if style == nil {
		style = styles.Fallback
	}

	formatter := formatters.Get("terminal16m")
	if formatter == nil {
		return "", fmt.Errorf("no terminal formatter available")
	}

	iterator, err := lexer.Tokenise(nil, code)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := formatter.Format(&buf, style, iterator); err != nil {
		return "", err
	}

	return buf.String(), nil
}
