package models

import (
	"io/ioutil"
)

// Article is something.
type Article struct {
	Title string
	Body  []byte
}

func (a *Article) save() error {
	filename := a.Title + ".txt"
	return ioutil.WriteFile(filename, a.Body, 0600)
}
