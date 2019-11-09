package main

import (
	"errors"
	"fmt"
)

// 定义一个字典
type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

// 给字典定义一个方法
// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }
func (d Dictionary) Search(word string) (string, error) {
	defintion, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return defintion, nil
}

func (d Dictionary) Add(word string, defintion string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = defintion
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func main() {
	dictionary := Dictionary{"test": "this is just a test"}

	_, ok := dictionary.Search("inknown")
	fmt.Println(ok)
}
