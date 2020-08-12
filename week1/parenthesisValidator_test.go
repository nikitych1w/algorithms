package week1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParenthesisValidator(t *testing.T) {
	var tests = []struct {
		in, out string
	}{
		{"([](){([])})", "Success"},
		{"{[()]}", "Success"},
		{"{*}", "Success"},
		{"{}", "Success"},
		{"*{}", "Success"},
		{"{{{}}}", "Success"},
		{"[]", "Success"},
		{"()[]{}", "Success"},

		{"{{[()]}", "1"},
		{"()[]}", "5"},
		{"{{[()]]", "7"},
		{"{{{[][][]", "3"},
		{"{*{{}", "3"},
		{"[[*", "2"},
		{"{{", "2"},
		{"}", "1"},
		{"{{{**[][][]", "3"},
		{"[(](){([])})", "3"},
	}

	for _, el := range tests {
		t.Run(el.in, func(t *testing.T) {
			res := validateParenthesis(el.in)
			assert.Equal(t, el.out, res, "values should be equal")
		})
	}
}
