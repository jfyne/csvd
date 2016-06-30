package detector

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
)

// NewReader returns a csv.Reader with the delimiter detected.
func NewReader(r io.Reader) *csv.Reader {
	b, _ := ioutil.ReadAll(r)
	reader := bytes.NewReader(b)

	csvReader := csv.NewReader(reader)

	sniff := newSniffer()
	sniff.analyse(csvReader)
	sniff.sniff()

	reader.Seek(0, 0)

	output := csv.NewReader(reader)
	output.Comma = sniff.delimiter

	return output
}
