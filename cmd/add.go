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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/mrinjamul/gpassmanager/gpm"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new password",
	Long:  ``,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(gpm.DatabaseFile); os.IsNotExist(err) {
		gpm.CreateDatabase()
	}
	var account gpm.Account
	var password string
	var accounts []gpm.Account

	data, err := ioutil.ReadFile(gpm.DatabaseFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(data) == 0 {
		color.Red("Warning: If you forget your master password your data will be lost !!")
		fmt.Println()
	}
	fmt.Print("password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	password = string(bytePassword)
	if password == "" {
		color.Red("Error: you haven't entered password")
		if len(data) == 0 {
			color.Red("Master password can't be empty")
		}
		os.Exit(0)
	}
	if len(data) == 0 && len(password) < 6 {
		color.Red("Master password must be greater than 5")
		os.Exit(0)
	}
	if len(data) != 0 && gpm.VerifyKey(bytePassword, data) == false {
		color.Red("Error: Wrong password !")
		os.Exit(1)
	}
	if len(data) != 0 && gpm.VerifyKey(bytePassword, data) {
		account, err := gpm.ReadPasswords(bytePassword)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		accounts = account
	}
	fmt.Print("Enter Account Name: ")
	fmt.Scanln(&account.AccountName)
	fmt.Print("Enter username: ")
	fmt.Scanln(&account.UserName)
	if account.UserName == "" {
		fmt.Println("username can't be empty!")
		os.Exit(0)
	}
	fmt.Print("Enter password: ")
	fmt.Scanln(&account.Password)
	if account.Password == "" {
		fmt.Println("password can't be empty!")
		os.Exit(0)
	}
	fmt.Print("Enter email: ")
	fmt.Scanln(&account.Email)
	fmt.Print("Enter mobile no: ")
	fmt.Scanln(&account.Phone)
	fmt.Print("Notes :")
	fmt.Scanln(&account.Notes)
	fmt.Println()
	gpm.LineBreak()
	prompt := promptui.Select{
		Label: "Do you want to save (Yes/No)",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if result == "Yes" {

		accounts = append(accounts, account)
		gpm.SavePasswords(bytePassword, accounts)
		color.Green("Password Saved!")

	} else {
		fmt.Println("Password Not Saved!")
	}

}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
