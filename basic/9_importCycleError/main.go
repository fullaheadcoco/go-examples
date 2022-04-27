package main

import (
	"9_importCycleError/parent"
)

func main() {
	p := parent.NewParent()
	c := p.CreateNewChild()
	c.PrintParentMessage()
}
