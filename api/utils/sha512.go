package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

// GerarHash: gera um hash usando o algoritmo Sha512 para os dados fornecidos
func GerarHash(dado string) (string, error) {
	h := sha512.New()
	_, err := h.Write([]byte(dado))
	if err != nil {
		return "", err
	}
	dadoHash := h.Sum(nil)
	hashHex := hex.EncodeToString(dadoHash)
	return hashHex, nil
}
