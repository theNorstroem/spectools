package util

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"strings"
)

func GetDependencyList() []string {
	deps := []string{}
	for _, d := range viper.GetStringSlice("dependencies") {
		// the version info is not needed
		parts := strings.Split(d, " ")
		// should only have 2 parts (repo, version)
		if len(parts) > 2 {
			fmt.Println(ScanForStringPosition(d, "./.spectools"), "Error")
			log.Fatal("config error or dependency not installed. Maybe you should run spectools install")
		}
		p := path.Join(viper.GetString("dependenciesDir"), parts[0])

		// remove .git from path, this looks nicer
		if strings.HasSuffix(p, ".git") {
			p = p[0 : len(p)-4]
		}

		// todo check for existence of p and give spectools install hint

		// load config to resolve Message and Service dirs
		depconf := viper.New()
		depconf.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
		depconf.AddConfigPath(p)
		depconf.SetConfigName(".spectools")
		err := depconf.ReadInConfig()
		if err == nil {

			tdir := depconf.GetString("typeSpecDir")
			if tdir != "" {
				if _, err := os.Stat(path.Join(p, tdir)); !os.IsNotExist(err) {
					// path/to/whatever exists
					deps = append(deps, path.Join(p, tdir))
				}

			}
			sdir := depconf.GetString("serviceSpecDir")
			if sdir != "" {
				if _, err := os.Stat(path.Join(p, sdir)); !os.IsNotExist(err) {
					deps = append(deps, path.Join(p, sdir))
				}
			}

		} else {
			// no .spectools config in target dir, use the complete path
			deps = append(deps, p)
		}

	}
	return deps
}
