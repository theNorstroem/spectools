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

func main() {
	// gen root command
	genMD(cmd.RootCmd.Use+".md", cmd.RootCmd)
	// generate the other commands
	for _, c := range cmd.RootCmd.Commands() {
		genMD(cmd.RootCmd.Use+"_"+c.Use+".md", c)
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
