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
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"github.com/mrinjamul/gpassmanager/gpm"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "remove an Account from password manager",
	Long:    ``,
	Run:     removeRun,
}

func removeRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		color.Red("Error: too short argument")
		color.Yellow("Usaage: gpassmanager remove [id]")
		os.Exit(0)
	}
	if len(args) > 1 {
		color.Red("Error: too much arguments")
		color.Yellow("Usaage: gpassmanager remove [id]")
		os.Exit(0)
	}
	i, err := strconv.Atoi(args[0])

	if err != nil {
		color.Red(args[0] + " is not a valid id\ninvalid syntax")
		os.Exit(0)
	}
	if _, err := os.Stat(gpm.DatabaseFile); os.IsNotExist(err) {
		gpm.CreateDatabase()
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
	accounts, err := gpm.ReadPasswords(bytePassword)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	if len(accounts) == 0 {
		fmt.Println("No passwords found !")
		os.Exit(0)
	}
	if i > 0 && i <= len(accounts) {
		colorFmt := color.New(color.FgRed, color.Bold)
		var response string
		var accountName string = ""
		username := accounts[i-1].UserName
		if accounts[i-1].AccountName != "" {
			accountName = " (" + accounts[i-1].AccountName + ")"
		}
		colorFmt.Print("Do you want to remove " + username + accountName + " (y/n) : ")
		fmt.Scanln(&response)
		switch strings.ToLower(response) {
		case "y", "yes":
			accounts = gpm.RemoveAccount(accounts, i-1)
			color.Yellow("[" + strconv.Itoa(i) + "] " + username + " has been removed")
			if len(accounts) != 0 {
				gpm.SavePasswords(bytePassword, accounts)
			} else {
				res := gpm.ConfirmPrompt("All password removed!\nDo you want to remove the master key ?")
				if res {
					err := gpm.CreateDatabase()
					if err != nil {
						fmt.Println(err)
					}
				} else {
					gpm.SavePasswords(bytePassword, accounts)
				}

			}

		case "n", "no":
			colorFmt = color.New(color.FgGreen)
			colorFmt.Println("Operation Canceled")
		default:
			colorFmt = color.New(color.FgGreen)
			colorFmt.Println("Operation Canceled")
		}
	} else {
		color.Red("[" + strconv.Itoa(i) + "] does not match any Account !")
	}

}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
