package utils

import (
	"encoding/base64"
	"errors"
	"net/url"
	"strings"

	"github.com/skip2/go-qrcode"
)


func IsUrl(originalUrl string) (*url.URL, error) {
	u, err := url.ParseRequestURI(originalUrl)
	if err != nil {
		return nil, errors.New("Invalid url")
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, errors.New("Invalid protocol")
	}

	if u.Host == "" {
		return nil, errors.New("Invalid domain")
	}

	if !strings.Contains(u.Host, ".") {
		return nil, errors.New("Invalid domain")
	}

	return u, nil
}

func GetQrCode(url string) (string, error) {
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return "", nil
	}
	return base64.StdEncoding.EncodeToString(png), nil
}