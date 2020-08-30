package runner

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Run(cmd *cobra.Command, args []string) {
	flow := "default"
	f := cmd.Flag("flow")
	if f != nil {
		flow = f.Value.String()
	}

	commandList := collectCommands(cmd.Parent())

	fmt.Println("running flow " + flow)
	seq := viper.GetStringSlice("flows." + flow)
	fmt.Println("starting sequence", seq)
	for _, step := range seq {
		commandList[step](cmd, args)
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
