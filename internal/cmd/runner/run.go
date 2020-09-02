package runner

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Run(cmd *cobra.Command, args []string) {
	flow := "default"

	f := cmd.Flag("flow")
	if f != nil {
		// override with flag
		flow = f.Value.String()
	}

	listOfConfiguredFlows := collectCommands(cmd.Parent())

	fmt.Println("running flow " + flow)
	seq := viper.GetStringSlice("flows." + flow)
	// flow not found
	if len(seq) == 0 {
		log.Fatal("flow has no sequence --flow=", flow)
	}
	fmt.Println("starting sequence", seq)
	for _, step := range seq {
		if listOfConfiguredFlows[step] != nil {
			// configured flows go first
			listOfConfiguredFlows[step](cmd, args)
		} else {
			// try commands
			commandList := viper.GetStringMapString("commands")
			if commandList[strings.ToLower(step)] != "" {
				// exec the command...
				cmd := exec.Command(commandList[strings.ToLower(step)], args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatalf("cmd.Run() failed with %s\n", err)
				}
			}
		}

	}

}

// commands are in parent
func collectCommands(rootCmd *cobra.Command) (commandList map[string]func(cmd *cobra.Command, args []string)) {
	commandList = map[string]func(cmd *cobra.Command, args []string){}
	for _, c := range rootCmd.Commands() {
		commandList[c.Use] = c.Run
	}
	return commandList
}
