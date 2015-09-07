package main

import (
	"github.com/go-chef/chef"
	"fmt"
)

type KeyMap map[string]*EncryptedDataBagItemKey

type EncryptedDataBagItem struct {
	Id   string
	Keys KeyMap
}

type EncryptedDataBagItemKey struct {
	Cipher        string
	EncryptedData string
	Iv            string
	Version       float64
}

func NewEncryptedDataBagItem(in chef.DataBagItem) (item *EncryptedDataBagItem) {
	item = new(EncryptedDataBagItem)
	item.Keys = make(map[string]*EncryptedDataBagItemKey)

	for k, v := range in.(map[string]interface{}) {
		switch k {
		case "id":
			item.Id = v.(string)
		default:
			item.Keys[k] = newEncryptedDataBagItemKey(v)
		}
	}

	return item
}

func (e *EncryptedDataBagItem) DecryptKey(bagKey string, decryptionKey []byte) (interface{}, error) {
	key, ok := e.Keys[bagKey]
	if !ok {
		return nil, fmt.Errorf("'%s' is not a key in the given databag", bagKey)
	}

	switch key.Version {
	case 1:
		return version1Decoder(decryptionKey, key.Iv, key.EncryptedData)
	default:
		return nil, fmt.Errorf("not implemented for encrypted bag version %d!", key.Version)
	}
}

func newEncryptedDataBagItemKey(in interface{}) *EncryptedDataBagItemKey {
	return &EncryptedDataBagItemKey{
		Cipher:        in.(map[string]interface{})["cipher"].(string),
		EncryptedData: in.(map[string]interface{})["encrypted_data"].(string),
		Iv:            in.(map[string]interface{})["iv"].(string),
		Version:       in.(map[string]interface{})["version"].(float64),
	}
}

