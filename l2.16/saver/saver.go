package saver

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type Saver struct {
	rootDir string
}

func NewSaver(rootDir string) *Saver {
	return &Saver{rootDir: rootDir}
}

func (s *Saver) Save(resp *http.Response, rawURL string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	host := u.Hostname()

	if strings.Contains(host, ":") {
		host = strings.Split(host, ":")[0]
	}

	filePath := u.Path
	if filePath == "" || filePath == "/" {
		filePath = "/index.html"
	}

	if strings.HasSuffix(filePath, "/") {
		filePath += "index.html"
	}

	fullPath := filepath.Join(s.rootDir, host, filepath.FromSlash(filePath))

	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return err
	}

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}
