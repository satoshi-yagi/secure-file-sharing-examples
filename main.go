package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ProtonMail/go-crypto/openpgp"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("please put 2 arguments: gpgFile and password")
		return
	}
	decrypter(args[0], args[1])
}

func decrypter(fileName, password string) {
	// pre-check
	if _, err := os.Stat(fileName); err != nil {
		fmt.Printf("File name '%s' not found\n", fileName)
		return
	}

	targetFileName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	if _, err := os.Stat(targetFileName); err == nil {
		fmt.Printf("File name '%s' exists already\n", targetFileName)
		return
	}

	f, err := os.Create(targetFileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	passphrase := []byte(password)
	dataIO, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer dataIO.Close()

	md, err := openpgp.ReadMessage(dataIO, nil, func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		return passphrase, nil
	}, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bytes, err := ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = f.WriteString(string(bytes))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
