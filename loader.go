package i18n

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func LoadDir[T any](dir string, unmarshal func([]byte, any) error) (*I18N[T], error) {
	in := New[T]()
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		nation, lang, err := parseFilename(info.Name())
		if err != nil {
			return err
		}
		value, err := loadFile[T](path, unmarshal)
		if err != nil {
			return err
		}
		in.Set(nation, lang, value)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return in, nil
}

func parseFilename(filename string) (nation, lang string, err error) {
	split := strings.Split(filename, ".")
	if len(split) < 2 {
		return "", "", fmt.Errorf("invalid filename: %s", filename)
	}
	return split[0], split[1], nil
}

func loadFile[T any](path string, unmarshal func([]byte, any) error) (rs T, err error) {
	f, err := os.Open(path)
	if err != nil {
		return rs, err
	}
	defer func(f *os.File) {
		e := f.Close()
		if e != nil {
			err = e
		}
	}(f)

	data, err := io.ReadAll(f)
	if err != nil {
		return rs, err
	}

	if err := unmarshal(data, &rs); err != nil {
		return rs, err
	}

	return rs, nil
}
