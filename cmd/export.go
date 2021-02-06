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
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/mrinjamul/gpassmanager/gpm"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export your data to a file (master key will be also exported)",
	Long: `Usage: gpassmanager export "export filename"
	or
gpassmanager export`,
	Run: exportRun,
}

func exportRun(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(gpm.DatabaseFile); os.IsNotExist(err) {
		color.Red("Error: data not exists !")
		os.Exit(0)
	}
	data, err := ioutil.ReadFile(gpm.DatabaseFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(data) == 0 {
		color.Red("Error: data does not exists !")
		os.Exit(0)
	}
	if len(args) > 1 {
		color.Red("Error: too much arguments")
		os.Exit(0)
	}

	var filename string
	if len(args) == 0 {
		color.Yellow("Generating export with default file name in current directory.\n[Example: gpm-13-01-2000-000.gpm]")
		currentTime := time.Now()
		rand.Seed(time.Now().UTC().UnixNano())
		filename = "gpm-" + currentTime.Format("02-01-2006") + "-" + strconv.Itoa(rand.Intn(100)) + ".gpm"
	} else {
		filename = args[0]
	}

	// simplify filename
	filename = gpm.GetFileName(filename, ".gpm")

	// check if file exists
	if _, err := os.Stat(filename); err == nil {
		fmt.Println(filename, "already exists !")
		var response string
		fmt.Print("Do you want to overwrite existing file (yes/no) :")
		fmt.Scanln(&response)
		switch strings.ToLower(response) {
		case "y", "yes":
			color.Yellow("Warning: " + filename + " will be overwritten !")
		case "n", "no":
			color.Green("Operation Canceled")
			os.Exit(0)
		default:
			color.Green("Operation Canceled")
			os.Exit(0)
		}
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		color.Red("Error: invalid location !")
		os.Exit(1)
	}
	color.Green("data exported to " + filename)

}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
