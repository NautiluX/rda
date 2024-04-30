package project

import (
	"fmt"
)

type ProjectType string

const (
	ProjectTypeUtility ProjectType = "utility"
	ProjectTypeProcess ProjectType = "process"
	ProjectTypeOther   ProjectType = "other"
)

type ProjectStage string

const (
	ProjectStageSandbox   ProjectStage = "sandbox"
	ProjectStageProposed  ProjectStage = "proposed"
	ProjectStageGraduated ProjectStage = "graduated"
)

type Project struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	Author      string       `yaml:"author"`
	Sponsor     string       `yaml:"sponsor"`
	Reference   string       `yaml:"reference"`
	ID          string       `yaml:"id"`
	Stage       ProjectStage `yaml:"stage"`
	Type        ProjectType  `yaml:"type"`
}

func NewProject(name, description, author, sponsor, reference string, projectType ProjectType) (p Project) {
	p.Name = name
	p.Description = description
	p.Author = author
	p.Sponsor = sponsor
	p.Reference = reference
	p.Stage = ProjectStageSandbox
	p.Type = projectType
	return
}

func (p *Project) Promote() {
	switch p.Stage {
	case ProjectStageSandbox:
		p.Stage = ProjectStageProposed
	case ProjectStageProposed:
		p.Stage = ProjectStageGraduated
	}
}

func (p Project) RenderMarkdown() (markdown string) {
	markdown += fmt.Sprintf("# %s - %s\n\n", p.ID, p.Name)
	markdown += fmt.Sprintf("%s\n\n", p.Description)
	markdown += fmt.Sprintln("---\n")
	markdown += fmt.Sprintf("Author: %s\n\n", p.Author)
	markdown += fmt.Sprintf("Sponsor: %s\n\n", p.Sponsor)
	markdown += fmt.Sprintf("Type: %s\n\n", p.Type)
	markdown += fmt.Sprintf("Reference: %s\n\n", p.Reference)
	markdown += fmt.Sprintf("Stage: **%s**\n\n", p.Stage)

	return markdown
}
