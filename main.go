package main

import (
	. "jWeb/components"
	"jWeb/filetypes/css"
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, reader *http.Request) {
	NewPage(css.NewCss("myStyle.css"), "Hello",
		[]Templated{
			NewDiv(DivData{
				Classes: "top",
				Children: []Templated{
					NewText("hey"),
				},
			}),
			NewDiv(DivData{
				Classes: "left",
				Children: []Templated{
					NewText("hey"),
					NewText("What's up?"),
				},
			}),
			NewDiv(DivData{
				Classes: "right",
				Children: []Templated{
					NewText("llalalla"),
				},
			}),
		}).RenderTemplate(writer)
}

func main() {
	http.HandleFunc("/", viewHandler)

	cssPath := "/" + css.GetBaseFolder()

	http.Handle(cssPath, http.StripPrefix(cssPath, http.FileServer(http.Dir(css.GetBaseFolder()))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
