package detector

import (
	"encoding/csv"
	"fmt"
	"sort"
	"strings"
)

type sniffer struct {
	sampleSize   int
	delimiter    rune
	frequencyMap frequencyMap
}

func newSniffer() *sniffer {
	delims := frequencyMap{
		',':  {},
		'\t': {},
		';':  {},
		':':  {},
	}

	return &sniffer{
		sampleSize:   15,
		delimiter:    ',',
		frequencyMap: delims,
	}
}

func (s *sniffer) analyse(r *csv.Reader) {
	for i := 0; i < s.sampleSize; i++ {
		line, err := r.Read()
		if err != nil {
			break
		}
		if len(line) > 1 {
			s.increment(r.Comma, len(line))
			continue
		}

		for potential := range s.frequencyMap {
			split := strings.Split(strings.TrimSpace(line[0]), string(potential))
			s.increment(potential, len(split))
		}
	}
}

func (s *sniffer) sniff() {
	ds := dialects{}
	for potential := range s.frequencyMap {
		p := dialect{
			delimiter:  potential,
			likelihood: []float64{},
		}
		for split, occurences := range s.frequencyMap[potential] {
			p.likelihood = append(p.likelihood, float64(split)/float64(occurences))
		}
		sort.Float64s(p.likelihood)
		ds = append(ds, p)
	}
	sort.Sort(ds)
	fmt.Println(ds)
	s.delimiter = ds[0].delimiter
}

func (s *sniffer) increment(r rune, amount int) {
	_, ok := s.frequencyMap[r][amount]
	if !ok {
		s.frequencyMap[r][amount] = 0
	}
	s.frequencyMap[r][amount]++
}
