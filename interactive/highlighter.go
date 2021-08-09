package interactive

import (
	"bytes"

	"github.com/alecthomas/chroma"
)

type Highlighter struct {
	lexer     chroma.Lexer
	formatter chroma.Formatter
	style     *chroma.Style
}

func newHighlighter(lexer chroma.Lexer, formatter chroma.Formatter, style *chroma.Style) *Highlighter {
	h := new(Highlighter)
	h.formatter = formatter
	h.lexer = lexer
	h.style = style
	return h
}

func (h *Highlighter) Highlight(text string) ([]byte, error) {
	iterator, err := h.lexer.Tokenise(nil, text)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer([]byte{})
	h.formatter.Format(buffer, h.style, iterator)
	return buffer.Bytes(), nil
}
