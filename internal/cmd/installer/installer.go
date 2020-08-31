package installer

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
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

		_, err := git.PlainClone("installedSpecs/nn", false, &git.CloneOptions{
			URL:      "git://github.com/theNorstroem/furoBaseSpecs.git",
			Progress: os.Stdout,
		})

		// fetch changes if we are not on the tag
		r, err := git.PlainOpen("installedSpecs/nn")
		r.Fetch(&git.FetchOptions{})

		// Get the working directory for the repository
		w, err := r.Worktree()

		// checkout version
		err = w.Checkout(&git.CheckoutOptions{Hash: plumbing.NewHash("v1.10.0")})

		fmt.Println(err)

	}
}
