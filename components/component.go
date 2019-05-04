package components

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type ComponentData interface{}

type Component struct {
	templates *template.Template
}

func NewComponent(templates ...string) Component {
	return Component{template.Must(template.ParseFiles(templates...))}
}

func (comp *Component) MustOrInternalServerError(writer http.ResponseWriter, err error) {
	if err != nil {
		http.Error(writer, "internal error", http.StatusInternalServerError)

		log.Fatal(err.Error())
	}
}

func (comp *Component) renderTemplate(writer http.ResponseWriter, tmpl string, compData ComponentData) {
	err := comp.writeTemplate(writer, tmpl, compData)
	comp.MustOrInternalServerError(writer, err)
}

func (comp *Component) writeTemplate(writer io.Writer, tmpl string, compData ComponentData) error {
	err := comp.templates.ExecuteTemplate(writer, tmpl, compData)
	return err
}
