# go-csv-detector

This library automatically detects the CSV delimiter and returns a `*csv.Reader` instance.

## Usage

```go
package main

import (
    "os"

    "github.com/jfyne/go-csv-detector"
)

func main() {
    file, err := os.OpenFile("example.csv", os.O_READONLY, os.ModePerm)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Get regular old "encoding/csv" with its delimiter detected.
    reader := detector.NewReader(file)

    // Use as usual
    reader.ReadAll()
}
```
