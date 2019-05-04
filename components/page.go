package components

import (
	"jWeb/filetypes/css"
)

type PageData struct {
	Style css.Css
	Title string
	Body  []Templated
}

type PageContainer struct {
	*Container
	Data PageData
}

func NewPage(style *css.Css, title string, body ...Templated) *PageContainer {
	page := PageData{
		Style: *style,
		Title: title,
		Body:  body,
	}

	component := &PageContainer{
		Container: NewContainer("page.html", ContainerData{
			Data:     &page,
			Children: &body,
		}),
		Data: page,
	}

	return component
}
