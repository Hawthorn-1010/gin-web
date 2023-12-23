package utils

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
	"os"
)

func GenerateShortenUrl(originUrl string) string {
	urlHashBytes := sha256Of(originUrl)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encode([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}

func sha256Of(s string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(s))
	return algorithm.Sum(nil)
}

func base58Encode(s []byte) string {
	encoding := base58.BitcoinEncoding
	res, err := encoding.Encode(s)
	if err != nil {
		errors.New("fail to encode!")
		os.Exit(1)
	}
	return string(res)
}
