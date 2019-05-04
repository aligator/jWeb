package components

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"regexp"
)

type ComponentData interface{}

type Component struct {
	templateName string
}

var tagnameRegexp = regexp.MustCompile("^[a-zA-Z0-9]+$")
var templates = template.New("none")

func (comp *Component) MustOrInternalServerError(writer http.ResponseWriter, err error) {
	if err != nil {
		http.Error(writer, "internal error", http.StatusInternalServerError)

		log.Fatal(err.Error())
	}
}

func (comp *Component) WriteTemplate(writer io.Writer) error {
	return comp.writeTemplate(writer, nil)
}

func (comp *Component) RenderTemplate(writer http.ResponseWriter) {
	comp.renderTemplate(writer, nil)
}

func (comp *Component) renderTemplate(writer http.ResponseWriter, compData ComponentData) {
	err := comp.writeTemplate(writer, compData)
	comp.MustOrInternalServerError(writer, err)
}

func (comp *Component) writeTemplate(writer io.Writer, compData ComponentData) error {
	err := templates.ExecuteTemplate(writer, comp.templateName, compData)
	return err
}

// Todo: combine the new Methods, as there is duplicated code
func NewComponent(filename string) *Component {

	if templ := templates.Lookup(filename); templ != nil {
		return &Component{
			templateName: filename,
		}
	} else {
		template.Must(templates.ParseFiles(filename))
	}

	return &Component{
		templateName: filename,
	}
}

func NewDirectTemplateComponent(name string, templStr string) *Component {
	if templ := templates.Lookup(name); templ != nil {
		return &Component{
			templateName: name,
		}
	} else {
		template.Must(templates.New(name).Parse(templStr))
	}

	return &Component{
		templateName: name,
	}
}

func Br() *Component {
	return NewDirectTemplateComponent("br", "<br>")
}
