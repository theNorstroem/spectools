package util

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"strings"
)

func GetDependencyList() []string {
	deps := []string{}
	for _, d := range viper.GetStringSlice("dependencies") {
		p := path.Join(viper.GetString("dependenciesDir"), strings.Split(d, " ")[0])
		// load config to resolve Message and Service dirs to append

		depconf := viper.New()
		depconf.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
		depconf.AddConfigPath(p)
		depconf.SetConfigName(".spectools")
		err := depconf.ReadInConfig()
		if err == nil {
			tdir := depconf.GetString("typeSpecDir")
			if tdir != "" {
				deps = append(deps, path.Join(p, tdir))
			}
			mdir := depconf.GetString("serviceSpecDir")
			if mdir != "" {
				deps = append(deps, path.Join(p, mdir))
			}
			// notify sub dependencies
			for _, subdep := range depconf.GetStringSlice("dependencies") {
				fmt.Println(p, "requires", subdep, "do not forget to add")
			}

		} else {
			deps = append(deps, p)
		}

	}
	return deps
}
