package main

import "fmt"

type Human struct {
	Name   string
	Height int
}

func (h Human) SayHi() {
	fmt.Println("Hihi, im ", h.Name)
}

func (h Human) SayHeigt() {
	fmt.Println("My height is ", h.Height)
}

type Action struct {
	Human
	ActionType string
}

func (a Action) SayActionType() {
	fmt.Println("My action is ", a.ActionType)
}
func main() {
	a := Action{
		Human: Human{
			Name:   "Bob",
			Height: 175,
		},
		ActionType: "Saying hi"}
	a.SayHi()
	a.SayHeigt()
	a.SayActionType()
}
