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
	"syscall"

	"github.com/fatih/color"
	"github.com/mrinjamul/gpassmanager/gpm"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "view all passwords",
	Long:  ``,
	Run:   viewRun,
}

func viewRun(cmd *cobra.Command, args []string) {
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
	for id, account := range accounts {
		gpm.LineBreak()
		fmt.Print("[" + strconv.Itoa(id+1) + "]" + "\t")
		if account.AccountName != "" {
			fmt.Print("Account:", account.AccountName)
		}
		fmt.Println()
		fmt.Println("username:", account.UserName)
		fmt.Println("Password:", account.Password)
		if account.Email != "" {
			fmt.Println("Email:", account.Email)
		}
		if account.Phone != "" {
			fmt.Println("Mobile no:", account.Phone)
		}
		if account.Notes != "" {
			fmt.Println("Notes:", account.Notes)
		}
		gpm.LineBreak()
	}
}

func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
