package main

import (
	"bufio"
	"bytes"
	. "jWeb/components"
	"jWeb/filetypes/css"
	"log"
	"net/http"
)

var buffer bytes.Buffer
var needReload = true

func viewHandler(writer http.ResponseWriter, reader *http.Request) {
	if needReload {
		var leftComps []Templated

		for i := 0; i < 9999; i++ {
			leftComps = append(leftComps, P("What's up?"))
		}

		buffer = bytes.Buffer{}
		var bWriter = bufio.NewWriter(&buffer)

		NewPage(css.NewCss("myStyle.css"), "Hello",
			Div("top", H1("hey")),

			Div("left", leftComps...),

			Div("right", P("llalalla"))).WriteTemplate(bWriter)

		bWriter.Flush()

		needReload = false
	}

	writer.Write(buffer.Bytes())
}

func main() {
	http.HandleFunc("/", viewHandler)

	cssPath := "/" + css.GetBaseFolder()

	http.Handle(cssPath, http.StripPrefix(cssPath, http.FileServer(http.Dir(css.GetBaseFolder()))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
