# go-csv-detector

This library automatically detects the CSV delimiter and returns a `*csv.Reader` instance.

## Caveats

This is quite a simple implementation, the following caveats should be considered:

- The whole reader will be read into memory before returning the CSV reader, so it is not suitable for massive CSV's.
- At the moment it checks for the following delimiters `, \t ; :`
- It's not bullet proof! Quick and simple.

## Usage

```go
package main

import (
    "os"
    "strings"

    "github.com/jfyne/go-csv-detector"
)

func main() {
    example := `first_name;last_name;username
"Rob";"Pike";rob
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`

    // Get regular old "encoding/csv" with its delimiter detected.
    reader := detector.NewReader(strings.NewReader(example))

    // reader.Comma will be ';' instead of the default. Use as usual
    reader.ReadAll()
}
```
