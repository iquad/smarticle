package smarticle

import (
	"io/ioutil"

	"github.com/iquad/smarticle/models"
)

func loadPage(title string) (*models.Article, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &models.Article{Title: title, Body: body}, nil
}

func main() {
	a1 := &Article{Title: "Test Article", Body: []byte("Lorem ipsum dolor sit amet.")}
	a1.save()
}
