package components

import (
	"io"
	"net/http"
)

type TextData struct {
	Text string
}

type TextComponent struct {
	Component
	Data TextData
}

func (comp *TextComponent) WriteTemplate(writer io.Writer) error {
	return comp.writeTemplate(writer, comp.Data)
}

func (comp *TextComponent) RenderTemplate(writer http.ResponseWriter) {
	comp.renderTemplate(writer, comp.Data)
}

func NewText(text string) *TextComponent {
	component := &TextComponent{
		NewComponent("text.html"),
		TextData{
			Text: text,
		},
	}

	return component
}
