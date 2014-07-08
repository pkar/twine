package twine

import (
	"fmt"
	"sync"
)

// TrieNode containes pointers to children nodes and a key value.
// There is a bool signifying the end of a word and a counter
// to keep track of the number of prefixes for a node.
type TrieNode struct {
	key      rune
	values   []interface{}
	children map[rune]*TrieNode
	isEnd    bool
	prefixes uint32 // how many words have this prefix
}

// Trie holds the root trie node
type Trie struct {
	Root *TrieNode
	mu   *sync.Mutex
}

// NewTrieNode initializes a node, as well
// as a map of runes and nodes.
func NewTrieNode(key rune) *TrieNode {
	return &TrieNode{
		key:      key,
		children: map[rune]*TrieNode{},
		values:   []interface{}{},
	}
}

// NewTrie initializes an empty root node.
func NewTrie() *Trie {
	return &Trie{Root: &TrieNode{children: map[rune]*TrieNode{}}, mu: &sync.Mutex{}}
}

// Insert updates the trie with key and appends a value in
// the end node.
func (t *Trie) Insert(key string, value interface{}) error {
	it := t.Root
	t.mu.Lock()
	for _, runeChar := range []rune(key) {
		if found, ok := it.children[runeChar]; ok {
			it = found
			it.prefixes++
		} else {
			found := NewTrieNode(runeChar)
			it.children[runeChar] = found
			it = found
			it.prefixes = 1
		}
	}
	t.mu.Unlock()
	it.isEnd = true
	it.values = append(it.values, value)
	return nil
}

// Get searches the trie and returns any values stored in the
// end node or a not found error.
func (t *Trie) Get(key string) ([]interface{}, error) {
	it := t.Root
	t.mu.Lock()
	defer t.mu.Unlock()

	for _, runeChar := range []rune(key) {
		if found, ok := it.children[runeChar]; ok {
			it = found
		} else {
			return nil, fmt.Errorf("not found")
		}
	}
	if !it.isEnd {
		return nil, fmt.Errorf("not found")
	}
	return it.values, nil
}

// Delete removes a given path within the trie by first finding
// if it exists. It removes all values at the end node.
func (t *Trie) Delete(key string) error {
	// check if there first
	_, err := t.Get(key)
	if err != nil {
		return err
	}

	it := t.Root
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, runeChar := range []rune(key) {
		if it.children[runeChar].prefixes == 1 {
			delete(it.children, runeChar)
			return nil
		} else {
			it.children[runeChar].prefixes--
			it = it.children[runeChar]
		}
	}
	it.isEnd = false
	return nil
}
