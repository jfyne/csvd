package detector

import (
	"fmt"
)

type dialect struct {
	delimiter  rune
	likelihood []float64
}

type dialects []dialect

func (d dialects) String() string {
	out := ""
	for _, dia := range d {
		out += fmt.Sprintf("%#U: %v\n", dia.delimiter, dia.likelihood)
	}
	return out
}

func (d dialects) Len() int {
	return len(d)
}

func (d dialects) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d dialects) Less(i, j int) bool {
	if len(d[i].likelihood) == 0 {
		return false
	}
	if len(d[j].likelihood) == 0 {
		return true
	}
	return d[i].likelihood[0] > d[j].likelihood[0]
}
