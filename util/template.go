package util

import (
	"os"
	"text/template"
)

func TemplatePlace(src, dst string, data interface{}) error {
	t, err := template.ParseFiles(src)
	if err != nil {
		return err
	}

	f, err := os.Create(dst)
	if err != nil {
		return err
	}

	err = t.Execute(f, data)
	f.Close()
	if err != nil {
		return err
	}

	return nil
}
