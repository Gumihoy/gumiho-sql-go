package main

import (
	"context"
	"fmt"
	"math"
	"strings"
	"sync"
	"sync/atomic"
	"time"
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

type A struct {
	A       string `json:"ab"`
	CreatAt string
	B       bool
}

// func (a *A) String() string {
// 	return a.a + "." + a.b
// }

var status int64

type Driver interface {
	Name() string
}

func main() {
	m := sync.Mutex{}
	m.Lock()
	m.Unlock()
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	atomic.StoreInt64(&status, 1)
	c.Broadcast()
	c.L.Unlock()
}

func IsSQLIdentifierPartWithCh(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9') ||
		(ch == '$') ||
		(ch == '_') ||
		(ch >= 0x0080 && ch <= 0xFFFF)
}

var map1 = make(map[string]bool)

func init() {
	map1["11"] = true
}

func removeQuoted(name string) string {
	if len(name) < 2 {
		return name
	}

	len := len(name)

	if (name[0] == '"' && name[len-1] == '"') ||
		(name[0] == '`' && name[len-1] == '`') {
		name = name[1 : len-1]
	}

	return name
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle:", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func changeA(a [3]int) {
	a[0] = 100
	fmt.Printf("inside: %p\n", &a)
	fmt.Printf("inside: %p\n", &a[0])
	fmt.Printf("inside: %p\n", &a[1])
}

// 1-A,2-B,..,26-Z,27-AA,28-AB,...,52-AZ,53-BA,54-BB,...,702-ZZ,703-AAA...
var T = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func IntToCode(num int) string {

	if num <= 0 {
		return ""
	}
	if num <= 26 {
		return T[num-1]
	}

	if num%26 == 0 {
		r := IntToCode(num/26-1) + "Z"
		return r
	}

	return IntToCode(num/26) + IntToCode(num%26)
}

func CodeToInt(code string) uint64 {
	r := uint64(0)
	l := len(code)

	if l == 0 {
		return 0
	}

	for i := 0; i < l; i++ {
		r = r + uint64(code[i]-'A'+1)*uint64(math.Pow(26, float64(i)))
	}
	return r
}

func p(out *strings.Builder) {
	out.WriteString("xxxx")

	fmt.Println("p", out.String())
}
