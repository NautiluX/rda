package cmd

import (
	"fmt"
	"os"

	"github.com/NautiluX/rda/pkg/project"
	"github.com/spf13/cobra"
)

type AddProjectArgument string

const (
	nameArg        AddProjectArgument = "name"
	descriptionArg AddProjectArgument = "description"
	authorArg      AddProjectArgument = "author"
	sponsorArg     AddProjectArgument = "sponsor"
	typeArg        AddProjectArgument = "type"
	referenceArg   AddProjectArgument = "reference"
)

// projectCmd represents the project command
var addProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := mandatoryString(cmd, nameArg)
		description := mandatoryString(cmd, descriptionArg)
		author := mandatoryString(cmd, authorArg)
		sponsor := optionalString(cmd, sponsorArg)
		projectType := mandatoryString(cmd, typeArg)
		reference := mandatoryString(cmd, referenceArg)
		newProject := project.NewProject(name, description, author, sponsor, reference, project.ProjectType(projectType))
		err := project.AddToRegistry(newProject)
		if err != nil {
			panic(err)
		}
		err = project.RenderRegistry()
		if err != nil {
			panic(err)
		}
	},
}

func mandatoryString(cmd *cobra.Command, arg AddProjectArgument) string {
	argument, err := cmd.Flags().GetString(string(arg))
	if err != nil {
		panic(fmt.Errorf("failed to get argument '%s': %w", arg, err))
	}
	if argument == "" {
		fmt.Printf("missing mandatory argument '%s'.\n", arg)
		os.Exit(1)
	}
	return argument
}

func optionalString(cmd *cobra.Command, arg AddProjectArgument) string {
	argument, err := cmd.Flags().GetString(string(arg))
	if err != nil {
		panic(fmt.Errorf("failed to get argument '%s': %w", arg, err))
	}
	return argument
}

func init() {
	addCmd.AddCommand(addProjectCmd)
	addProjectCmd.PersistentFlags().StringP(string(nameArg), "n", "", "Project name")
	addProjectCmd.PersistentFlags().StringP(string(descriptionArg), "d", "", "Project description")
	addProjectCmd.PersistentFlags().StringP(string(authorArg), "a", "", "Primary Author (owner)")
	addProjectCmd.PersistentFlags().StringP(string(sponsorArg), "s", "", "Sponsor (optional)")
	addProjectCmd.PersistentFlags().StringP(string(typeArg), "t", "", "Project Type (process, utility, other)")
	addProjectCmd.PersistentFlags().StringP(string(referenceArg), "r", "", "Project Reference (e.g. doc, repository)")
}
