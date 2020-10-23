package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
    basic
	/
   /
	m
*/

type L interface {
	Life()
	Say()
}

type P struct {
	l string
}

func (x *P) String() string {
	return "p:String: " + x.l
}

func NewP(l string) *P {
	return &P{l}
}

func (p *P) Life() {
	fmt.Println("p:" + p.l)
}

func (p *P) Say() {
	fmt.Println("p:" + p.l)
}

type ExP struct {
	*P
}

func newExp(l string) *ExP {
	return &ExP{NewP(l)}
}

func (e *ExP) Say() {
	fmt.Println("e:" + e.l)
}

type SeleP struct {
	*P
	Exp L
}

// type Kind struct {
// 	name string
// }

type Kind string

// func NewKind(name string) *Kind {
// 	return &Kind{name: name}
// }
//
// func (kind *Kind) String() string  {
// 	return kind.name
// }

var (
	EQ   = "="
	COMM = "="
	SEMI = "="
)

func newSeleP(l string) *SeleP {

	x := new(SeleP)
	x.P = NewP(l)
	// x.exp = newExp(l)
	// x.l = x.exp.l
	return x
}

func main() {
	a := '孓'
	fmt.Println(a)
	fmt.Println(strconv.ParseInt("ffff", 16, 64))
}

func p(out *strings.Builder) {
	out.WriteString("xxxx")

	fmt.Println("p", out.String())
}
