# json2duration

duration is a Go package that provides functionality to convert integer values from JSON configurations into time.Duration values.
It supports conversions for:

- nanoseconds
- milliseconds
- seconds
- minutes
- hours

## Installation

To install the package, run:

```sh
go get github.com/AlexandrKobalt/json2duration
```

## Usage

Importing the Package

```go
import "github.com/AlexandrKobalt/json2duration"
```

## Example

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/AlexandrKobalt/json2duration"
)

func main() {
    jsonData := {"duration": 120}

    var data struct {
        Duration duration.Seconds json:"duration"
    }

    err := json.Unmarshal([]byte(jsonData), &data)
    if err != nil {
        panic(err)
    }

    fmt.Println("Duration in seconds:", data.Duration.Seconds())
    fmt.Println("Duration:", data.Duration.Duration)
}
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.
