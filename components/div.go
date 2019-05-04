package components

type DivData struct {
	Classes  string
	Children []Templated
}

type DivContainer struct {
	*Container
	Data *DivData
}

func NewDiv(classes string, children ...Templated) *DivContainer {
	div := &DivData{
		Classes:  classes,
		Children: children,
	}

	component := &DivContainer{
		Container: NewContainer(ContainerData{
			TemplateName: "div.html",
			Data:         &div,
			Children:     &div.Children,
		}),
		Data: div,
	}

	return component
}
