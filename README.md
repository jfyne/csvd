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
