package util

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcutil/base58"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.Encode(bytes)
	return string(encoding)
}

func GenerateShortUrl(initialLink string) string {
	if initialLink == "" {
		return ""
	}
	urlHashBytes := sha256Of(initialLink)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
