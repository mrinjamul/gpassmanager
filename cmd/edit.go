/*
Copyright © 2020-2021 Injamul Mohammad Mollah

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
	"log"
	"os"
	"strconv"
	"syscall"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/mrinjamul/gpassmanager/gpm"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"ed", "modify"},
	Short:   "Edit a account details",
	Long:    `Edit a account details`,
	Run:     editRun,
}

func editRun(cmd *cobra.Command, args []string) {
	// Check if database exists
	_, err := os.Stat(gpm.DatabaseFile)
	if os.IsNotExist(err) {
		color.Red("Error: No account exists")
		os.Exit(1)
	}
	// Get raw data for checking
	data, err := ioutil.ReadFile(gpm.DatabaseFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(data) == 0 {
		fmt.Println("No passwords found !")
		os.Exit(0)
	}
	// only take arg as integer
	id, err := strconv.Atoi(args[0])
	if err != nil || id == 0 {
		color.Red(args[id] + " is not a valid id\ninvalid syntax")
		os.Exit(0)
	}
	if id < 0 {
		color.Red(args[id] + " is not a valid id\ninvalid syntax")
		os.Exit(0)
	}

	// as array index
	id = id - 1

	// secure user input
	fmt.Print("password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	// password verifications
	if len(data) != 0 && gpm.VerifyKey(bytePassword, data) == false {
		color.Red("Error: Wrong password !")
		os.Exit(1)
	}
	// decrypt and get All accounts
	accounts, err := gpm.ReadPasswords(bytePassword)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	if len(accounts) == 0 {
		fmt.Println("No passwords found !")
		os.Exit(0)
	}
	// View section
	printableData := "\n"
	printableData += gpm.LineBreak()
	printableData += "[" + strconv.Itoa(id+1) + "]" + "\t" + "Account: " + accounts[id].AccountName + "\n"
	printableData += "Username: " + accounts[id].UserName + "\n"
	printableData += "Password: " + accounts[id].Password + "\n"
	if accounts[id].Email != "" {
		printableData += "Email: " + accounts[id].Email + "\n"
	}
	if accounts[id].Phone != "" {
		printableData += "Mobile no: " + accounts[id].Phone + "\n"
	}
	if accounts[id].Notes != "" {
		printableData += "Notes: " + accounts[id].Notes + "\n"
	}
	printableData += gpm.LineBreak()
	printableData += "\n"
	fmt.Print(printableData)

	// Edit section
	// scanner for taking user input as line
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Account Name: ")
	if scanner.Scan() {
		if scanner.Text() != "" {
			accounts[id].AccountName = scanner.Text()
		}
	}
	fmt.Print("Enter username: ")
	if scanner.Scan() {
		if scanner.Text() != "" {
			accounts[id].UserName = scanner.Text()
		}
	}
	fmt.Print("Enter password: ")
	if scanner.Scan() {
		if scanner.Text() != "" {
			accounts[id].Password = scanner.Text()
		}
	}

	fmt.Print("Enter email: ")
	if scanner.Scan() {
		if scanner.Text() != "" {
			accounts[id].Email = scanner.Text()
		}
	}
	fmt.Print("Enter mobile no: ")
	if scanner.Scan() {
		if scanner.Text() != "" {
			accounts[id].Phone = scanner.Text()
		}
	}
	fmt.Print("Notes :")
	if scanner.Scan() {
		if scanner.Text() != "" {
			accounts[id].Notes = scanner.Text()
		}
	}
	fmt.Println()
	fmt.Println(gpm.LineBreak())
	// prompt for confirmations
	prompt := promptui.Select{
		Label: "Do you want to save (Yes/No)",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if result == "Yes" {
		gpm.SavePasswords(bytePassword, accounts)
		color.Green("Password Saved!")
	} else {
		fmt.Println("Password Not Saved!")
	}
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
