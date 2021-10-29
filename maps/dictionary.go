package maps

const (
	// having precise errors gives you more informatino on what went wrong
	ErrNotFound = DictionaryErr("could not find word")
	ErrWordExists = DictionaryErr("already exists. did not add")
	ErrWordDoesNotExist = DictionaryErr("cannot update because it dose not exist")
)

// create DictionaryErr type which implements error interface
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// create Dictionary type as a thin wrapper around map
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]
	// found differentiates between a word that doesn't exist and one that doesn't have a definition
	if !found {
		return "", ErrNotFound
	}
	return definition, nil
}

// only add new words, don't modify
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
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
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	// built in delete function on maps
	delete(d, word)
}