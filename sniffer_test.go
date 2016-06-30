package detector

import (
	"encoding/csv"
	"strings"
	"testing"
)

const (
	csv1 = `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	csv2 = `first_name;last_name;username
"Rob";"Pike";rob
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`
)

func TestAnalyse(t *testing.T) {
	r1 := csv.NewReader(strings.NewReader(csv1))
	s1 := newSniffer()
	s1.analyse(r1)

	r2 := csv.NewReader(strings.NewReader(csv2))
	s2 := newSniffer()
	s2.analyse(r2)

	for k, v := range s1.frequencyMap[','] {
		if k != 3 && v != 4 {
			t.Fail()
		}
	}

	for k, v := range s2.frequencyMap[';'] {
		if k != 3 && v != 4 {
			t.Fail()
		}
	}
}

func TestSniff(t *testing.T) {
	r1 := csv.NewReader(strings.NewReader(csv1))
	s1 := newSniffer()
	if s1.analyse(r1) != ',' {
		t.Fail()
	}

	r2 := csv.NewReader(strings.NewReader(csv2))
	s2 := newSniffer()
	if s2.analyse(r2) != ';' {
		t.Fail()
	}
}

func TestIncrement(t *testing.T) {
	s := newSniffer()

	s.increment(',', 1)

	val, ok := s.frequencyMap[','][1]
	if !ok {
		t.Fail()
	}
	if val != 1 {
		t.Fail()
	}
}
