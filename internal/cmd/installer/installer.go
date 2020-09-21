package installer

import (
	"fmt"
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
				err := CloneWithGitCommand(packageRepoDir, dep.Repository)
				if err != nil {
					log.Fatal(err)
				}
			}

			// fetch the changes
			err := FetchWithGitCommand(packageRepoDir)
			if err != nil {
				log.Fatal(err)
			}

			// checkout requested version
			err = CheckoutWithGitCommand(packageRepoDir, dep.Version)
			if err != nil {
				log.Fatal(err)
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

func CheckoutWithGitCommand(packageRepoDir string, version string) error {
	fmt.Println("git checkout", version)
	cmd := exec.Command("git", "checkout", version)
	cmd.Dir = packageRepoDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
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
