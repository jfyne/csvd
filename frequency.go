package detector

import (
	"fmt"
)

type frequencyMap map[rune]map[int]int

func (f frequencyMap) String() string {
	out := ""
	for r := range f {
		out += fmt.Sprintf("%#U: %v\n", r, f[r])
	}
	return out
}
