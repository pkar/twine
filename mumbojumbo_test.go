package twine

import (
	"os"
	"strings"
	"testing"
)

func TestNewMumboJumbo(t *testing.T) {
	fName := "big.txt.gz"
	fi, err := os.Open(fName)
	if err != nil {
		t.Fatal(err)
	}
	_, err = NewMumboJumbo(fi, 4)
	if err != nil {
		t.Fatal(err)
	}
}

var suggestTests = []struct {
	in  string
	out string
}{
	{"caaat", "cat"},
	{"acaaat", ""},
	{"france", "français"},
}

func TestSuggest(t *testing.T) {
	words := "filipowicz français cat"
	r := strings.NewReader(words)
	mj, err := NewMumboJumbo(r, 4)
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range suggestTests {
		res, err := mj.Suggest(tt.in)
		if err != nil && tt.out != "" {
			t.Error(err)
		}
		if tt.out != res {
			t.Errorf("parseLine(%s) => %s, want %s", tt.in, res, tt.out)
		}
	}
}

var parseLineTests = []struct {
	in  string
	out []string
}{
	{"a b c d", []string{"a", "b", "c", "d"}},
	{"", []string{}},
	{"Schübler d.? \"' e fff français", []string{"schübler", "d", "e", "fff", "français"}},
	{"records--of", []string{"records", "of"}},
}

func TestParseLine(t *testing.T) {
	mj := &MumboJumbo{}
	for _, tt := range parseLineTests {
		res := mj.parseLine(tt.in)
		if len(res) != len(tt.out) {
			t.Errorf("parseLine(%s) => %v, want %v", tt.in, res, tt.out)
			continue
		}
		for i, out := range res {
			if out != res[i] {
				t.Errorf("parseLine(%s) => %v, want %v", tt.in, res, tt.out)
				break
			}
		}
	}
}

var sanitizeWordTests = []struct {
	in  string
	out string
}{
	{"", ""},
	{"Schübler?", "schübler"},
	{".Schübler?", "schübler"},
	{".Schübler!", "schübler"},
	{".Schübler ", "schübler"},
	{"records--of", "recordsof"},
}

func TestSanitizeWord(t *testing.T) {
	mj := &MumboJumbo{}
	for _, tt := range sanitizeWordTests {
		res := mj.sanitizeWord(tt.in)
		if res != tt.out {
			t.Errorf("sanitizeWord(%s) => %s, want %s", tt.in, res, tt.out)
			break
		}
	}
}
