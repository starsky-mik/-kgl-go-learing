package maps

type DictionaryError string

const (
	KeyNotExistsError     = DictionaryError("could not find key")
	KeyAlreadyExistsError = DictionaryError("this key already exists")
)

func (e DictionaryError) Error() string {
	return string(e)
}
