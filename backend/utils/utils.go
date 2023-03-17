package utils

import "math/rand"

var runes = []rune("0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

func RandomURL(size int) string {
	str := make([]rune, size)
	len := len(runes)

	for i := range str {
		str[i] = runes[rand.Intn(len)]
	}

	return string(str)
}

func StringCompressor(s string) string {

}
