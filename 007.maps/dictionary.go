package maps

type Dictionary map[string]string

const (
	ErrNotFound      = DictionaryError("Word not found")
	ErrWordExists    = DictionaryError("Word already present")
	ErrWordNotExists = DictionaryError("Word not present to update")
)

type DictionaryError string

func (d DictionaryError) Error() string {
	return string(d)
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if ok {
		return val, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Update(key, val string) error {

	_, err := d.Search(key)

	switch err {
	case nil:
		d[key] = val
	case ErrNotFound:
		return ErrWordNotExists
	default:
		return err

	}

	return nil
}

func (d Dictionary) Add(key, val string) error {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = val
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
