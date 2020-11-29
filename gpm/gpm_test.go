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

// TestGetVersion tests
func TestGetLincense(t *testing.T) {
	out := GetLicense()
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
	// TODO
}

// TestSavePasswords save all password
func TestSavePasswords(t *testing.T) {
	// TODO
}

// TestReadPasswords save all password
func TestReadPasswords(t *testing.T) {
	// TODO
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
	// TODO
}

// TestGeneratePassword will return Password
func TestGeneratePassword(t *testing.T) {
	pass := GeneratePassword(0)
	if pass == "" {
		t.Errorf("Want string but got nil")
	}
	if len(pass) != 12 {
		t.Errorf("Want 12 length of string but got %d", len(pass))
	}

	pass = GeneratePassword(16)
	if pass == "" {
		t.Errorf("Want string but got nil")
	}
	if len(pass) != 16 {
		t.Errorf("Want 12 length of string but got %d", len(pass))
	}

	for i := 1; i < 8; i++ {
		pass = GeneratePassword(i)
		if pass == "" {
			t.Errorf("Want string but got nil")
		}
		if len(pass) != 8 {
			t.Errorf("Want 8 length of string but got %d", len(pass))
		}
	}
}

// TestSortSlice sort arrays
func TestSortSlice(t *testing.T) {
	slice := []int{4, 2, 8, 3, 8, 5, 2, 1}
	slice = SortSlice(slice)
	want := []int{8, 8, 5, 4, 3, 2, 2, 1}
	if len(slice) != len(want) {
		t.Errorf("Both result is not same")
	}
	for i := range slice {
		if slice[i] != want[i] {
			t.Errorf("Want %d but got %d", want[i], slice[i])
		}
	}
}

// TestRemoveDuplicate removes duplicate from slice
func TestRemoveDuplicate(t *testing.T) {
	slice := []int{4, 2, 8, 3, 8, 5, 2, 1}
	slice = SortSlice(slice)
	slice = RemoveDuplicate(slice)
	want := []int{8, 5, 4, 3, 2, 1}
	if len(slice) != len(want) {
		t.Errorf("Both result is not same")
	}
	for i := range want {
		if slice[i] != want[i] {
			t.Errorf("Want %d but got %d", want[i], slice[i])
		}
	}
}

// TestTGetFileName simplify filenames for use (Note: only 3 char ext)
func TestGetFileName(t *testing.T) {
	testcases := []struct {
		name     string
		filename string
		ext      string
		result   string
	}{
		{"smaller file name (gpm)", "abc", ".gpm", "abc.gpm"},
		{"smaller file name with extensions (gpm)", "abc.gpm", ".gpm", "abc.gpm"},
		{"normal file name (gpm)", "backup", ".gpm", "backup.gpm"},
		{"normal file name with extensions (gpm)", "backup.gpm", ".gpm", "backup.gpm"},

		{"smaller file name (csv)", "abc", ".csv", "abc.csv"},
		{"smaller file name with extensions (csv)", "abc.csv", ".csv", "abc.csv"},
		{"normal file name (csv)", "backup", ".csv", "backup.csv"},
		{"normal file name with extensions (csv)", "backup.csv", ".csv", "backup.csv"},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			filename := GetFileName(testcase.filename, testcase.ext)
			if filename != testcase.result {
				t.Errorf("Wants to be %v; but got %v", testcase.result, filename)
			}
		})
	}
}

//TestReadCSV returns parsed Passwords (*Google Password csv file)
func TestReadCSV(t *testing.T) {
	// TODO
}

// TestConvertToAccount converts CSVPassword into Account
func TestConvertToAccount(t *testing.T) {
	var csvPass []CSVPassword = make([]CSVPassword, 10)
	testAccounts := ConvertToAccount(csvPass)
	if len(testAccounts) != 10 {
		t.Errorf("Accounts should be same number of password; but go %v", len(testAccounts))
	}
}

// TestCopy test for copy
func TestCopy(t *testing.T) {
	// TODO
}
