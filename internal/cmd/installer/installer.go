package installer

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/otiai10/copy"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theNorstroem/spectools/pkg/util"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func Run(cmd *cobra.Command, args []string) {

	freshInstall := false
	f := cmd.Flag("fresh")
	if f != nil {
		freshInstall = f.Value.String() == "true"
		fmt.Println("Fresh install requested")
	}

	fmt.Println("Installing dependencies")
	deps := viper.GetStringSlice("dependencies")
	fmt.Println(deps)
	for _, d := range deps {
		strings.Split(d, " ")

		dep := util.ParseDependency(d)
		spectoolRepository, err := os.UserHomeDir()
		if err != nil {
			// use tmpdir if ho home
			spectoolRepository = os.TempDir()
		}
		spectoolRepository = spectoolRepository + "/.spectools"
		packageRepoDir := path.Join(spectoolRepository, dep.Path)
		if !util.DirExists(dep.DependencyPath) {
			mkdirRecursive(dep.DependencyPath)
		}
		if dep.Kind == util.GIT {
			// removie repodir if freshInstall is requested
			if freshInstall {
				os.RemoveAll(packageRepoDir)
			}
			// create path if it does not exist
			if !util.DirExists(packageRepoDir) {
				// create
				mkdirRecursive(packageRepoDir)
				// clone if it is new
				_, err := git.PlainClone(packageRepoDir, false, &git.CloneOptions{
					URL:      dep.Repository,
					Depth:    1,
					Progress: os.Stdout,
				})

				if err != nil {
					// use exec
					log.Println(err)
					log.Println("switching to git executable")
					e := CloneWithGitCommand(packageRepoDir, dep.Repository)
					if e != nil {
						log.Fatal(err)
					}
				}
			}

			// fetch the changes
			r, err := git.PlainOpen(packageRepoDir)
			if err != nil {
				log.Fatal(err)
			}
			err = r.Fetch(&git.FetchOptions{Tags: git.AllTags})
			if err != nil && err != git.NoErrAlreadyUpToDate {
				fmt.Println(dep.Repository, err.Error())
				log.Println("switching to git executable")
				e := FetchWithGitCommand(packageRepoDir)
				if e != nil {
					log.Fatal(err)
				}
			}

			// Get the working directory for the repository
			w, err := r.Worktree()
			if err != nil {
				log.Fatal(err)
			}
			// checkout version
			err = w.Checkout(&git.CheckoutOptions{Branch: plumbing.NewTagReferenceName(dep.Version)})
			if err != nil {
				// try branch
				err = w.Checkout(&git.CheckoutOptions{Branch: plumbing.NewBranchReferenceName(dep.Version)})
				if err != nil {
					log.Fatal(err)
				}
			}

			// clear dep path
			err = os.RemoveAll(dep.DependencyPath)
			if err != nil {
				fmt.Println(err)
			}
			// copy from packageRepoDir to dep.DependencyPath
			copy.Copy(packageRepoDir, dep.DependencyPath, copy.Options{
				OnSymlink: nil,
				Skip: func(src string) (bool, error) {
					return path.Base(src) == ".git", nil
				},
				AddPermission: 0,
				Sync:          false,
			})

		} else {
			// todo discuss to implement file system dep ???
			fmt.Println("File system deps are not implemented yet")
		}

	}
}

func FetchWithGitCommand(packageRepoDir string) error {
	fmt.Println("git fetch --depth=-1")
	cmd := exec.Command("git", "fetch", "--depth=1")
	cmd.Dir = packageRepoDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CloneWithGitCommand(packageRepoDir string, repository string) error {
	fmt.Println("git clone --depth=-1", repository, ".")
	cmd := exec.Command("git", "clone", "--depth=1", repository, ".")
	cmd.Dir = packageRepoDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func mkdirRecursive(subdir string) {

	pathSegments := strings.Split(subdir, "/")
	p := "./"
	if pathSegments[0] == "" {
		p = "/"
	}
	for _, folder := range pathSegments {
		newDir := path.Clean(p + folder)
		if !util.DirExists(newDir) {
			os.Mkdir(newDir, 0755)
		}

		p = p + folder + "/"
	}
}
