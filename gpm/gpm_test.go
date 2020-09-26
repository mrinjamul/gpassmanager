package gpm

import (
	"testing"
)

// TestGetVersion tests
func TestGetVersion(t *testing.T) {
	out := GetVersion()
	if out == "" || len(out) == 0 {
		t.Errorf("Want strings but got nil")
	}
}

// TestGetHomeDir tests
func TestGetHomeDir(t *testing.T) {
	out := GetHomeDir()
	if out == "" || len(out) == 0 {
		t.Errorf("Want strings but got nil")
	}
}

// TestderiveKey hash the RAW password
func TestDeriveKey(t *testing.T) {
	password := []byte("password")
	key, salt, err := deriveKey(password, nil)
	if len(key) == 0 || len(salt) == 0 {
		t.Errorf("Want strings but got nil")
	}
	if err != nil {
		t.Errorf("Want nil but got errors")
	}
}

// TestEncrypt tests
func TestEncrypt(t *testing.T) {
	password := []byte("password")
	data := []byte("Data")
	ciphertext, err := Encrypt(password, data)
	if ciphertext == nil {
		t.Errorf("Want strings but got nil")
	}
	if err != nil {
		t.Errorf("Want nil but got errors")
	}
}

// TestDecrypt tests
func TestDecrypt(t *testing.T) {
	password := []byte("password")
	data := []byte("Data")
	ciphertext, _ := Encrypt(password, data)
	plaintext, err := Decrypt(password, ciphertext)
	if plaintext == nil {
		t.Errorf("Want strings but got nil")
	}
	if err != nil {
		t.Errorf("Want nil but got errors")
	}
}

// TestVerifyKey verify the password
func TestVerifyKey(t *testing.T) {
	password := []byte("password")
	fakepassword := []byte("fakepassword")
	data := []byte("Data")
	ciphertext, _ := Encrypt(password, data)
	res := VerifyKey(password, ciphertext)
	if res != true {
		t.Errorf("Want true but got false")
	}
	res = VerifyKey(fakepassword, ciphertext)
	if res != false {
		t.Errorf("Want false but got true")
	}
}

// TestCreateDatabase create folder and the file to save data
func TestCreateDatabase(t *testing.T) {
	// err := CreateDatabase()
	// if err != nil {
	// 	t.Errorf("Want nil but got error")
	// }
}

// TestSavePasswords save all password
func TestSavePasswords(t *testing.T) {

}

// TestReadPasswords save all password
func TestReadPasswords(t *testing.T) {

}

// RemoveAccount removes an acccount
func TestRemoveAccount(t *testing.T) {
	var data = make([]Account, 2)
	data = RemoveAccount(data, 0)
	if len(data) == 0 || len(data) != 1 {
		t.Errorf("Want %T but got ", data)
		t.Error(data)
	}
	data = make([]Account, 1)
	data = RemoveAccount(data, 0)
	if len(data) != 0 {
		t.Errorf("Want %T but got ", data)
		t.Error(data)
	}
}

// TestConfirmPrompt will prompt to user for yes or no
func TestConfirmPrompt(t *testing.T) {

}
