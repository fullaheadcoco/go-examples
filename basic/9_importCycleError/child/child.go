package child

type IParent interface {
	PrintMessage()
}

type Child struct {
	parent IParent
}

func (child *Child) PrintParentMessage() {
	child.parent.PrintMessage()
}

func NewChild(parent IParent) *Child {
	return &Child{parent: parent}
}
