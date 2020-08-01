package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IHuman interface {
	SayHi()
}

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段
	school string
}

type Employee struct {
	Human   //匿名字段
	company string
}

//在human上面定义了一个method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h *Student) SayHi() {
	fmt.Println(" Studio ....")
}

func (h *Employee) SayHi() {
	fmt.Println(" Employee ....")
}

type S struct {
	h IHuman
}

type S2 struct {
	S
}

func main() {
	//message := make([]rune, 10)
	var m strings.Builder
	m.WriteString("line ")
	m.WriteRune(1)
	messa := "line " + strconv.Itoa(1) + ", col " + strconv.Itoa(0) + ""
	fmt.Println(messa)
	fmt.Println(m.String())
}
