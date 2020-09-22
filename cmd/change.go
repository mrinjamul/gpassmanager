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
	"bytes"
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

// changeCmd represents the change command
var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change Master Password",
	Long: `Change Master Password
Example, gpassmanager change --passwd`,
	Run: changeRun,
}

func changeRun(cmd *cobra.Command, args []string) {
	if passwordOpt {
		if _, err := os.Stat(gpm.DatabaseFile); os.IsNotExist(err) {
			fmt.Println("No password found !")
			os.Exit(0)
		}
		data, err := ioutil.ReadFile(gpm.DatabaseFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if len(data) == 0 {
			fmt.Println("No passwords found !")
			os.Exit(0)
		}
		fmt.Print("password: ")
		bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println()

		if len(data) != 0 && gpm.VerifyKey(bytePassword, data) == false {
			color.Red("Error: Wrong password !")
			os.Exit(1)
		}
		fmt.Print("New password: ")
		byteNewPassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if string(bytePassword) == string(byteNewPassword) {
			color.Red("Error: Password already in use !!")
			os.Exit(0)
		}

		if string(byteNewPassword) == "" {
			color.Red("Error: you haven't entered password")
			os.Exit(0)
		}

		if bytes.ContainsAny(byteNewPassword, "0123456789") {
			color.Red("Error: master key can't have numbers !!")
			color.Yellow("Tips: Use passphrases instead")
			os.Exit(0)
		}

		if len(string(byteNewPassword)) < 6 {
			color.Red("Master password must be greater than 5")
			os.Exit(0)
		}

		fmt.Print("Verify password: ")
		byteVerifyPassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println()

		if string(byteNewPassword) != string(byteVerifyPassword) {
			color.Red("Error: both password is not same!")
			os.Exit(0)
		}

		prompt := promptui.Select{
			Label: "Do you want to save changes (Yes/No)",
			Items: []string{"Yes", "No"},
		}
		_, result, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}

		if result == "Yes" {

			accounts, err := gpm.ReadPasswords(bytePassword)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			gpm.SavePasswords(byteVerifyPassword, accounts)

			color.Green("Password changed successfully !!")
		} else {
			fmt.Println("Password remains unchanged !!")
		}
	} else {
		color.Yellow("You haven't selected any options.\nTry Again !!")
		color.Yellow("Usage: gpassmanager change --passwd")
	}
}

var (
	passwordOpt bool
)

func init() {
	rootCmd.AddCommand(changeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// changeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// changeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	changeCmd.Flags().BoolVarP(&passwordOpt, "passwd", "p", false, "change master key for the Data")
}
