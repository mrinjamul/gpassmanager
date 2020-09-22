// Package gpm ...
/*
Copyright © 2020 Injamul Mohammad Mollah

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package gpm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/scrypt"
)

// Account contains all details for a account
type Account struct {
	AccountName string
	UserName    string
	Email       string
	Phone       string
	Password    string
	Notes       string
}

// DatabaseFile is the db file
var DatabaseFile string = GetHomeDir() + "/.config/gpassmanager/gpassmanager.bin"
var appname = "gpassmanager"

// GetHomeDir returns homedir
func GetHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory.")
	}
	return home
}

// GetVersion returns version name, and code
func GetVersion() string {
	var version = "0.2.0"
	return version
}

// Encrypt returns encrypted data and errors
func Encrypt(key, data []byte) ([]byte, error) {
	key, salt, err := deriveKey(key, nil)
	if err != nil {
		return nil, err
	}
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	ciphertext = append(ciphertext, salt...)
	return ciphertext, nil
}

// Decrypt returns decrypted data and errors
func Decrypt(key, data []byte) ([]byte, error) {
	salt, data := data[len(data)-32:], data[:len(data)-32]
	key, _, err := deriveKey(key, salt)
	if err != nil {
		return nil, err
	}
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// deriveKey hash the RAW password
func deriveKey(password, salt []byte) ([]byte, []byte, error) {
	if salt == nil {
		salt = make([]byte, 32)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}
	key, err := scrypt.Key(password, salt, 1<<16, 8, 1, 32)
	if err != nil {
		return nil, nil, err
	}
	return key, salt, nil
}

// VerifyKey verify the password
func VerifyKey(key, data []byte) bool {
	_, err := Decrypt(key, data)
	if err != nil {
		return false
	}
	return true
}

// CreateDatabase create folder and the file to save data
func CreateDatabase() error {
	dbLoc := GetHomeDir() + "/.config/gpassmanager/"

	err := os.MkdirAll(dbLoc, 0755)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(DatabaseFile, nil, 0644)
	if err != nil {
		return err
	}
	return nil
}

// LineBreak prints lots of '-'
func LineBreak() {
	fmt.Println("------------------------------")
}

// SavePasswords save all password
func SavePasswords(key []byte, accounts []Account) error {
	bytePasswords, err := json.Marshal(accounts)
	if err != nil {
		return err
	}
	ciphertext, err := Encrypt(key, bytePasswords)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(DatabaseFile, ciphertext, 0644)
	if err != nil {
		return err
	}
	return nil
}

// ReadPasswords decrypt all password
func ReadPasswords(key []byte) ([]Account, error) {
	data, err := ioutil.ReadFile(DatabaseFile)
	if err != nil {
		return []Account{}, err
	}
	plaintext, err := Decrypt(key, data)
	if err != nil {
		return []Account{}, err
	}
	var accounts []Account
	if err := json.Unmarshal(plaintext, &accounts); err != nil {
		return []Account{}, err
	}

	return accounts, nil
}

// RemoveAccount removes an acccount
func RemoveAccount(slice []Account, s int) []Account {
	return append(slice[:s], slice[s+1:]...)
}
