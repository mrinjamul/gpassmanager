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
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"github.com/mrinjamul/gpassmanager/gpm"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import password(s) from a file",
	Long:  `gpassmanager import "[file location]"`,
	Run:   importRun,
}

func importRun(cmd *cobra.Command, args []string) {
	// check if database is existss
	if _, err := os.Stat(gpm.DatabaseFile); os.IsNotExist(err) {
		gpm.CreateDatabase()
	}

	// Rich import (csv file)
	if csvOpt {
		// required variables
		var accounts []gpm.Account
		// requirements check
		if len(args) == 0 {
			color.Red("Error: too short argument")
			os.Exit(0)
		}
		if len(args) > 1 {
			color.Red("Error: too much arguments")
			os.Exit(0)
		}
		// simplify filename
		filename := gpm.GetFileName(args[0], ".csv")

		// check if import file is exists
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			color.Red("Error: csv file doesn't exists !")
			os.Exit(1)
		}
		// get RAW data for checking file status
		data, err := ioutil.ReadFile(gpm.DatabaseFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// new user message
		if len(data) == 0 {
			color.Red("Warning: If you forget your master password your data will be lost !!")
			color.Yellow("Master password can contains characters and symbols.")
			fmt.Println()
		}
		// get password
		fmt.Print("password: ")
		bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		// new user password confirmation
		if len(data) == 0 {
			fmt.Print("Verify password: ")
			byteVerifyPassword, _ := terminal.ReadPassword(int(syscall.Stdin))
			fmt.Println()
			if string(bytePassword) != string(byteVerifyPassword) {
				color.Red("Error: both password isn't same !")
				os.Exit(0)
			}
		}
		// check if password if not null
		if string(bytePassword) == "" {
			color.Red("Error: you haven't entered password")
			if len(data) == 0 {
				color.Red("Master password can't be empty")
			}
			os.Exit(0)
		}
		// master key can't have numbers
		/*
			(there is a problem in encryption algo)
			need to figure out fix !
		*/
		if len(data) == 0 && bytes.ContainsAny(bytePassword, "0123456789") {
			color.Red("Error: master key can't have numbers !!")
			color.Yellow("Tips: Use passphrases instead")
			os.Exit(0)
		}

		if len(data) == 0 && len(string(bytePassword)) < 6 {
			color.Red("Master password must be greater than 5")
			os.Exit(0)
		}
		//  user message
		if len(data) == 0 {
			color.Green("New User created successfully !")
		}
		// verify password
		if len(data) != 0 && gpm.VerifyKey(bytePassword, data) == false {
			color.Red("Error: Wrong password !")
			os.Exit(1)
		}
		// read decrypt data by the pass key
		if len(data) != 0 && gpm.VerifyKey(bytePassword, data) {
			account, err := gpm.ReadPasswords(bytePassword)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			accounts = account
		}
		// read csv file
		csvpasswords, err := gpm.ReadCSV(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// convert from csv to accounts
		newAccounts := gpm.ConvertToAccount(csvpasswords)
		// append new data into old ones
		for _, account := range newAccounts {
			accounts = append(accounts, account)
		}
		// prompt for save changes
		message := strconv.Itoa(len(newAccounts)) + " password(s) added !\n"
		response := gpm.ConfirmPrompt(message + "Do want to save changes?")
		if response {
			gpm.SavePasswords(bytePassword, accounts)
			color.Green(filename + " imported successfully !")
		} else {
			color.Yellow("Operation Canceled")
			os.Exit(0)
		}

	} else { // normal import (gpm file)

		if len(args) == 0 {
			color.Red("Error: too short argument")
			os.Exit(0)
		}
		if len(args) > 1 {
			color.Red("Error: too much arguments")
			os.Exit(0)
		}
		// simplify filename
		filename := gpm.GetFileName(args[0], ".gpm")
		// check if data file exists
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			color.Red("Error: data file doesn't exists !")
			os.Exit(1)
		}
		data, err := ioutil.ReadFile(gpm.DatabaseFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Overwrite prompts
		if len(data) != 0 {
			fmt.Println("data already exists !")
			var response string
			fmt.Print("Do you want to overwrite existing data (yes/no) :")
			fmt.Scanln(&response)
			switch strings.ToLower(response) {
			case "y", "yes":
				color.Red("Warning: data will be overwritten !")
			case "n", "no":
				color.Green("Operation Canceled")
				os.Exit(0)
			default:
				color.Green("Operation Canceled")
				os.Exit(0)
			}
		}

		backup, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if len(backup) < 3 {
			color.Red("Error: invalid backup file !")
			os.Exit(0)
		}

		err = ioutil.WriteFile(gpm.DatabaseFile, backup, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		color.Green(filename + " data imported !")
	}
}

var (
	csvOpt bool
)

func init() {
	rootCmd.AddCommand(importCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	importCmd.Flags().BoolVarP(&csvOpt, "csv", "c", false, "Import CSV file into the password manager (Currently Google password csv file is supported)")
}
