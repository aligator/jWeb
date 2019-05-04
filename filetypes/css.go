package filetypes

import (
	"log"
	"regexp"
)

type Css struct {
	filename string
}

var baseRegexp = regexp.MustCompile("^([a-zA-Z0-9]+/)$")
var filenameRegexp = regexp.MustCompile("^([a-zA-Z0-9]+.css)$")
var baseFolder = "css/"

func NewCss(filename string) *Css {
	if filenameRegexp.MatchString(filename) {
		return &Css{
			filename: filename,
		}
	} else {
		log.Fatal(filename, " is not a valid filename for a css file")
		return nil
	}
}

func GetCssBaseFolder() string {
	return baseFolder
}

func SetCssBaseFolder(folderPath string) {
	if baseRegexp.MatchString(baseFolder) {
		baseFolder = folderPath
	} else {
		log.Fatal(folderPath, " is not a valid pathname")
	}
}

func (css *Css) GetPath() string {
	return baseFolder + "/" + css.filename
}
