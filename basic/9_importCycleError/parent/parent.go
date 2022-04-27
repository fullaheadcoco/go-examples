package parent

import (
	"9_importCycleError/child"
	//package command-line-arguments
	//	imports 10_importCycleError/parent_error
	//	imports 10_importCycleError/child_error
	//	imports 10_importCycleError/parent_error: import cycle not allowed
	//"9_importCycleError/child_error"
	"fmt"
)

type Parent struct {
	message string
}

func (parent *Parent) PrintMessage() {
	fmt.Println(parent.message)
}

func (parent *Parent) CreateNewChild() *child.Child {
	return child.NewChild(parent)
}

func NewParent() *Parent {
	return &Parent{message: "Hello World"}
}
