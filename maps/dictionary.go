package maps

import "errors"

var (
	ErrNotFoundWord  = errors.New("could not find the word")
	ErrWordExists    = errors.New("word already defined")
	ErrWordNotExists = errors.New("word not exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	translation, ok := d[word]
	if !ok {
		return "", ErrNotFoundWord
	}

	return translation, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFoundWord:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFoundWord:
		return ErrWordNotExists
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil

}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
