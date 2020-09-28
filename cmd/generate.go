// Package cmd ...
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
package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"

	"github.com/fatih/color"
	"github.com/mrinjamul/gpassmanager/gpm"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generate secure password",
	Long:    ``,
	Run:     generateRun,
}

func generateRun(cmd *cobra.Command, args []string) {
	var (
		passLength int = 0
		err        error
	)

	if len(args) > 1 {
		color.Red("Error: too much arguments")
		color.Yellow("Usage: gpassmanager generate [password length]")
		os.Exit(0)
	}

	if len(args) != 0 {
		passLength, err = strconv.Atoi(args[0])
		if err != nil {
			color.Red(args[0] + " is not a valid password length\ninvalid syntax")
			os.Exit(0)
		}
		if passLength == 0 {
			color.Red("Error: password length can't be 0")
			os.Exit(0)
		}
		if passLength < 8 {
			color.Yellow("Password length should be greater than 8")
			color.Yellow("Warning: using default password length")
		}
	}
	// } else {
	// 	var length string
	// 	fmt.Print("Enter password length: ")
	// 	fmt.Scanln(&length)
	// 	if length != "" {
	// 		passLength, err = strconv.Atoi(length)
	// 		if err != nil {
	// 			color.Red(length + " is not a valid password length\ninvalid syntax")
	// 			os.Exit(0)
	// 		}
	// 	}
	// 	if passLength < 8 {
	// 		color.Yellow("Password length should be greater than 8")
	// 		color.Yellow("Warning: using default password length")
	// 	}
	// }

	pass := gpm.GeneratePassword(passLength)
	fmt.Println("Password: " + pass + "")

	if _, err := os.Stat(gpm.DatabaseFile); os.IsNotExist(err) {
		gpm.CreateDatabase()
	}
	data, err := ioutil.ReadFile(gpm.DatabaseFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(data) == 0 {
		os.Exit(0)
	}
	response := gpm.ConfirmPrompt("Do you want to save to the password store?")
	if response == true {
		fmt.Println("Yes")
		fmt.Print("password: ")
		bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if gpm.VerifyKey(bytePassword, data) == false {
			color.Red("Error: Wrong password !")
			os.Exit(1)
		}
		var account gpm.Account
		var accounts []gpm.Account

		if len(data) != 0 && gpm.VerifyKey(bytePassword, data) {
			accounts, err = gpm.ReadPasswords(bytePassword)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		}

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Enter Account Name: ")
		if scanner.Scan() {
			account.AccountName = scanner.Text()
		}
		fmt.Print("Enter username: ")
		if scanner.Scan() {
			account.UserName = scanner.Text()
		}
		if account.UserName == "" {
			fmt.Println("username can't be empty!")
			os.Exit(0)
		}
		account.Password = pass
		fmt.Print("Enter email: ")
		if scanner.Scan() {
			account.Email = scanner.Text()
		}
		fmt.Print("Enter mobile no: ")
		if scanner.Scan() {
			account.Phone = scanner.Text()
		}
		fmt.Print("Notes :")
		if scanner.Scan() {
			account.Notes = scanner.Text()
		}
		fmt.Println()

		accounts = append(accounts, account)
		gpm.SavePasswords(bytePassword, accounts)
		color.Green("Password Saved!")

	} else {
		fmt.Println("No")
	}

}

var (
	passLength int
)

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// generateCmd.Flags().IntVarP(&passLength, "length", "l", 12, "Set password length")
}
