# go-dd

A Go package similar to Laravel's `dd()` function for debugging.

## Installation

```sh
go get github.com/sherlockzzw/go-dd

```
## Usage

```go
package main

import (
	dd "github.com/sherlockzzw/go-dd"
)

func main() {
	data := []interface{}{
		"见字如晤，展信佳颜!",
		42,
		[]int{1, 2, 3},
		map[string]int{"one": 1, "two": 2},
	}

	dd.Dump(data...)
}
```

