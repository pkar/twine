# twine


Some string similarity helpers written in Go.

```go
package main

import "github.com/pkar/twine"

func main() {
	codes := twine.DoubleMetaphone("cabrillo", 4)
	// Output: [2]string{"KPRL", "KPR"}
	dist := twine.LevenshteinDistance("abc", "abd")
	// Output: 1
}

```
