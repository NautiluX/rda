package cmd

import (
	"github.com/NautiluX/rda/pkg/project"
	"github.com/spf13/cobra"
)

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := project.RenderRegistry()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)
}
