package utils

import (
	"crypto/sha1"
	"encoding/base64"
)

func RandomURL(id string, url string, createdAt string) string {
	data := []byte(id + url + createdAt)

	hasher := sha1.New()
	hasher.Write(data)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))[:5]
}
