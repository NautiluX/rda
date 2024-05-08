package cmd

import (
	"fmt"

	"github.com/NautiluX/rda/pkg/project"
	"github.com/spf13/cobra"
)

const (
	projectIdArg string = "id"
	addEpicArg   string = "epic"
)

// projectCmd represents the project command
var promoteProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectId, err := cmd.Flags().GetString(projectIdArg)
		if err != nil {
			panic(fmt.Errorf("failed to get argument '%s': %w", projectIdArg, err))
		}
		epic, err := cmd.Flags().GetString(addEpicArg)
		if err != nil {
			panic(fmt.Errorf("failed to get argument '%s': %w", projectIdArg, err))
		}
		projectToPromote, err := project.GetProjectById(projectId)
		if err != nil {
			panic(err)
		}
		projectToPromote.Promote(epic)
		err = project.WriteProjectYaml(*projectToPromote)
		if err != nil {
			panic(err)
		}
		err = project.RenderRegistry()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	promoteCmd.AddCommand(promoteProjectCmd)
	promoteProjectCmd.PersistentFlags().StringP(projectIdArg, "i", "", "Project ID (e.g. RDA0001)")
	promoteProjectCmd.PersistentFlags().StringP(addEpicArg, "e", "", "Epic link to add")
}
