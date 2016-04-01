// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"github.com/KaseyPowers/packer/utils"
	"github.com/spf13/cobra"
)

// tagsCmd represents the tags command
var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Print out this projects tags",
	Long: `Use this to check if your tag is in the repo`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		min, err := cmd.Flags().GetString("min")
	  if err != nil {
	    utils.Error(err)
	  }
		max, err := cmd.Flags().GetString("max")
	  if err != nil {
	    utils.Error(err)
	  }


		for _ , val := range utils.GitTags(min, max) {
			fmt.Println(val)
		}
	},
}

func init() {
	RootCmd.AddCommand(tagsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tagsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tagsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	tagsCmd.Flags().StringP("min", "m", "", "Only display tags greater than or equal to this value")
	tagsCmd.Flags().String("max", "", "Only display tags less than than or equal to this value")

}
