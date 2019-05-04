package components

import (
	"io"
	"net/http"
)

type Templated interface {
	RenderTemplate(writer http.ResponseWriter)
	WriteTemplate(writer io.Writer) error
}
