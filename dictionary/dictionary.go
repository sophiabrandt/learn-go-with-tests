package dictionary

// custom error for a dictionary
type dictionaryError string

// Error implement the error interface
func (de dictionaryError) Error() string {
	return string(de)
}

// ErrNotFound is returned when a word is not found in the dictionary
const ErrNotFound = dictionaryError("could not find the word you're looking for")

// ErrWordExists is returned when a word already exists in the dictionary
const ErrWordExists = dictionaryError("cannot add word because it already exists")

// ErrorDoesNotExist is returned when a word does not exist in the dictionary
const ErrWordDoesNotExist = dictionaryError("cannot update word because it does not exist")

// Dictionary type that contains words
type Dictionary map[string]string

// Search returns a word from the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add adds a word to the dictionary
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

// Update changes the definition of a word in the dictionary
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

// Delete removes a word from the dictionary
func (d Dictionary) Delete(word string) error {
	delete(d, word)
	return nil
}
