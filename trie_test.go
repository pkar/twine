package twine

import "testing"

func TestNewTrie(t *testing.T) {
	tr := NewTrie()
	if tr == nil {
		t.Fatal("unable to create trie")
	}
}

func TestTrieInsert(t *testing.T) {
	tr := NewTrie()
	err := tr.Insert("abc", "d")
	if err != nil {
		t.Fatal(err)
	}
}

func TestTrieGet(t *testing.T) {
	tr := NewTrie()
	tr.Insert("abc", 2)
	tr.Insert("abc", "123")
	v, err := tr.Get("abc")
	if err != nil {
		t.Fatalf("%s %#v", err, tr.Root)
	}
	if len(v) != 2 || v[0] != 2 || v[1] != "123" {
		t.Fatalf("not found got=> %+v want => [2, \"123\"]", v)
	}
}

func TestTrieDelete(t *testing.T) {
	tr := NewTrie()
	tr.Insert("abc", 2)
	tr.Insert("abc", "123")
	err := tr.Delete("abc")
	if err != nil {
		t.Fatal(err)
	}
	_, err = tr.Get("abc")
	if err == nil {
		t.Fatal("want empty")
	}
}
