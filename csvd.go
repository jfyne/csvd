package csvd

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
)

// NewReader returns a csv.Reader with the delimiter detected. As a second argument you
// can pass in a *Sniffer instance to use instead of the defaults, this can provide a different
// set of delimiters to look for.
func NewReader(r io.Reader, s ...*Sniffer) *csv.Reader {
	var sniffer *Sniffer
	if len(s) != 0 {
		sniffer = s[0]
	} else {
		sniffer = defaultSniffer()
	}
	return newReaderFromSniffer(r, sniffer)
}

// NewSnifferReader returns a csv.Reader using a provided sniffer.
func newReaderFromSniffer(r io.Reader, s *Sniffer) *csv.Reader {
	b, _ := ioutil.ReadAll(r)
	reader := bytes.NewReader(b)

	csvReader := csv.NewReader(reader)

	s.analyse(csvReader)

	reader.Seek(0, 0)

	output := csv.NewReader(reader)
	output.Comma = s.delimiter

	return output
}
