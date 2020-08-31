package util

import (
	"github.com/spf13/viper"
	"os"
	"path"
	"strings"
)

func GetDependencyList() []string {
	deps := []string{}
	for _, d := range viper.GetStringSlice("dependencies") {
		p := path.Join(viper.GetString("dependenciesDir"), strings.Split(d, " ")[0])
		if strings.HasSuffix(p, ".git") {

		}
		// load config to resolve Message and Service dirs to append

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
			deps = append(deps, p)
		}

	}
	return deps
}
