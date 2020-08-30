package main

import (
	"bytes"
	"github.com/spf13/cobra/doc"
	"github.com/theNorstroem/spectools/cmd"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var docPath = "./doc"
var commandList = map[string]*cobra.Command{

	cmd.RootCmd.Use + ".md":                                   cmd.RootCmd,
	cmd.RootCmd.Use + "_" + cmd.MuService2SpecCmd.Use + ".md": cmd.MuService2SpecCmd,
	cmd.RootCmd.Use + "_" + cmd.MuType2SpecCmd.Use + ".md":    cmd.MuType2SpecCmd,
	cmd.RootCmd.Use + "_" + cmd.WatchCmd.Use + ".md":          cmd.WatchCmd,
	cmd.RootCmd.Use + "_" + cmd.InitCmd.Use + ".md":           cmd.InitCmd,
	cmd.RootCmd.Use + "_" + cmd.RunCmd.Use + ".md":            cmd.RunCmd,
}

func main() {
	for filename, command := range commandList {
		genMD(filename, command)
	}
}

func genMD(filename string, cmd *cobra.Command) {
	out := new(bytes.Buffer)
	err := doc.GenMarkdown(cmd, out)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(docPath+"/"+filename, out.Bytes(), 0644)

	if err != nil {
		log.Fatal(err)
	}

}
