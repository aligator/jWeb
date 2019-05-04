package main

import (
	. "jWeb/components"
	"jWeb/filetypes/css"
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, reader *http.Request) {

	var leftComps []Templated

	for i := 0; i < 5; i++ {
		leftComps = append(leftComps, P("What's up?"))
	}

	NewPage(css.NewCss("myStyle.css"), "Hello",
		Div("top", H1("hey")),

		Div("left", leftComps...),

		Div("right", P("llalalla"))).RenderTemplate(writer)
}

func main() {
	http.HandleFunc("/", viewHandler)

	cssPath := "/" + css.GetBaseFolder()

	http.Handle(cssPath, http.StripPrefix(cssPath, http.FileServer(http.Dir(css.GetBaseFolder()))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
