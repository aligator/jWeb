package main

import (
	. "jWeb/components"
	"jWeb/filetypes/css"
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, reader *http.Request) {
	NewPage(css.NewCss("myStyle.css"), "Hello",
		Div("top", H1("hey")),

		Div("left",
			P("What's up?"),
			Br(), Br(), Br(),
			P("lalilu")),

		Div("right", P("llalalla"))).RenderTemplate(writer)
}

func main() {
	http.HandleFunc("/", viewHandler)

	cssPath := "/" + css.GetBaseFolder()

	http.Handle(cssPath, http.StripPrefix(cssPath, http.FileServer(http.Dir(css.GetBaseFolder()))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
