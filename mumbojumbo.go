package twine

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"math"
	"strings"
	"sync"
	"unicode"

	log "github.com/golang/glog"
)

// MumboJumbo is a spell checker.
type MumboJumbo struct {
	Metaphones map[string]map[string]struct{}
	CodeLength int
	mu         *sync.Mutex
}

// NewMumboJumbo reads from io and attempts to first
// unzip a gzip io, if it fails it just reads the file
// line by line.
func NewMumboJumbo(in io.Reader, codeLength int) (*MumboJumbo, error) {
	mj := &MumboJumbo{
		Metaphones: map[string]map[string]struct{}{},
		CodeLength: codeLength,
		mu:         &sync.Mutex{},
	}
	r := bufio.NewReader(in)
	gzipCheck, err := r.Peek(2)
	if err != nil || len(gzipCheck) < 2 {
		log.Error(err)
		return nil, err
	}
	if gzipCheck[0] == 31 && gzipCheck[1] == 139 {
		var err error
		fz, err := gzip.NewReader(r)
		if err != nil {
			log.Error(err)
		} else {
			r = bufio.NewReader(fz)
		}
	}

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	totalWords := 0
	for scanner.Scan() {
		words := mj.parseLine(scanner.Text())
		for _, word := range words {
			if len(word) == 0 {
				continue
			}
			totalWords++
			res, err := DoubleMetaphone(word, mj.CodeLength)
			if err != nil {
				log.Error(err)
				continue
			}
			if res[0] != "" {
				if _, ok := mj.Metaphones[res[0]]; !ok {
					mj.Metaphones[res[0]] = map[string]struct{}{}
				}
				mj.Metaphones[res[0]][word] = struct{}{}
			}
			if res[1] != "" {
				if _, ok := mj.Metaphones[res[1]]; !ok {
					mj.Metaphones[res[1]] = map[string]struct{}{}
				}
				mj.Metaphones[res[1]][word] = struct{}{}
			}
		}
	}
	log.V(1).Infof("total words:%d metaphones: %d", totalWords, len(mj.Metaphones))
	/*
		for k, v := range mj.Metaphones {
			log.Error(k, " ", v)
		}
	*/

	return mj, nil
}

// parseLine sanitizes input as a string and returns an array of
// lowercased words. Accounts for unicode characters by
// converting to runes.
func (m *MumboJumbo) parseLine(line string) []string {
	words := []string{}

	runes := []rune(strings.ToLower(line))
	var iter bytes.Buffer
	for _, r := range runes {
		switch {
		case unicode.IsLetter(r):
			iter.WriteRune(r)
		default:
			if iter.Len() > 0 {
				words = append(words, iter.String())
				iter.Reset()
			}
		}
	}
	if iter.Len() > 0 {
		words = append(words, iter.String())
	}

	return words
}

// sanitizeWord lowercases a string and removes any punctuations.
func (m *MumboJumbo) sanitizeWord(word string) string {
	runes := []rune(strings.ToLower(word))
	var iter bytes.Buffer
	for _, r := range runes {
		switch {
		case unicode.IsLetter(r):
			iter.WriteRune(r)
			//default:
			//	break
		}
	}

	return iter.String()
}

// Suggest takes an input word and returns the best suggestion for the word.
// If there's no suggestions an error is returned.
func (m *MumboJumbo) Suggest(input string) (string, error) {
	dm, err := DoubleMetaphone(input, m.CodeLength)
	if err != nil {
		log.Error(err)
		return "", err
	}
	suggMap := map[string]int{}
	m.mu.Lock()
	for k, _ := range m.Metaphones[dm[0]] {
		suggMap[k] = LevenshteinDistance(input, k)
	}
	m.mu.Unlock()
	if len(suggMap) == 0 {
		return "", fmt.Errorf("no suggestion")
	}

	if dm[1] != "" {
		alts, ok := m.Metaphones[dm[1]]
		if ok {
			for alt, _ := range alts {
				suggMap[alt] = LevenshteinDistance(input, alt)
			}
		}
	}

	min := math.MaxInt32
	bestWord := ""
	for word, dist := range suggMap {
		if dist < min {
			bestWord = word
			min = dist
		}
	}

	return bestWord, nil
}
