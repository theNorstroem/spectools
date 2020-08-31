package installer

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Installing dependencies")
	deps := viper.GetStringSlice("dependencies")
	fmt.Println(deps)
	for _, d := range deps {
		strings.Split(d, " ")
		// check it ends with .git

		// create path if it does not exist

		// clone if it is new

		// git pull

		// checkout version
	}
}
