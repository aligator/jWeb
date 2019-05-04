package components

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type ComponentData interface{}

type Component struct {
	templ        *template.Template
	templateName string
}

func NewComponent(templ string) Component {
	return Component{
		templ:        template.Must(template.ParseFiles(templ)),
		templateName: templ,
	}
}

func (comp *Component) MustOrInternalServerError(writer http.ResponseWriter, err error) {
	if err != nil {
		http.Error(writer, "internal error", http.StatusInternalServerError)

		log.Fatal(err.Error())
	}
}

func (comp *Component) renderTemplate(writer http.ResponseWriter, compData ComponentData) {
	err := comp.writeTemplate(writer, compData)
	comp.MustOrInternalServerError(writer, err)
}

func (comp *Component) writeTemplate(writer io.Writer, compData ComponentData) error {
	err := comp.templ.Execute(writer, compData)
	return err
}
