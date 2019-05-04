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

func NewText(text string, tag string) *TextComponent {
	if !tagnameRegexp.MatchString(tag) {
		panic(tag + " is not a valid tag")
	}

	component := &TextComponent{
		Component: *NewDirectTemplateComponent(tag, "<"+tag+">{{.Text}}</"+tag+">"),
		Data: TextData{
			Text: text,
		},
	}

	return component
}

func H1(text string) *TextComponent {
	return NewText(text, "h1")
}

func P(text string) *TextComponent {
	return NewText(text, "p")
}
