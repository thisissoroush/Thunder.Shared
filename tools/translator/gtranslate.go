package translator

import (
	_lng "github.com/thisissoroush/thunder.shared/primitive/enum"

	_gt "github.com/bas24/googletranslatefree"
)

func Translate(text string, source _lng.Language, destination _lng.Language) *string {

	if source == destination {
		return &text
	}

	result, err := _gt.Translate(text, string(source), string(destination))

	if err != nil {
		return nil
	}
	return &result
}
