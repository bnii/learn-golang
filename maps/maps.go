package maps

type Dictionary map[string]string

const (
	ErrNotFound          = DictionaryError("not found")
	ErrAlreadyExists     = DictionaryError("already exists")
	ErrUpdateNonExisting = DictionaryError("trying to update non-existing")
	ErrDeleteNonExisting = DictionaryError("trying to delete non-existing")
)

type DictionaryError string

func (d DictionaryError) Error() string {
	return string(d)
}

func (d Dictionary) Search(word string) (string, error) {
	word, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return word, nil
}

func (d Dictionary) Add(key, val string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = val
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}
	return nil
}
func (d Dictionary) Update(key, val string) error {
	_, err := d.Search(key)

	if err == ErrNotFound {
		return ErrUpdateNonExisting
	}

	d[key] = val
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrDeleteNonExisting
	case nil:
		delete(d, key)
	default:
		return err
	}
	return nil
}
