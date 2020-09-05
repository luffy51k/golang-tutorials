package main

import "encapsulation"

func main() {
	e := encapsulation.Encapsulation{}

	e.Expose() // "AHHHH! I'm exposed!"

	//e.hide() // if uncommented, causes the following error
	// ./main.go:10: e.hide undefined (cannot refer
	// to unexported field or method encapsulation.
	// (*Encapsulation)."".hide)

	e.Unhide() // "Shhhh... this is super secret"
	// "...jk"
}
