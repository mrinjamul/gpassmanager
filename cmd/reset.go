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
	"strings"

	"github.com/fatih/color"
	"github.com/mrinjamul/gpassmanager/gpm"
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Erase all passwords including master key",
	Long: `It's like a hard reset.
If you forget your master key then you have to perform a hard reset`,
	Run: resetRun,
}

func resetRun(cmd *cobra.Command, args []string) {
	colorFmt := color.New(color.FgRed, color.Bold)
	var response string
	colorFmt.Print("Do you want to erase all passwords (y/n) : ")
	fmt.Scanln(&response)
	switch strings.ToLower(response) {
	case "y", "yes":
		err := ioutil.WriteFile(gpm.DatabaseFile, nil, 0644)
		if err != nil {
			colorFmt.Printf("%v", err)
		}
		colorFmt.Println("All passwords has been cleared.")
	case "n", "no":
		colorFmt = color.New(color.FgGreen)
		colorFmt.Println("Operation Canceled")
	default:
		colorFmt = color.New(color.FgGreen)
		colorFmt.Println("Operation Canceled")
	}
}

func init() {
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
