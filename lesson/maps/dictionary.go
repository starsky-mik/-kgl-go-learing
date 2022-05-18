package maps

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	res, ok := d[key]

	if !ok {
		return "", KeyNotExistsError
	}

	return res, nil
}

func (d *Dictionary) Add(key, value string) error {
	_, err := (*d).Search(key)
	if err != KeyNotExistsError {
		return KeyAlreadyExistsError
	}

	(*d)[key] = value

	return nil
}

func (d *Dictionary) Set(key, value string) error {
	(*d)[key] = value

	return nil
}

func (d Dictionary) Remove(key string) error {
	_, err := d.Search(key)
	if err == KeyNotExistsError {
		return err
	}

	delete(d, key)
	return nil
}
