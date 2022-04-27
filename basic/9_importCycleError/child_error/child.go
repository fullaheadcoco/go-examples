package child

import (
	"9_importCycleError/parent"
)

type Child struct {
	parent *parent.Parent
}

func (child *Child) PrintParentMessage() {
	child.parent.PrintMessage()
}

func NewChild(parent *parent.Parent) *Child {
	return &Child{parent: parent}
}
