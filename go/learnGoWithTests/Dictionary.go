package learnGoWithTests

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	if value, ok := d[word]; ok != true {
		return "", ErrNotFound
	} else {
		return value, nil
	}
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
