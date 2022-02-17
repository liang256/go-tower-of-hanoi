package main

import (
	"fmt"
	"go-tower-of-hanoi/src/stack"
)

func main() {
	pillar_a := stack.New()
	pillar_b := stack.New()
	pillar_c := stack.New()
	pillar_a.Push(5)
	pillar_a.Push(4)
	pillar_a.Push(3)
	pillar_a.Push(2)
	pillar_a.Push(1)

	printPillars(pillar_a, pillar_b, pillar_c)
	fmt.Println("--")

	hanoi(5, pillar_a, pillar_b, pillar_c)

	// fmt.Println("--")
	// printPillars(pillar_a, pillar_b, pillar_c)
}

func hanoi(layer int, from *stack.Stack, tmp *stack.Stack, to *stack.Stack) {
	printPillars(from, tmp, to)

	if layer == 0 {
		return
	}
	/** base case : when layer == 1, it's our base case **/

	/** step01: move the above plate to the pillar_mid, to clear the room for the plate below **/
	hanoi(layer-1, from, to, tmp)

	/** step02: main target **/
	if ele, ok := from.Pop(); ok {
		to.Push(ele)
	}

	/** step03: move the original above plate back **/
	hanoi(layer-1, tmp, from, to)
}

func printPillars(a *stack.Stack, b *stack.Stack, c *stack.Stack) {
	space := 13
	printPillar("A", a, space)
	printPillar("B", b, space)
	printPillar("C", c, space)
	fmt.Printf("\n")
}

func printPillar(name string, p *stack.Stack, space int) {

	var remain int
	if p.Size() > 0 {
		remain = space - 2*p.Size() - 1
	} else {
		remain = space - 2
	}

	fmt.Printf(
		"%s:%v%*s",
		name,
		p.ToSlice(), remain, "",
	)
}
