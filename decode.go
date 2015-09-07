package main

import (
	"encoding/base64"
	"fmt"
	"crypto/sha256"
	"crypto/aes"
	"errors"
	"crypto/cipher"
	"encoding/json"
)

func decodeBase64(str string) []byte {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
	}
	return data
}

func unPKCS7Padding(data []byte) []byte {
	dataLen := len(data)
	endIndex := int(data[dataLen-1])

	if 16 > endIndex {
		return data[:dataLen-endIndex]
	}
	return nil
}

type version1Item struct {
	Content interface{} `json:"json_wrapper"`
}

func version1Decoder(key []byte, iv, encryptedData string) (interface{}, error)  {
	ciphertext := decodeBase64(encryptedData)
	initVector := decodeBase64(iv)
	keySha := sha256.Sum256(key)

	block, err := aes.NewCipher(keySha[:])
	if err != nil {
		return nil, err
	}

	if len(ciphertext) % aes.BlockSize != 0 {
		return nil, errors.New("Ciphertext is wrong length")
	}

	mode := cipher.NewCBCDecrypter(block, initVector)
	mode.CryptBlocks(ciphertext, ciphertext)

	ciphertext = unPKCS7Padding(ciphertext)

	var item version1Item
	_ = json.Unmarshal(ciphertext, &item)

	if item.Content == nil {
		return nil, fmt.Errorf("Decryption failed, check your decryption key.")
	}

	return item.Content, nil
}