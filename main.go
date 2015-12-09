package main

import (
	"fmt"
	"github.com/go-chef/chef"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		printAndExit(fmt.Errorf("No argument given"))
	}
	variablePath := os.Args[1]
	bagName, bagItem, keyName, err := parsePath(variablePath)
	if err != nil {
		printAndExit(fmt.Errorf("Path '%s' is invalid", variablePath))
	}

	nodeName := readEnvVar("CHEF_NODE_NAME")
	clientKeyPath := readEnvVar("CHEF_CLIENT_KEY_PATH")
	serverUrl := readEnvVar("CHEF_SERVER_URL")
	decryptionKeyPath := readEnvVar("CHEF_DECRYPTION_KEY_PATH")
        sslSkipCheck  := readEnvVar("CHEF_SSL_CHECK")

	key, err := ioutil.ReadFile(clientKeyPath)
	if err != nil {
		printAndExit(err)
	}

	decryptionKey, err := ioutil.ReadFile(decryptionKeyPath)
	if err != nil {
		printAndExit(err)
	}

        var sslOption bool = false

        if (sslSkipCheck == "true") {
         sslOption = true
        }

	client, err := chef.NewClient(&chef.Config{
		Name:    nodeName,
		Key:     string(key),
		SkipSSL: sslOption,
		BaseURL: fmt.Sprintf("%s/foo", serverUrl), // /foo is needed here because of how URLs are parsed by go-chef
		SkipSSL: (os.Getenv("CHEF_SKIP_SSL") == "1"),
	})
	if err != nil {
		printAndExit(err)
	}

	item, err := client.DataBags.GetItem(bagName, bagItem)
	if err != nil {
		printAndExit(err)
	}

	encrypted := NewEncryptedDataBagItem(item)

	unencrypted, err := encrypted.DecryptKey(keyName, decryptionKey)
	if err != nil {
		printAndExit(err)
	}

	fmt.Print(unencrypted)
}
