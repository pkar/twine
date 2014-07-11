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

	// MumboJumbo spelling suggestions using double metaphones. Provide an io.Reader
	// which by default is stdin. Text can be anything as long as there are space 
	// separated words.
	// cat filename.text | mumbo
	// or
	// mumbo -file=file.txt
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
	mj, err := twine.NewMumboJumbo(ioIn)
	mj.Suggest("caaat")
	// Output: cat

	// DoubleMetaphone is used by mumbo jumbo to encode words to 
	// a given length encoding.
	codes := twine.DoubleMetaphone("cabrillo", 4)
	// Output: [2]string{"KPRL", "KPR"}

	// LevenshteinDistance provides edit distances and is used by MumboJumbo
	// after finding a matching code.
	dist := twine.LevenshteinDistance("abc", "abd")
	// Output: 1

	// Trie is a simple trie implementation
	tr := twine.NewTrie()
	tr.Insert("abc", 2)
	tr.Insert("abc", "123")
	vals, err = tr.Get("abc")
	// Output: [2, "123"]

	err := tr.Delete("abc")
}
```
[![wercker status](https://app.wercker.com/status/7798e32da599f66f46af6c7e4a595e07/m "wercker status")](https://app.wercker.com/project/bykey/7798e32da599f66f46af6c7e4a595e07)
