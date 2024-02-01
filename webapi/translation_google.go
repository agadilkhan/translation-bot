package webapi

import (
	"fmt"
	translator "github.com/Conight/go-googletrans"
)

type Translation struct {
	Source      string
	Destination string
	Original    string
	Translation string
}

type TranslationWebApi struct {
	conf translator.Config
}

func New() *TranslationWebApi {
	conf := translator.Config{
		UserAgent:   []string{"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1"},
		ServiceUrls: []string{"translate.google.com"},
	}

	return &TranslationWebApi{
		conf: conf,
	}
}

func (t *TranslationWebApi) Translate(translation Translation) (Translation, error) {
	trans := translator.New(t.conf)

	res, err := trans.Translate(translation.Original, translation.Source, translation.Destination)
	if err != nil {
		return Translation{}, fmt.Errorf("TranslationWebAPI - Translate - trans.Translate: %w", err)
	}

	translation.Translation = res.Text

	return translation, nil
}
