package utils

import "testing"

func TestGenerator(t *testing.T) {
	str := ""
	shortUrl := GenerateShortenUrl(str)
	t.Log(shortUrl)
}
