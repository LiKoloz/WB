package downloader

import (
	"fmt"
	"net/http"
)

// Downloader отвечает за выполнение HTTP-запросов к указанным URL.
type Downloader struct {
	client *http.Client
}

// NewDownloader создаёт новый экземпляр Downloader с заданным HTTP-клиентом.
func NewDownloader(client *http.Client) *Downloader {
	return &Downloader{client: client}
}

// Get выполняет GET-запрос к указанному URL и возвращает HTTP-ответ.
func (d *Downloader) Get(rawURL string) (*http.Response, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Go-Mirror/1.0)")
	resp, err := d.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		resp.Body.Close()
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	return resp, nil
}
