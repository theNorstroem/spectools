/*
Copyright © 2020 Veith Zäch <veithz@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a configured flow.",
	Long: `Runs a configured flow of spectools commands.

To configure a flow of commands just add a "flow" in the flows section of your .spectools config.
A flow is just a list of commands which gets executed in order

Example Config:
[.spectools]
flows:
  type:
    - cleanTypeProtoDir
    - muType2Spec
    - TypeSpec2Proto

Command:

This config will run "cleanTypeProtoDir",  muType2Spec"" and "TypeSpec2Proto" in sequence`,
	Run: func(cmd *cobra.Command, args []string) {
		flow := "default"
		f := cmd.Flag("flow")
		if f != nil {
			flow = f.Value.String()
		}

		fmt.Println("running flow " + flow)
		c := viper.GetStringSlice("flow." + flow)
		fmt.Println(c)
		muType2SpecCmd.Run(cmd, args)
		muService2SpecCmd.Run(cmd, args)

	},
}

// needed for the documentation generator
var RunCmd = runCmd

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	runCmd.Flags().StringP("flow", "f", "all", "A configured flow from the .spectools config")
}
