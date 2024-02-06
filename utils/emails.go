package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	fmt.Println("Parsing templates")
	if err != nil {
		return nil, err
	}
	return template.ParseFiles(paths...)
}
