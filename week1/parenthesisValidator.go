package week1

import (
	bds "github.com/nikitych1w/stepik-algorithms/basic-data-structures"
	"strconv"
)

func validateParenthesis(in string) string {
	s := bds.Stack{}

	symb := make(map[string]string)
	symb["("] = ")"
	symb["{"] = "}"
	symb["["] = "]"

	for i, el := range in {
		str := string(el)

		if str == "{" || str == "(" || str == "[" {
			s.Push(bds.Symbol{Val: str, Pos: i + 1})
		}

		if str == "}" || str == ")" || str == "]" {
			popped, ok := s.Pop()

			if !ok || str != symb[popped.Val] {
				return strconv.Itoa(i + 1)
			}
		}
	}

	if len(s) > 0 {
		el, _ := s.Pop()
		return strconv.Itoa(el.Pos)
	}

	return "Success"
}
