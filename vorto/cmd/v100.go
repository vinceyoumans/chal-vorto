/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	v100 "github.com/vinceyoumans/chal-vorto/vorto/pkg/v100"

	"github.com/spf13/cobra"

	util "github.com/vinceyoumans/chal-vorto/vorto/pkg/util"
)

// v100Cmd represents the v100 command
var v100Cmd = &cobra.Command{
	Use:   "v100",
	Short: "vorto challenge",
	Long: `Based on the vorto challenge, given a testfile 
of a list of Loads, with Pickup and dropOff lat longs,
this app will output a list of drivers and their assigned load.
The result will include the minimum drivers count so that they do not
Violate 12 hour rule.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("v100 called")
		problemPath, err := cmd.Flags().GetString("ProblemPath")
		if err != nil {
			fmt.Println(err)
			return
		}
		// slogPkg.MakeOutputDirs()
		util.MakeOutputDirs()

		result := v100.V100Start(problemPath)

		for _, valResult := range result {
			strNumbers := make([]string, len(valResult))
			for i, num := range valResult {
				strNumbers[i] = strconv.Itoa(num)
			}
			fmt.Printf("[%s]\n", strings.Join(strNumbers, ", "))
		}

	},
}

func init() {
	rootCmd.AddCommand(v100Cmd)

	v100Cmd.Flags().StringP("ProblemPath", "T", "../training/Problems/problem20.txt", "select a Single Problem to work on")
}
