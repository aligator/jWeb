package main

import (
	. "jWeb/components"
	"jWeb/filetypes/css"
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, reader *http.Request) {
	NewPage(css.NewCss("myStyle.css"), "Hello",
		NewDiv("top", NewText("hey")),
		NewDiv("left", NewText("What's up?"), NewText("lalilu")),
		NewDiv("right", NewText("llalalla"))).RenderTemplate(writer)
}

func main() {
	http.HandleFunc("/", viewHandler)

	cssPath := "/" + css.GetBaseFolder()

	http.Handle(cssPath, http.StripPrefix(cssPath, http.FileServer(http.Dir(css.GetBaseFolder()))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
