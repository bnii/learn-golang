package walk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//
// golang challenge: write a function walk(x interface{}, fn func(string))
// which takes a struct x and calls fn for all strings fields found inside.
// difficulty level: recursively.
func TestWalk(t *testing.T) {

	type Person struct {
		Age  int
		Name string
	}

	cases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		{"happy",
			struct {
				S1 string
				S2 string
			}{
				S1: "s1",
				S2: "s2",
			},
			[]string{"s1", "s2"},
		},
		{"not only string",
			struct {
				S1 string
				S2 string
				I1 int
			}{
				S1: "s1",
				S2: "s2",
				I1: 123,
			},
			[]string{"s1", "s2"},
		},
		{
			"nested fields",
			struct {
				S1 string
				P1 Person
			}{
				S1: "s1",
				P1: Person{
					Name: "p1",
					Age:  11,
				},
			},
			[]string{"s1", "p1"},
		},
		{
			"pointers",
			&Person{
				Age:  123,
				Name: "nm",
			},
			[]string{"nm"},
		},
		{
			"slices",
			[]Person{
				{Name: "p1"},
				{Name: "p2"},
			},
			[]string{"p1", "p2"},
		},
		{
			"arrays",
			[2]Person{
				{Name: "p1"},
				{Name: "p2"},
			},
			[]string{"p1", "p2"},
		},
		{
			"maps",
			map[string]string{
				"k1": "v1",
				"k2": "v2",
			},
			[]string{"v1", "v2"},
		},
	}

	for _, tst := range cases {

		t.Run(tst.Name, func(t *testing.T) {

			slc := []string{}

			Walk(tst.Input, func(s string) {
				slc = append(slc, s)
			})

			assert.ElementsMatch(t, tst.Expected, slc)
		})
	}
}
