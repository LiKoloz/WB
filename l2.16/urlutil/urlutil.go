package urlutil

import (
	"net/url"
	"strings"
)

// Normalize приводит URL к единому виду: без фрагмента, схема в нижнем регистре, убирает лишние слеши
func Normalize(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	u.Fragment = ""
	if u.Scheme != "" {
		u.Scheme = strings.ToLower(u.Scheme)
	}
	if u.Host != "" {
		u.Host = strings.ToLower(u.Host)
	}
	if u.Path == "/" {
		u.Path = ""
	}
	return u.String(), nil
}

// SameDomain проверяет, принадлежит ли URL тому же домену, что и base (без учёта поддоменов)
func SameDomain(base, other string) bool {
	u1, err := url.Parse(base)
	if err != nil {
		return false
	}
	u2, err := url.Parse(other)
	if err != nil {
		return false
	}
	return u1.Hostname() == u2.Hostname()
}
