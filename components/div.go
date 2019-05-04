package components

import (
	"io"
	"net/http"
)

type DivData struct {
	Classes  string
	Children []Templated
}

type DivContainer struct {
	*Container
	Data *DivData
}

func (cont *DivContainer) WriteTemplate(writer io.Writer) error {
	return cont.Container.WriteTemplate(writer)
}

func (cont *DivContainer) RenderTemplate(writer http.ResponseWriter) {
	cont.Container.RenderTemplate(writer)
}

func NewDiv(classes string, children ...Templated) *DivContainer {
	div := &DivData{
		Classes:  classes,
		Children: children,
	}

	component := &DivContainer{
		Container: NewContainer(ContainerData{
			TemplateName: "div.html",
			Data:         &div,
			Children:     &div.Children,
		}),
		Data: div,
	}

	return component
}
