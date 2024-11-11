package builder

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/AEKDA/aebuilder/internal/tags"
)

type App struct {
	Tags   []tags.Tag
	Source string
	Name   string
	Output string
}

var regEx = regexp.MustCompile(`^([^:]+):([^:]+)$`)

func New(path, output, name string, t []string) (App, error) {
	var tagsModel []tags.Tag

	for _, v := range t {
		matches := regEx.FindStringSubmatch(v)
		if len(matches) != 3 {
			return App{}, errors.New("Неудалось считать тэг")
		}
		model, err := tags.NewTag(matches[1], matches[2])
		if err != nil {
			return App{}, err
		}
		tagsModel = append(tagsModel, model)
	}

	return App{
		Source: path,
		Name:   name,
		Tags:   tagsModel,
		Output: output,
	}, nil
}

func (a *App) Run() error {
	fmt.Println(a)
	return a.ReadAst()
}
