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
	mathrand "math/rand"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	homedir "github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/scrypt"
)

// Account contains all details for a account
type Account struct {
	AccountName string `json:"account_name"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	Notes       string `json:"notes"`
}

// CSVPassword contains csv parsed passwords
type CSVPassword struct {
	Name     string `csv:"name"`
	URL      string `csv:"url"`
	Username string `csv:"username"`
	Password string `csv:"password"`
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
	var version = "0.6.0"
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
	fmt.Println("----------------------------------------")
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

// ConfirmPrompt will prompt to user for yes or no
func ConfirmPrompt(message string) bool {
	var response string
	fmt.Print(message + " (yes/no) :")
	fmt.Scanln(&response)

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		return false
	}
}

// GeneratePassword will return Password
func GeneratePassword(length int) string {
	if length == 0 {
		length = 12
	}
	if length < 8 {
		length = 8
	}

	mathrand.Seed(time.Now().Unix())

	var (
		numberSet      = "0123456789"
		lowerCharSet   = "abcdedfghijklmnopqrst"
		upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		specialCharSet = "~!@#$%^&*()_+`-={}|[]\\:\"<>?/"
		CharSet        = lowerCharSet + upperCharSet
		// specialCharSet = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
		// allCharSet     = CharSet + specialCharSet + numberSet
	)

	var (
		minUpperCase    int    = (length - 4) / 4
		minNum          int    = (length - 4) / 4
		minSpecialChar  int    = (length - 4) / 4
		remainingLength int    = length - minSpecialChar - minNum - minUpperCase
		password        string = ""
	)

	for i := 0; i < minUpperCase; i++ {
		random := mathrand.Intn(len(upperCharSet))
		password += string(upperCharSet[random])
	}

	for i := 0; i < minNum; i++ {
		random := mathrand.Intn(len(numberSet))
		password += string(numberSet[random])
	}

	for i := 0; i < minSpecialChar; i++ {
		random := mathrand.Intn(len(specialCharSet))
		password += string(specialCharSet[random])
	}

	for i := 0; i < remainingLength; i++ {
		random := mathrand.Intn(len(CharSet))
		password += string(CharSet[random])
	}

	RunePassword := []rune(password)
	mathrand.Shuffle(len(RunePassword), func(i, j int) {
		RunePassword[i], RunePassword[j] = RunePassword[j], RunePassword[i]
	})

	password = string(RunePassword)

	return password
}

// SortSlice sort arrays
func SortSlice(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool { return slice[i] > slice[j] })
	return slice
}

// RemoveDuplicate removes duplicate from slice
func RemoveDuplicate(slice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// GetFileName simplify filenames for use (Note: only 3 char ext)
func GetFileName(filename, extension string) string {
	if len(filename) > 4 {
		if filename[len(filename)-4:] != extension {
			filename += extension
		}
	} else {
		filename += extension
	}
	return filename
}

//ReadCSV returns parsed Passwords (*Google Password csv file)
func ReadCSV(filename string) ([]CSVPassword, error) {
	passFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer passFile.Close()

	passwords := []CSVPassword{}

	if err := gocsv.UnmarshalFile(passFile, &passwords); err != nil {
		return []CSVPassword{}, err
	}
	return passwords, nil
}

/*
 * Google Password csv file structure
 *
 * name,url,username,password
 *
 * name | url | username | password
 * ---- | --- | -------- | --------
 */

// ConvertToAccount converts CSVPassword into Account
func ConvertToAccount(csvpasswords []CSVPassword) []Account {
	var accounts []Account
	for _, csvpassword := range csvpasswords {
		var account Account
		account.AccountName = csvpassword.Name
		account.UserName = csvpassword.Username
		account.Password = csvpassword.Password
		account.Notes = csvpassword.URL
		accounts = append(accounts, account)
	}
	return accounts
}
