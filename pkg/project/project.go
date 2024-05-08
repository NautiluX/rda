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
	ProjectStageSandbox    ProjectStage = "sandbox"
	ProjectStageIncubation ProjectStage = "incubation"
	ProjectStageGraduated  ProjectStage = "graduated"
)

type Project struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Author      string            `yaml:"author"`
	Sponsor     string            `yaml:"sponsor"`
	Reference   string            `yaml:"reference"`
	ID          string            `yaml:"id"`
	Epics       map[string]string `yaml:"epics"`
	Stage       ProjectStage      `yaml:"stage"`
	Type        ProjectType       `yaml:"type"`
}

func NewProject(name, description, author, sponsor, reference, sandboxEpic string, projectType ProjectType) (p Project) {
	p.Name = name
	p.Description = description
	p.Author = author
	p.Sponsor = sponsor
	p.Reference = reference
	p.Stage = ProjectStageSandbox
	p.Type = projectType
	p.Epics = map[string]string{string(ProjectStageSandbox): sandboxEpic}
	return
}

func (p *Project) Promote(epic string) {
	switch p.Stage {
	case ProjectStageSandbox:
		p.Stage = ProjectStageIncubation
	case ProjectStageIncubation:
		p.Stage = ProjectStageGraduated
	}
	if epic != "" {
		p.Epics[string(p.Stage)] = epic
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
	markdown += fmt.Sprintln("## Epics\n")
	for stage, link := range p.Epics {
		markdown += fmt.Sprintf("* %s: %s\n", stage, link)
	}

	return markdown
}
