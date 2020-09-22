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
	"strings"

	"github.com/fatih/color"
	"github.com/mrinjamul/gpassmanager/gpm"
	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import data from a file",
	Long: `	 gpassmanager import "[file location]"`,
	Run: importRun,
}

func importRun(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(gpm.DatabaseFile); os.IsNotExist(err) {
		gpm.CreateDatabase()
	}

	if len(args) == 0 {
		color.Red("Error: too short argument")
		os.Exit(0)
	}
	if len(args) > 1 {
		color.Red("Error: too much arguments")
		os.Exit(0)
	}
	filename := args[0]
	if len(filename) > 4 {
		if filename[len(filename)-4:] != ".gpm" {
			filename += ".gpm"
		}
	} else {
		filename += ".gpm"
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		color.Red("Error: data file doesn't exists !")
		os.Exit(1)
	}
	data, err := ioutil.ReadFile(gpm.DatabaseFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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

func init() {
	rootCmd.AddCommand(importCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
