package tags

import (
	"errors"

	changer "github.com/ku/go-change-case"
)

type caseType string

const (
	CamelCase    caseType = "camel"
	SnakeCase             = "snake"
	ConstantCase          = "constant"
	DotCase               = "dot"
	KebabCase             = "kebab"
	PascalCase            = "pascal"
)

func (c caseType) Convert(s string) string {
	var res string
	switch c {
	case SnakeCase:
		res = changer.Snake(s)
	case CamelCase:
		res = changer.Camel(s)
	case ConstantCase:
		res = changer.Constant(s)
	case DotCase:
		res = changer.Dot(s)
	case KebabCase:
		res = changer.Param(s)
	case PascalCase:
		res = changer.Pascal(s)
	default:
		panic("неизвестный кейс")
	}
	return res
}

func CaseFrom(s string) (caseType, error) {
	switch s {
	case "camel":
		return CamelCase, nil
	case "snake":
		return SnakeCase, nil
	case "constant":
		return ConstantCase, nil
	case "dot":
		return DotCase, nil
	case "kebab":
		return KebabCase, nil
	case "pascal":
		return PascalCase, nil
	}
	return caseType("undefined"), errors.New("undefined case")
}

type Tag struct {
	Name string
	Case caseType
}

func NewTag(n string, c string) (Tag, error) {
	ca, err := CaseFrom(c)
	return Tag{
		Name: n, Case: ca,
	}, err
}
