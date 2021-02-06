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

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search password account in the password store",
	Long:  ``,
	Run:   searchRun,
}

func searchRun(cmd *cobra.Command, args []string) {
	// Check if database exists or create
	if _, err := os.Stat(gpm.DatabaseFile); os.IsNotExist(err) {
		gpm.CreateDatabase()
	}
	// requirements check
	if len(args) == 0 {
		color.Red("Error: too short argument")
		os.Exit(0)
	}
	if len(args) > 1 {
		color.Red("Error: too much arguments")
		os.Exit(0)
	}
	// Get raw data for checking
	data, err := ioutil.ReadFile(gpm.DatabaseFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(data) == 0 {
		color.Yellow("No data found !")
		os.Exit(0)
	}
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
		color.Yellow("No data found !")
		os.Exit(0)
	}
	// Found flag
	isFound := false
	// results counts
	var founds int = 0
	// prints found results
	printableData := "\n"
	for id, account := range accounts {
		if strings.Contains(account.AccountName, args[0]) || strings.Contains(account.UserName, args[0]) || strings.Contains(account.Notes, args[0]) {
			founds++
			isFound = true
			printableData += "[" + strconv.Itoa(id+1) + "]" + "\t" + "Account: " + account.AccountName + "\n"
			printableData += "Notes: " + account.Notes + "\n"
			printableData += gpm.LineBreak()
		}
	}
	if isFound {
		printableData = " " + strconv.Itoa(founds) + " result(s) found !" + "\n" + printableData
		printableData += "\n"
		err := gpm.ToPager(printableData)
		if err != nil {
			gpm.PagerErrorLogger(err)
		}
	} else {
		color.Red("No result found !")
	}
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
