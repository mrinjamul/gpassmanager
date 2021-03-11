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
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"

	"github.com/fatih/color"
	"github.com/mrinjamul/gpassmanager/gpm"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "view a particular password or entire passwords",
	Long: `view a particular password or entire passwords
Example: gpassmanager view
then gpassmanager view 1`,
	Run: viewRun,
}

func viewRun(cmd *cobra.Command, args []string) {
	// Check if database exists or create
	if _, err := os.Stat(gpm.DatabaseFile); os.IsNotExist(err) {
		gpm.CreateDatabase()
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
	// Show as per instructions
	if viewAll {
		printableData := "\n"
		// View all passwords
		for id, account := range accounts {
			printableData += gpm.LineBreak()
			printableData += "[" + strconv.Itoa(id+1) + "]" + "\t" + "Account: " + account.AccountName + "\n"
			printableData += "Username: " + account.UserName + "\n"
			printableData += "Password: " + account.Password + "\n"
			if account.Email != "" {
				printableData += "Email: " + account.Email + "\n"
			}
			if account.Phone != "" {
				printableData += "Mobile no: " + account.Phone + "\n"
			}
			if account.Notes != "" {
				printableData += "Notes: " + account.Notes + "\n"
			}
			printableData += gpm.LineBreak()
		}
		printableData += "\n"
		err := gpm.ToPager(printableData)
		if err != nil {
			gpm.PagerErrorLogger(err)
		}
	} else if len(args) == 0 { // print only lists with index
		printableData := "\n"
		for id, account := range accounts {
			printableData += gpm.LineBreak()
			printableData += "[" + strconv.Itoa(id+1) + "]" + "\t" + "Account: " + account.AccountName + "\n"
		}
		printableData += gpm.LineBreak() + "\n"
		err := gpm.ToPager(printableData)
		if err != nil {
			gpm.PagerErrorLogger(err)

		}
	} else {
		viewList := []int{}
		for id := range args {
			i, err := strconv.Atoi(args[id])
			if err != nil || i == 0 {
				color.Red(args[id] + " is not a valid id\ninvalid syntax")
				os.Exit(0)
			}
			if i > 0 {
				viewList = append(viewList, i-1)
			}
		}
		viewList = gpm.RemoveDuplicate(viewList)
		// fmt.Println(viewList)
		printableData := "\n"
		for _, id := range viewList {
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
		}
		printableData += "\n"
		err := gpm.ToPager(printableData)
		if err != nil {
			gpm.PagerErrorLogger(err)
		}
	}
}

var (
	viewAll bool
)

func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	viewCmd.Flags().BoolVarP(&viewAll, "all", "a", false, "view all passwords in the store")
}
