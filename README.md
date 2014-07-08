# twine


Some string similarity helpers written in Go.


```bash
# get the glog package for logging.
go get github.com/golang/glog
```

```go
package main

import (
	"io"
	"flag"
	"os"

	"github.com/pkar/twine"
)

func main() {
	path := flag.String("file", "STDIN", "path to file, if not given os.Stdin assumed")
	flag.Parse()

	var ioIn io.Reader
	if *path == "STDIN" {
		ioIn = os.Stdin
	} else {
		var err error
		ioIn, err = os.Open(fName)
		if err != nil {
			os.Exit(1)
		}
	}

	mj, err := NewMumboJumbo(ioIn)
	mj.Suggest("caaat")
	// Output: cat

	codes := twine.DoubleMetaphone("cabrillo", 4)
	// Output: [2]string{"KPRL", "KPR"}
	dist := twine.LevenshteinDistance("abc", "abd")
	// Output: 1
}

```
