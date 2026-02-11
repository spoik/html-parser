package stringreader

import "fmt"

type Reader struct {
	string   string
	position int
	length   int
}

func New(string string) *Reader {
	return &Reader{
		string:   string,
		length: len(string),
		position: -1,
	}
}

func (sr *Reader) String() string {
	return sr.string
}

func (sr *Reader) Position() int {
	return sr.position
}

func (sr *Reader) AtEnd() bool {
	return sr.position == sr.length
}

func (sr *Reader) ReadNext() (byte, error) {
	err := sr.advancePosition()

	if err != nil {
		return 0, err
	}

	return sr.string[sr.position], nil
}

func (sr *Reader) advancePosition() error {
	if sr.position == sr.length {
		return fmt.Errorf("At the end of the string.")
	}

	sr.position++
	return nil
}
