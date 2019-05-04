package components

import (
	"bufio"
	"bytes"
	"html/template"
	"io"
	"net/http"
)

type ContainerData struct {
	Data     interface{}
	Children *[]Templated
}

type finalContainerData struct {
	Data     interface{}
	Children template.HTML
}

type Container struct {
	Component
	Data ContainerData
}

func (cont *Container) WriteTemplate(writer io.Writer) error {
	renderedContent, err := cont.renderContainer(writer)

	if err != nil {
		return err
	}

	return cont.writeTemplate(writer, finalContainerData{
		Data:     &cont.Data.Data,
		Children: template.HTML(renderedContent),
	})
}

func (cont *Container) RenderTemplate(writer http.ResponseWriter) {
	renderedContent, err := cont.renderContainer(writer)

	cont.MustOrInternalServerError(writer, err)

	cont.renderTemplate(writer, finalContainerData{
		Data:     &cont.Data.Data,
		Children: template.HTML(renderedContent),
	})
}

func (cont *Container) renderContainer(writer io.Writer) (string, error) {
	var bodyBuffer bytes.Buffer
	var bodyWriter = bufio.NewWriter(&bodyBuffer)

	for _, comp := range *cont.Data.Children {
		err := comp.WriteTemplate(bodyWriter)
		if err != nil {
			return "", err
		}
	}

	bodyWriter.Flush()
	return bodyBuffer.String(), nil
}

func NewContainer(templateName string, compData ContainerData) *Container {
	component := &Container{
		Component: NewComponent(templateName),
		Data:      compData,
	}

	return component
}
