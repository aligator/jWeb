package components

import (
	"io"
	. "jWeb/filetypes"
	"net/http"
)

type PageData struct {
	Style Css
	Title string
	Body  []Templated
}

type PageContainer struct {
	*Container
	Data PageData
}

func (cont *PageContainer) WriteTemplate(writer io.Writer) error {
	panic("implement me")
}

func (cont *PageContainer) RenderTemplate(writer http.ResponseWriter) {
	cont.Container.RenderTemplate(writer)
}

func NewPage(style *Css, title string, body []Templated) *PageContainer {
	page := PageData{
		Style: *style,
		Title: title,
		Body:  body,
	}

	component := &PageContainer{
		Container: NewContainer(ContainerData{
			TemplateName: "page.html",
			Data:         &page,
			Children:     &body,
		}),
		Data: page,
	}

	return component
}
