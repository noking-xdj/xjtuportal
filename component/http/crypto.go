package http

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"errors"
)

var portalCryptoKey = []byte("1234567890000000")

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7Unpad(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("http/crypto: empty data")
	}
	padding := int(data[len(data)-1])
	if padding == 0 || padding > len(data) {
		return nil, errors.New("http/crypto: invalid padding")
	}
	for _, v := range data[len(data)-padding:] {
		if int(v) != padding {
			return nil, errors.New("http/crypto: invalid padding")
		}
	}
	return data[:len(data)-padding], nil
}

func EncryptPortalPayload(data []byte) (string, error) {
	if len(data) == 0 {
		return "", nil
	}
	block, err := aes.NewCipher(portalCryptoKey)
	if err != nil {
		return "", err
	}
	paddedData := pkcs7Pad(data, block.BlockSize())
	encryptedData := make([]byte, len(paddedData))
	mode := cipher.NewCBCEncrypter(block, portalCryptoKey)
	mode.CryptBlocks(encryptedData, paddedData)
	return hex.EncodeToString(encryptedData), nil
}

func DecryptPortalPayload(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var encryptedHex string
	if err := json.Unmarshal(data, &encryptedHex); err != nil {
		encryptedHex = string(data)
	}
	encryptedData, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(portalCryptoKey)
	if err != nil {
		return nil, err
	}
	if len(encryptedData)%block.BlockSize() != 0 {
		return nil, errors.New("http/crypto: invalid ciphertext length")
	}
	decryptedData := make([]byte, len(encryptedData))
	mode := cipher.NewCBCDecrypter(block, portalCryptoKey)
	mode.CryptBlocks(decryptedData, encryptedData)
	return pkcs7Unpad(decryptedData)
}
