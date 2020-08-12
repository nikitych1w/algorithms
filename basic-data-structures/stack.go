package basic_data_structures

import "fmt"

type Symbol struct {
	Val string
	Pos int
}

type Stack []Symbol

func (s *Stack) Push(val Symbol) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (Symbol, bool) {
	if len(*s) == 0 {
		return Symbol{}, false
	}

	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return val, true
}

func (s Stack) Print() {
	for _, el := range s {
		fmt.Printf("[ '%s', %d ]", el.Val, el.Pos)
	}
	fmt.Println()
}
