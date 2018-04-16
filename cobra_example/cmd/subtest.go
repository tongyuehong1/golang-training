// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var a bool
var aaa bool

// subtestCmd represents the subtest command
var subtestCmd = &cobra.Command{
	Use:   "subtest",
	Short: "A subcommand example",
	Long: "Subtest is a test application for subcommand",
	Run: func(cmd *cobra.Command, args []string) {
		if aaa {
			fmt.Println(123456789)
		}

		if a {
			fmt.Println(1111111)
		}
		fmt.Println("subtest called")
	},
}

func init() {
	RootCmd.AddCommand(subtestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subtestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	 subtestCmd.Flags().BoolVarP(&aaa, "ta", "t", false, "An example for using flags")
	 subtestCmd.Flags().BoolVar(&a, "s", true, "short" )
}
