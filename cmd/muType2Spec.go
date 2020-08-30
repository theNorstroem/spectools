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
)

// muType2SpecCmd represents the muType2Spec command
var muType2SpecCmd = &cobra.Command{
	Use:   "muType2Spec",
	Short: "Updates the type specs with the definitions from the type µSpecs.",
	Long: `The converter will update your type specs and also delete specs and fields if they are not in the µSpec file anymore.

Do not forget to set your µSpec folder in the .spectools config.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("muType2Spec called")
	},
}

// needed for the documentation generator
var MuType2SpecCmd = muType2SpecCmd

func init() {
	rootCmd.AddCommand(muType2SpecCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// muType2SpecCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// muType2SpecCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
